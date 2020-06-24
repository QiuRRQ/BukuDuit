package usecase

import (
	queue "bukuduit-go/helpers/amqp"
	"bukuduit-go/helpers/messages"
	"bukuduit-go/helpers/str"
	"bukuduit-go/usecase/viewmodel"
	"errors"
	"fmt"
)

type OtpUseCase struct {
	*UcContract
}

// RequestOtp ...

func (uc OtpUseCase) RequestOtp(mobilePhoneNumber,action string) (res viewmodel.OtpVm, err error) {
	if action == "register" {
		userUc := UserUseCase{UcContract:uc.UcContract}
		isExist,err := userUc.IsMobilePhoneExist(mobilePhoneNumber)
		if err != nil {
			return res,err
		}
		if isExist {
			return res, errors.New(messages.PhoneAlreadyExist)
		}
	}

	err = uc.GetFromRedis("otp"+mobilePhoneNumber,&res)
	if err == nil {
		uc.RemoveFromRedis("otp" + mobilePhoneNumber)
	}

	//generate otp save to redis
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
		return res,err
	}

	return res, err
}

func (uc OtpUseCase) generateOtpCode(mobilePhone string, otpVm viewmodel.OtpVm) (res string, err error) {
	res = str.RandomNumberString(6)
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

func (uc OtpUseCase) SubmitOtp(key string, otp string) (res viewmodel.OtpVm, err error) {
	otpVm := viewmodel.OtpVm{}
	fmt.Println(key)
	err = uc.GetFromRedis(key, &otpVm)
	if err != nil {
		return res, errors.New(messages.ExpiredOtp)
	}

	if otpVm.Otp != otp {
		err = uc.LimitRetryByKey("invalidOtp"+otpVm.MobilePhone, MaxOtpSubmitRetry)
		if err != nil {
			return res, err
		}
	}

	return res, nil
}
