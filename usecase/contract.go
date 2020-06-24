package usecase

import (
	"bukuduit-go/helpers/aes"
	"bukuduit-go/helpers/aesfront"
	queue "bukuduit-go/helpers/amqp"
	"bukuduit-go/helpers/jwe"
	"bukuduit-go/helpers/jwt"
	"bukuduit-go/helpers/messages"
	"bukuduit-go/helpers/str"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v7"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sfreiberg/gotwilio"
	"github.com/streadway/amqp"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	defaultLimit    = 10
	maxLimit        = 50
	defaultOrderBy  = "id"
	defaultSort     = "asc"
	PasswordLength  = 6
	defaultLastPage = 0
	OtpLifeTime       = "3m"
	MaxOtpSubmitRetry = 3
)

//globalsmscounter
var GlobalSmsCounter int

// AmqpConnection ...
var AmqpConnection *amqp.Connection

// AmqpChannel ...
var AmqpChannel *amqp.Channel

//X-Request-ID
var xRequestID interface{}

type UcContract struct {
	E          *echo.Echo
	DB         *sql.DB
	Redis      *redis.Client
	Aes        aes.Credential
	AesFront   aesfront.Credential
	Jwe        jwe.Credential
	Validate   *validator.Validate
	Translator ut.Translator
	JwtConfig  middleware.JWTConfig
	JwtCred    jwt.JwtCredential
	Twilio     *gotwilio.Twilio
}

func (uc UcContract) setPaginationParameter(page, limit int, order, sort string) (int, int, int, string, string) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 || limit > maxLimit {
		limit = defaultLimit
	}

	if order == "" {
		order = defaultOrderBy
	}

	if sort == "" {
		sort = defaultSort
	}

	offset := (page - 1) * limit

	return offset, limit, page, order, sort
}

func (uc UcContract) setPaginationResponse(page, limit, total int) (paginationResponse viewmodel.PaginationVm) {
	var lastPage int

	if total > 0 {
		lastPage = total / limit

		if total%limit != 0 {
			lastPage = lastPage + 1
		}
	} else {
		lastPage = defaultLastPage
	}

	paginationResponse = viewmodel.PaginationVm{
		CurrentPage: page,
		LastPage:    lastPage,
		Total:       total,
		PerPage:     limit,
	}

	return paginationResponse
}

func (uc UcContract) GetRandomString(length int) string {
	if length == 0 {
		length = PasswordLength
	}

	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")

	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	password := b.String()

	return password
}

func (uc UcContract) StoreToRedistWithExpired(key string, val interface{}, duration string) error {
	dur, err := time.ParseDuration(duration)
	if err != nil {
		return err
	}

	b, err := json.Marshal(val)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = uc.Redis.Set(key, string(b), dur).Err()

	return err
}

func (uc UcContract) StoreToRedis(key string, val interface{}) error {
	b, err := json.Marshal(val)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = uc.Redis.Set(key, string(b), 0).Err()

	return err
}

func (uc UcContract) GetFromRedis(key string, cb interface{}) error {
	res, err := uc.Redis.Get(key).Result()
	if err != nil {
		return err
	}

	if res == "" {
		return errors.New("[Redis] Value of " + key + " is empty.")
	}

	err = json.Unmarshal([]byte(res), &cb)
	if err != nil {
		return err
	}

	return err
}

func (uc UcContract) RemoveFromRedis(key string) error {
	return uc.Redis.Del(key).Err()
}

func (uc UcContract) RefNumberGenerator(code string, count int, now time.Time) (codeNumber string) {
	codeNumber = code + now.Format("0601") + fmt.Sprintf("%06d", count+1)

	return codeNumber
}

func (uc UcContract) LimitRetryByKey(key string, limit float64) (err error) {
	var count float64
	res := map[string]interface{}{}

	err = uc.GetFromRedis(key, &res)
	if err != nil {
		err = nil
		res = map[string]interface{}{
			"counter": count,
		}
	}
	count = res["counter"].(float64) + 1
	if count > limit {
		uc.RemoveFromRedis(key)

		return errors.New(messages.MaxRetryKey)
	}

	res["counter"] = count
	uc.StoreToRedistWithExpired(key, res, "24h")

	return err
}

func (uc UcContract) GenerateAesKey(redisKey string, value interface{}) (res string, err error) {
	rand := str.RandStringBytesMaskImprSrc(10)
	res, err = uc.Aes.Encrypt(rand)
	if err != nil {
		return res, errors.New(messages.InternalServer)
	}
	err = uc.StoreToRedistWithExpired(redisKey, value, "24h")
	if err != nil {
		return res, errors.New(messages.InternalServer)
	}
	redisValue := map[string]interface{}{
		"key":   res,
		"count": 0,
	}
	err = uc.StoreToRedistWithExpired(redisKey, redisValue, "24h")
	if err != nil {
		return res, errors.New(messages.InternalServer)
	}

	return res, err
}

func (uc UcContract) SetXRequestID(ctx echo.Context) {
	xRequestID = ctx.Get(echo.HeaderXRequestID)
}

func (uc UcContract) GetXRequestID() interface{} {
	return xRequestID
}

func (uc UcContract) GenerateOrderID(key, duration string, ID int) (res string, err error) {
	res = str.RandomNumberString(6)
	redisValue := map[string]interface{}{
		"orderID": res,
	}

	err = uc.StoreToRedistWithExpired("NPLT-"+"-"+res+"-"+string(ID), redisValue, duration)

	return res, err
}

func (uc UcContract) PushToQueue(queueBody map[string]interface{}, queueType, deadLetterType string) (err error) {
	mqueue := queue.NewQueue(AmqpConnection, AmqpChannel)

	_, _, err = mqueue.PushQueueReconnect(os.Getenv("AMQP_URL"), queueBody, queueType, deadLetterType)
	if err != nil {
		return err
	}

	return err
}
