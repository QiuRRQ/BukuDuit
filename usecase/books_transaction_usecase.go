package usecase

import (
	"bukuduit-go/db/models"
	"bukuduit-go/db/repositories/actions"
	"bukuduit-go/helpers/messages"
	request "bukuduit-go/server/requests"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type BooksTransactionUseCase struct {
	*UcContract
}

func (uc BooksTransactionUseCase) BrowseByUser(shopID string) (res []viewmodel.BooksTransactionVm, err error) {
	model := actions.NewBooksTransactionModel(uc.DB)
	booksTrans, err := model.BrowseByShop(shopID)
	if err != nil {
		return res, err
	}

	for _, books := range booksTrans {
		res = append(res, viewmodel.BooksTransactionVm{
			ID:          books.ID,
			ShopID:      books.ShopID.String,
			DebtTotal:   books.DebtTotal,
			CreditTotal: books.CreditTotal,
			CreatedAt:   books.CreatedAt,
			UpdatedAt:   books.UpdatedAt.String,
			DeletedAt:   books.DeletedAt.String,
		})
	}

	return res, err
}

//fucntion for hutang list lunas
func (uc BooksTransactionUseCase) Read(ID, lunas string) (res viewmodel.BooksTransactionVm, err error) { //lunas = 1
	model := actions.NewBooksTransactionModel(uc.DB)

	businessCard, err := model.Read(ID)
	if err != nil {
		return res, err
	}

	res = viewmodel.BooksTransactionVm{
		ID:          businessCard.ID,
		ShopID:      businessCard.ShopID.String,
		DebtTotal:   businessCard.DebtTotal,
		CreditTotal: businessCard.CreditTotal,
		CreatedAt:   businessCard.CreatedAt,
		UpdatedAt:   businessCard.UpdatedAt.String,
		DeletedAt:   businessCard.UpdatedAt.String,
	}

	return res, err
}

func (uc BooksTransactionUseCase) Edit(input request.BooksTransactionRequest, ID string) (err error) {
	model := actions.NewBooksTransactionModel(uc.DB)
	now := time.Now().UTC()
	isExist, err := uc.IsShopIDExist(input.ShopID)
	if err != nil {
		return err
	}
	if !isExist {
		return errors.New(messages.DataNotFound)
	}

	body := viewmodel.BooksTransactionVm{
		ID:          ID,
		ShopID:      input.ShopID,
		DebtTotal:   input.DebtTotal,
		CreditTotal: input.CreditTotal,
		UpdatedAt:   now.Format(time.RFC3339),
	}
	_, err = model.Edit(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc BooksTransactionUseCase) Add(input *request.BooksTransactionRequest, userID string) (err error) {
	model := actions.NewBooksTransactionModel(uc.DB)
	now := time.Now().UTC()

	body := viewmodel.BooksTransactionVm{
		ShopID:      input.ShopID,
		DebtTotal:   input.DebtTotal,
		CreditTotal: input.CreditTotal,
		CreatedAt:   now.Format(time.RFC3339),
		UpdatedAt:   now.Format(time.RFC3339),
	}
	_, err = model.Add(body, userID, nil)
	if err != nil {
		return err
	}

	return nil
}

func (uc BooksTransactionUseCase) Delete(ID string) (err error) {
	fmt.Println(ID)
	model := actions.NewBooksTransactionModel(uc.DB)
	now := time.Now().UTC()

	isExist, err := uc.IsShopIDExist(ID)
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

func (uc BooksTransactionUseCase) DeleteByUser(userID string, tx *sql.Tx) (err error) {
	model := actions.NewBooksTransactionModel(uc.DB)
	now := time.Now().UTC()

	err = model.DeleteByShop(userID, now.Format(time.RFC3339), now.Format(time.RFC3339), tx)
	if err != nil {
		return err
	}

	return nil
}

func (uc BooksTransactionUseCase) IsShopIDExist(shopID string) (res bool, err error) {
	model := actions.NewBooksTransactionModel(uc.DB)
	count, err := model.CountByShop(shopID)
	if err != nil {
		return res, err
	}

	return count > 0, err
}

func (uc BooksTransactionUseCase) IsShopExist(shopID string) (res bool, err error, data []models.BooksTransaction) {
	model := actions.NewBooksTransactionModel(uc.DB)
	count, err := model.CountByShop(shopID)
	if err != nil {
		return res, err, data
	}
	data, err = model.BrowseByShop(shopID)

	if err != nil {
		return res, err, data
	}

	return count > 0, err, data
}
