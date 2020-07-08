package main

import (
	"bukuduit-go/helpers/logruslogger"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sfreiberg/gotwilio"
	"log"
	"os"
	"runtime"

	"bukuduit-go/helpers/aes"
	amqpPkg "bukuduit-go/helpers/amqp"
	"bukuduit-go/helpers/amqpconsumer"
	"bukuduit-go/usecase"

	"github.com/go-redis/redis/v7"
	"github.com/streadway/amqp"
)

var (
	uri          *string
	formURL      = flag.String("form_url", "http://localhost", "The URL that requests are sent to")
	logFile      = flag.String("log_file", "system.log", "The file where errors are logged")
	threads      = flag.Int("threads", 1, "The max amount of go routines that you would like the process to use")
	maxprocs     = flag.Int("max_procs", 1, "The max amount of processors that your application should use")
	paymentsKey  = flag.String("payments_key", "secret", "Access key")
	exchange     = flag.String("exchange", amqpPkg.OtpExchange, "The exchange we will be binding to")
	exchangeType = flag.String("exchange_type", "direct", "Type of exchange we are binding to | topic | direct| etc..")
	queue        = flag.String("queue", amqpPkg.Otp, "Name of the queue that you would like to connect to")
	routingKey   = flag.String("routing_key", amqpPkg.OtpDeadLetter, "queue to route messages to")
	workerName   = flag.String("worker_name", "worker.name", "name to identify worker by")
	verbosity    = flag.Bool("verbos", false, "Set true if you would like to log EVERYTHING")

	// Hold consumer so our go routine can listen to
	// it's done error chan and trigger reconnects
	// if it's ever returned
	conn      *amqpconsumer.Consumer
	envConfig map[string]string
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading ..env file")
	}

	flag.Parse()
	runtime.GOMAXPROCS(*maxprocs)
	uri = flag.String("uri", os.Getenv("AMQP_URL"), "The rabbitmq endpoint")
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading ..env file")
	}

	file := false
	// Open a system file to start logging to
	if file {
		f, err := os.OpenFile(*logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		defer f.Close()
		if err != nil {
			log.Printf("error opening file: %v", err.Error())
		}
		log.SetOutput(f)
	}

	conn := amqpconsumer.NewConsumer(*workerName, *uri, *exchange, *exchangeType, *queue)

	if err := conn.Connect(); err != nil {
		log.Printf("Error: %v", err)
	}

	deliveries, err := conn.AnnounceQueue(*queue, *routingKey)
	if err != nil {
		log.Printf("Error when calling AnnounceQueue(): %v", err.Error())
	}
	//setup twilio
	twilio := gotwilio.NewTwilioClient(os.Getenv("TWILLIO_SID"), os.Getenv("TWILIO_TOKEN"))

	//setup redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	//check redis
	pong, err := redisClient.Ping().Result()
	fmt.Println("Redis ping status: "+pong, err)

	// AES credential
	aesCredential := aes.Credential{
		Key: os.Getenv("AES_KEY"),
	}

	cUC := usecase.UcContract{
		Redis:  redisClient,
		Aes:    aesCredential,
		Twilio: twilio,
	}

	conn.Handle(deliveries, handler, *threads, *queue, *routingKey, cUC)
}

func handler(deliveries <-chan amqp.Delivery, uc *usecase.UcContract) {
	ctx := "bukuduit.otplistener"
	for d := range deliveries {
		var formData map[string]interface{}
		err := json.Unmarshal(d.Body, &formData)
		if err != nil {
			log.Printf("Error unmarshaling data: %s", err.Error())
		}

		smsUc := usecase.SmsUseCase{
			UcContract: uc,
		}
		err = smsUc.SendSms(formData["message"].(string), formData["phone"].(string))
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, "{body: "+string(d.Body)+", err: "+err.Error()+"}", ctx, "err", formData["id"].(string))

			// Get fail counter from redis
			failCounter := amqpconsumer.FailCounter{}
			err = uc.GetFromRedis("amqpFail"+formData["id"].(string), &failCounter)
			if err != nil {
				failCounter = amqpconsumer.FailCounter{
					Counter: 1,
				}
			}

			if failCounter.Counter > amqpconsumer.MaxFailCounter {
				log.Printf("Rejected message")
				d.Reject(false)
			} else {
				// Save the new counter to redis
				failCounter.Counter++
				err = uc.StoreToRedistWithExpired("amqpFail"+formData["id"].(string), failCounter, "10m")

				log.Printf("Unacknowledged message")
				d.Nack(false, true)
			}
		} else {
			logruslogger.Log(logruslogger.InfoLevel, string(d.Body), ctx, "success", formData["id"].(string))
			log.Printf("Acknowledged message")
			d.Ack(false)
		}
	}

	return
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
