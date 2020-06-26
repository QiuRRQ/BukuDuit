package usecase

import (
	"bukuduit-go/db/repositories/actions"
	"bukuduit-go/helpers/enums"
	"bukuduit-go/helpers/messages"
	request "bukuduit-go/server/requests"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type BusinessCardUseCase struct {
	*UcContract
}

func (uc BusinessCardUseCase) BrowseByUser(userID string) (res []viewmodel.BusinessCardVm, err error) {
	model := actions.NewBusinessCardModel(uc.DB)
	businessCards, err := model.BrowseByUser(userID)
	if err != nil {
		return res, err
	}

	for _, businessCard := range businessCards {
		res = append(res, viewmodel.BusinessCardVm{
			ID:          businessCard.ID,
			FullName:    businessCard.FullName.String,
			BookName:    businessCard.BookName,
			MobilePhone: businessCard.MobilePhone.String,
			TagLine:     businessCard.TagLine.String,
			Address:     businessCard.Address.String,
			Email:       businessCard.Email.String,
			Avatar:      businessCard.Avatar.String,
			CreatedAt:   businessCard.CreatedAt,
			UpdatedAt:   businessCard.UpdatedAt.String,
			DeletedAt:   businessCard.DeletedAt.String,
		})
	}

	return res, err
}

//fucntion for hutang list
func (uc BusinessCardUseCase) Read(ID string) (res viewmodel.BusinessCardVm, err error) {
	model := actions.NewBusinessCardModel(uc.DB)
	userCustomerUC := UserCustomerUseCase{UcContract: uc.UcContract}
	transactionUC := TransactionUseCase{UcContract: uc.UcContract}
	var debtTotal int
	var creditTotal int

	dataTransaction, err := transactionUC.BrowseByShop(ID)
	if err != nil {
		return res, err
	}
	dataUserCustomer, err := userCustomerUC.BrowseByShop(ID)
	if err != nil {
		return res, err
	}

	for _, k := range dataUserCustomer {
		creditTotal = creditTotal + int(k.Debt)
	}

	for _, v := range dataTransaction {

		if v.Type == enums.Debet {
			debtTotal = debtTotal + int(v.Amount)
		}
	}

	businessCard, err := model.Read(ID)
	if err != nil {
		fmt.Println(3)
		return res, err
	}

	res = viewmodel.BusinessCardVm{
		ID:                  businessCard.ID,
		FullName:            businessCard.FullName.String,
		BookName:            businessCard.BookName,
		MobilePhone:         businessCard.MobilePhone.String,
		TagLine:             businessCard.TagLine.String,
		Address:             businessCard.Address.String,
		Email:               businessCard.Email.String,
		UserCustomers:       dataUserCustomer,
		TotalCustomerCredit: int32(creditTotal),
		TotalOwnerCredit:    int32(debtTotal),
		CreatedAt:           businessCard.CreatedAt,
		UpdatedAt:           businessCard.UpdatedAt.String,
		DeletedAt:           businessCard.UpdatedAt.String,
	}

	return res, err
}

func (uc BusinessCardUseCase) Edit(input *request.BusinessCardRequest, ID string) (err error) {
	model := actions.NewBusinessCardModel(uc.DB)
	now := time.Now().UTC()
	isExist, err := uc.IsBusinessCardExist(ID)
	if err != nil {
		return err
	}
	if !isExist {
		return errors.New(messages.DataNotFound)
	}

	body := viewmodel.BusinessCardVm{
		ID:          ID,
		FullName:    input.FullName,
		BookName:    input.BookName,
		MobilePhone: input.MobilePhone,
		TagLine:     input.TagLine,
		Address:     input.Address,
		Email:       input.Email,
		Avatar:      input.Avatar,
		UpdatedAt:   now.Format(time.RFC3339),
	}
	_, err = model.Edit(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc BusinessCardUseCase) Add(input *request.BusinessCardRequest, userID string) (err error) {
	model := actions.NewBusinessCardModel(uc.DB)
	now := time.Now().UTC()

	body := viewmodel.BusinessCardVm{
		BookName:  input.BookName,
		CreatedAt: now.Format(time.RFC3339),
		UpdatedAt: now.Format(time.RFC3339),
	}
	_, err = model.Add(body, userID, nil)
	if err != nil {
		return err
	}

	return nil
}

func (uc BusinessCardUseCase) Register(userID, bookName, createdAt, updatedAt string, tx *sql.Tx) (err error) {
	model := actions.NewBusinessCardModel(uc.DB)
	body := viewmodel.BusinessCardVm{
		BookName:  bookName,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	_, err = model.Add(body, userID, tx)
	if err != nil {
		return err
	}

	return nil
}

func (uc BusinessCardUseCase) Delete(ID string) (err error) {
	fmt.Println(ID)
	model := actions.NewBusinessCardModel(uc.DB)
	now := time.Now().UTC()

	isExist, err := uc.IsBusinessCardExist(ID)
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

func (uc BusinessCardUseCase) DeleteByUser(userID string, tx *sql.Tx) (err error) {
	model := actions.NewBusinessCardModel(uc.DB)
	now := time.Now().UTC()

	err = model.DeleteByUser(userID, now.Format(time.RFC3339), now.Format(time.RFC3339), tx)
	if err != nil {
		return err
	}

	return nil
}

func (uc BusinessCardUseCase) IsBusinessCardExist(ID string) (res bool, err error) {
	model := actions.NewBusinessCardModel(uc.DB)
	count, err := model.CountByPk(ID)
	if err != nil {
		return res, err
	}

	return count > 0, err
}
