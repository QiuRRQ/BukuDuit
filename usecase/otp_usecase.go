package usecase

import (
	queue "bukuduit-go/helpers/amqp"
	"bukuduit-go/helpers/messages"
	"bukuduit-go/helpers/str"
	"bukuduit-go/usecase/viewmodel"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type OtpUseCase struct {
	*UcContract
}

// RequestOtp ...
func (uc OtpUseCase) RequestOtp(mobilePhoneNumber string) (res viewmodel.OtpVm, err error) {
	// Cehck Phone uniqe
	userUc := UserUseCase{UcContract: uc.UcContract}
	isExist, err := userUc.IsMobilePhoneExist(mobilePhoneNumber)
	if err != nil {
		return res,err
	}
	if isExist {
		return res, errors.New(messages.DataAlreadyExist)
	}

	//generate otp save to redis
	rand.Seed(time.Now().UTC().UnixNano())
	res.MobilePhone = mobilePhoneNumber
	res.Otp, err = uc.generateOtpCode(mobilePhoneNumber, res)
	if err != nil {
		return res, err
	}
	//add otp invalid counter
	err = uc.addInvalidOtpCounter(mobilePhoneNumber)
	if err != nil {
		return res, err
	}

	// Send sms Queue
	requestID := fmt.Sprintf("%v", uc.GetXRequestID())
	queueBody := map[string]interface{}{
		"qid":     xRequestID,
		"phone":   mobilePhoneNumber,
		"message": res.Otp,
		"type":    "sms",
		"id":      requestID,
	}
	err = uc.PushToQueue(queueBody, queue.Otp, queue.OtpDeadLetter)
	if err != nil {
		fmt.Println(err)
	}

	return res, err
}

func (uc OtpUseCase) generateOtpCode(mobilePhone string, otpVm viewmodel.OtpVm) (res string, err error) {
	res = str.RandomNumberString(4)
	otpVm.Otp = res
	err = uc.StoreToRedistWithExpired("otp"+mobilePhone, otpVm, OtpLifeTime)
	if err != nil {
		return res, errors.New(messages.OtpStoreToRedis)
	}

	return res, err
}

func (uc OtpUseCase) addInvalidOtpCounter(mobilePhone string) (err error) {
	c := viewmodel.InvalidOtpCounterVM{
		Counter: 0,
	}
	err = uc.StoreToRedistWithExpired("invalidOtp"+mobilePhone, c, OtpLifeTime)
	if err != nil {
		return errors.New(messages.InvalidOtpStoreToRedis)
	}

	return err
}
