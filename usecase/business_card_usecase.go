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

//fucntion for hutang list lunas
func (uc BusinessCardUseCase) Read(ID, lunas string) (res viewmodel.BusinessCardVm, err error) { //lunas = 1
	model := actions.NewBusinessCardModel(uc.DB)
	userCustomerUC := UserCustomerUseCase{UcContract: uc.UcContract}
	bookDebtUc := BooksDebtUseCase{UcContract:uc.UcContract}
	var debtTotal int
	var creditTotal int

	tempDataUserCustomer, err := userCustomerUC.BrowseByShop(ID)
	if err != nil {
		return res, err
	}

	var elems int
	var amount int32
	var typeBook string
	var dataUserCustomer = make([]viewmodel.UserCustomerDebetCreditVm, elems)
	for _, data := range tempDataUserCustomer {
		if lunas == "1" {
			bookdebtsLunas, err := bookDebtUc.BrowseByUser(data.ID,enums.Lunas)
			if err == nil {
				if bookdebtsLunas.CreditTotal == 0{
					dataUserCustomer = append(dataUserCustomer, viewmodel.UserCustomerDebetCreditVm{
						ID:          data.ID,
						FullName:    data.FullName,
						Amount: int32(bookdebtsLunas.CreditTotal),
						Type: typeBook,
					})
				}
			}
		} else {
			bookdebtsNunggak, err := bookDebtUc.BrowseByUser(data.ID,"")
			if err == nil {
				creditTotal = creditTotal + bookdebtsNunggak.CreditTotal
				debtTotal = debtTotal + bookdebtsNunggak.DebtTotal
				if bookdebtsNunggak.CreditTotal != 0 {
					typeBook = enums.Credit
					amount = int32(bookdebtsNunggak.CreditTotal)
				}else{
					typeBook = enums.Debet
					amount = int32(bookdebtsNunggak.DebtTotal)
				}
				dataUserCustomer = append(dataUserCustomer, viewmodel.UserCustomerDebetCreditVm{
					ID:          data.ID,
					FullName:    data.FullName,
					Amount: amount,
					Type: typeBook,
				})
			}
		}
	}

	//for _, k := range dataUserCustomer {
	//	if k.Debt > 0 {
	//		creditTotal = creditTotal + int(k.Debt)
	//	} else {
	//		debtTotal = debtTotal + (int(k.Debt) * -1)
	//	}
	//}

	businessCard, err := model.Read(ID)
	if err != nil {
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
