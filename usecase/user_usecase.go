package usecase

import (
	"bukuduit-go/db/repositories/actions"
	queue "bukuduit-go/helpers/amqp"
	"bukuduit-go/helpers/hashing"
	"bukuduit-go/helpers/messages"
	"bukuduit-go/helpers/str"
	request "bukuduit-go/server/requests"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type UserUseCase struct {
	*UcContract
}

func (uc UserUseCase) ForgotMyPin(input request.UserRequest) (err error) {
	pin := str.RandomNumberString(6)

	isMobilePhoneExist, err := uc.IsMobilePhoneExist(input.MobilePhone)
	if err != nil {
		return err
	}
	if isMobilePhoneExist {
		hasPin, _ := hashing.HashAndSalt(pin)
		err = uc.EditPin(input.MobilePhone, hasPin)
		if err != nil {
			return err
		}

		// Send sms Queue
		requestID := fmt.Sprintf("%v", uc.GetXRequestID())
		queueBody := map[string]interface{}{
			"qid":     xRequestID,
			"phone":   input.MobilePhone,
			"message": `Hai sobat! OTP untuk BukuDuit.com anda adalah ` + pin,
			"type":    "sms",
			"id":      requestID,
		}
		err = uc.PushToQueue(queueBody, queue.Otp, queue.OtpDeadLetter)
		if err != nil {
			return err
		}
	}

	return err
}

func (uc UserUseCase) MyAccount(userID string) (res viewmodel.UserVm, err error) {

	businessUsecase := BusinessCardUseCase{UcContract: uc.UcContract}
	paymentUC := PaymentAccountUseCase{UcContract: uc.UcContract}
	user, err := uc.ReadByPk(userID)
	if err != nil {
		return res, err
	}

	var ownedShop []viewmodel.BusinessCardVm
	ownedShop, err = businessUsecase.BrowseByUser(userID)
	fmt.Println(ownedShop)
	if err != nil {
		return res, err
	}

	var detail []viewmodel.ListPaymentAcc
	var temp []viewmodel.PaymentAccountVm
	for _, i := range ownedShop {
		temp, err = paymentUC.BrowseByShop(i.ID)

		detail = append(detail, viewmodel.ListPaymentAcc{
			ListAccPayment: temp,
		})
	}
	fmt.Println(detail)

	if err != nil {
		return res, err
	}

	var paymentDetails []viewmodel.PaymentAccountVm

	fmt.Println(paymentDetails)
	res = viewmodel.UserVm{
		ID:             user.ID,
		MobilePhone:    user.MobilePhone,
		PaymentDetails: detail,
		Name:           user.Name,
	}
	return res, err
}
func (uc UserUseCase) ReadByPk(ID string) (res viewmodel.UserVm, err error) {
	model := actions.NewUserModel(uc.DB)
	user, err := model.ReadByPk(ID)
	if err != nil {
		return res, err
	}

	res = viewmodel.UserVm{
		ID:          user.ID,
		MobilePhone: user.MobilePhone,
		Name:        user.Name.String,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt.String,
		DeletedAt:   user.DeletedAt.String,
	}

	return res, err
}

func (uc UserUseCase) ReadBy(column, value string) (res viewmodel.UserVm, err error) {
	model := actions.NewUserModel(uc.DB)
	user, err := model.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	res = viewmodel.UserVm{
		ID:          user.ID,
		MobilePhone: user.MobilePhone,
		Name:        user.Name.String,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt.String,
		DeletedAt:   user.DeletedAt.String,
	}

	return res, err
}

func (uc UserUseCase) Edit(input request.UserRequest) (err error) {
	model := actions.NewUserModel(uc.DB)
	now := time.Now().UTC()
	fmt.Println(input.ID)
	if input.OldPin != "" {
		isPINMatch, err := uc.IsPINMatch(input.MobilePhone, input.OldPin)
		if err != nil {
			fmt.Println("error pin match")
			return errors.New(messages.CredentialDoNotMatch)
		}
		if !isPINMatch {
			fmt.Println("error pin tidak sama")
			return errors.New(messages.CredentialDoNotMatch)
		}

		hasPin, _ := hashing.HashAndSalt(input.NewPin)
		err = uc.EditPin(input.MobilePhone, hasPin)
		if err != nil {
			return err
		}
	}

	res := viewmodel.UserVm{
		ID:          input.ID,
		Name:        input.Name,
		MobilePhone: input.MobilePhone,
		UpdatedAt:   now.Format(time.RFC3339),
	}
	if input.Name != "" {
		_, err = model.Edit(res)
		if err != nil {
			return err
		}
	}

	return nil
}

func (uc UserUseCase) EditPin(phone, pin string) (err error) {
	model := actions.NewUserModel(uc.DB)
	now := time.Now().UTC()
	_, err = model.EditPin(phone, pin, now.Format(time.RFC3339))
	if err != nil {
		return err
	}

	return nil
}

func (uc UserUseCase) Add(mobilePhone, pin string, tx *sql.Tx) (res string, err error) {
	model := actions.NewUserModel(uc.DB)
	now := time.Now().UTC()
	isMobilePhoneExist, err := uc.IsMobilePhoneExist(mobilePhone)
	if err != nil {
		return res, err
	}
	if isMobilePhoneExist {
		return res, errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.UserVm{
		MobilePhone: mobilePhone,
		CreatedAt:   now.Format(time.RFC3339),
	}
	hasPin, _ := hashing.HashAndSalt(pin)
	res, err = model.Add(body, hasPin, tx)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc UserUseCase) CountBy(column, value string) (res int, err error) {
	model := actions.NewUserModel(uc.DB)
	res, err = model.CountBy(column, value)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc UserUseCase) IsMobilePhoneExist(mobilePhone string) (res bool, err error) {
	count, err := uc.CountBy("mobile_phone", mobilePhone)
	if err != nil {
		return res, err
	}

	return count > 0, err
}

func (uc UserUseCase) IsPINMatch(mobilePhone, PIN string) (res bool, err error) {
	model := actions.NewUserModel(uc.DB)
	user, err := model.ReadBy("mobile_phone", mobilePhone)
	if err != nil {
		return res, err
	}
	res = hashing.CheckHashString(PIN, user.Pin)

	return res, err
}
