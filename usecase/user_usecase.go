package usecase

import (
	"bukuduit-go/db/repositories/actions"
	"bukuduit-go/helpers/hashing"
	"bukuduit-go/helpers/messages"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
	"errors"
	"time"
)

type UserUseCase struct {
	*UcContract
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
	if err != nil {
		return res, err
	}

	var detail []viewmodel.PaymentAccountVm
	for _, i := range ownedShop {
		detail, err = paymentUC.BrowseByShop(i.ID)
	}

	if err != nil {
		return res, err
	}

	var paymentDetails []viewmodel.PaymentAccountVm

	for _, v := range detail {
		paymentDetails = append(paymentDetails, viewmodel.PaymentAccountVm{
			ID:            v.ID,
			ShopID:        v.ShopID,
			Name:          v.Name,
			PaymentNumber: v.PaymentNumber,
		})
	}

	res = viewmodel.UserVm{
		ID:             user.ID,
		MobilePhone:    user.MobilePhone,
		PaymentDetails: paymentDetails,
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

func (uc UserUseCase) EditPin(ID, pin string) (err error) {
	model := actions.NewUserModel(uc.DB)
	now := time.Now().UTC()
	_, err = model.EditPin(ID, pin, now.Format(time.RFC3339))
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
