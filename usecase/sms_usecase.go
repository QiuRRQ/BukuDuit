package usecase

import (
	"os"
)

type SmsUseCase struct {
	*UcContract
}

func (uc SmsUseCase) SendSms(otp,receiver string) (err error) {
	sender := os.Getenv("TWILLIO_TESTING_PHONE")
	message := `Hai sobat! OTP untuk BukuDuit.com anda adalah `+otp
	_,_,err = uc.Twilio.SendSMS(sender,receiver, message,"","")
	if err != nil {
		return err
	}

	return nil
}