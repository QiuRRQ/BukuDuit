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

func (uc UserUseCase) ReadByPk(ID string) (res viewmodel.UserVm, err error) {
	model := actions.NewUserModel(uc.DB)
	user, err := model.ReadByPk(ID)
	if err != nil {
		return res, err
	}

	res = viewmodel.UserVm{
		ID:          user.ID,
		MobilePhone: user.MobilePhone,
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

func (uc UserUseCase) Add(mobilePhone,pin string, tx *sql.Tx) (res string,err error){
	model := actions.NewUserModel(uc.DB)
	now := time.Now().UTC()
	isMobilePhoneExist, err := uc.IsMobilePhoneExist(mobilePhone)
	if err != nil {
		return res,err
	}
	if isMobilePhoneExist {
		return res,errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.UserVm{
		MobilePhone: mobilePhone,
		CreatedAt:   now.Format(time.RFC3339),
	}
	hasPin,_ := hashing.HashAndSalt(pin)
	res, err = model.Add(body,hasPin,tx)
	if err != nil {
		return res,err
	}

	return res,err
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

func (uc UserUseCase) IsPINMatch(mobilePhone, PIN string) (res bool, err error){
	model := actions.NewUserModel(uc.DB)
	user,err := model.ReadBy("mobile_phone",mobilePhone)
	if err != nil {
		return res,err
	}
	res = hashing.CheckHashString(PIN,user.Pin)

	return res,err
}
