package usecase

import (
	"os"
)

type SmsUseCase struct {
	*UcContract
}

func (uc SmsUseCase) SendSms(message, receiver string) (err error) {
	sender := os.Getenv("TWILLIO_TESTING_PHONE")
	_, _, err = uc.Twilio.SendSMS(sender, receiver, message, "", "")
	if err != nil {
		return err
	}

	return nil
}
