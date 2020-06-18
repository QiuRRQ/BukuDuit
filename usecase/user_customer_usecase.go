package usecase

import (
	"bukuduit-go/db/repositories/actions"
	"bukuduit-go/helpers/messages"
	request "bukuduit-go/server/requests"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
	"errors"
	"time"
)

type UserCustomerUseCase struct {
	*UcContract
}

func (uc UserCustomerUseCase) BrowseByShop(shopId string) (res []viewmodel.UserCustomerVm, err error) {
	model := actions.NewUserCustomerModel(uc.DB)
	userCustomers, err := model.BrowseByBusiness(shopId)
	if err != nil {
		return res, err
	}

	for _, userCustomer := range userCustomers {
		res = append(res, viewmodel.UserCustomerVm{
			ID:          userCustomer.ID,
			FullName:    userCustomer.FullName,
			MobilePhone: userCustomer.MobilePhone,
			Debt:        userCustomer.Debt.Int32,
			CreatedAt:   userCustomer.CreatedAt,
			UpdatedAt:   userCustomer.UpdatedAt.String,
			DeletedAt:   userCustomer.DeletedAt.String,
		})
	}

	return res, err
}

func (uc UserCustomerUseCase) Read(ID string) (res viewmodel.UserCustomerVm, err error) {
	model := actions.NewUserCustomerModel(uc.DB)
	userCustomer, err := model.Read(ID)
	if err != nil {
		return res, err
	}

	res = viewmodel.UserCustomerVm{
		ID:          userCustomer.ID,
		FullName:    userCustomer.FullName,
		MobilePhone: userCustomer.MobilePhone,
		Debt:        userCustomer.Debt.Int32,
		CreatedAt:   userCustomer.CreatedAt,
		UpdatedAt:   userCustomer.UpdatedAt.String,
		DeletedAt:   userCustomer.DeletedAt.String,
	}

	return res, err
}

func (uc UserCustomerUseCase) EditDebt(ID string, debt int32, tx *sql.Tx) (err error) {
	model := actions.NewUserCustomerModel(uc.DB)
	now := time.Now().UTC()

	err = model.EditDebt(ID, now.Format(time.RFC3339), debt, tx)
	if err != nil {
		return err
	}

	return nil
}

func (uc UserCustomerUseCase) Add(input *request.UserCustomerRequest) (err error) {
	model := actions.NewUserCustomerModel(uc.DB)
	now := time.Now().UTC()

	body := viewmodel.UserCustomerVm{
		FullName:    input.FullName,
		MobilePhone: input.MobilePhone,
		Debt:        0,
		CreatedAt:   now.Format(time.RFC3339),
		UpdatedAt:   now.Format(time.RFC3339),
	}
	_, err = model.Add(body, input.ShopID)
	if err != nil {
		return err
	}

	return nil
}

func (uc UserCustomerUseCase) Delete(ID string) (err error) {
	model := actions.NewUserCustomerModel(uc.DB)
	now := time.Now().UTC()

	isExist, err := uc.IsUserCustomerExist(ID)
	if err != nil {
		return err
	}
	if !isExist {
		return errors.New(messages.DataNotFound)
	}

	_, err = model.Delete(ID, now.Format(time.RFC3339), now.Format(time.RFC3339))
	if err != nil {
		return err
	}

	return nil
}

func (uc UserCustomerUseCase) CountByPk(ID string) (res int, err error) {
	model := actions.NewUserCustomerModel(uc.DB)
	res, err = model.CountByPk(ID)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc UserCustomerUseCase) IsUserCustomerExist(ID string) (res bool, err error) {
	count, err := uc.CountByPk(ID)
	if err != nil {
		return res, err
	}

	return count > 0, err
}
