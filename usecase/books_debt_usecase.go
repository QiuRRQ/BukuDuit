package usecase

import (
	"bukuduit-go/db/models"
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

type BooksDebtUseCase struct {
	*UcContract
}

func (uc BooksDebtUseCase) BrowseByUser(customerID, status string) (res viewmodel.BooksDebtVm, err error) {
	model := actions.NewBooksDebtModel(uc.DB)
	books, err := model.BrowseByCustomer(customerID, status)
	if err != nil {
		return res, err
	}

	for _, book := range books {
		res = viewmodel.BooksDebtVm{
			ID:             book.ID,
			CustomerID:     book.CustomerID,
			SubmissionDate: book.SubmissionDate,
			BillDate:       book.BillDate.String,
			DebtTotal:      book.DebtTotal,
			CreditTotal:    book.CreditTotal,
			Status:         book.Status.String,
			CreatedAt:      book.CreatedAt,
			UpdatedAt:      book.UpdatedAt.String,
			DeletedAt:      book.DeletedAt.String,
		}
	}

	return res, err
}

//read data
func (uc BooksDebtUseCase) Read(ID, status string) (res viewmodel.BooksDebtVm, err error) {
	model := actions.NewBooksDebtModel(uc.DB)
	userCustomerUC := BooksDebtUseCase{UcContract: uc.UcContract}

	fmt.Println(userCustomerUC)

	books, err := model.Read(ID)
	if err != nil {
		return res, err
	}

	res = viewmodel.BooksDebtVm{
		ID:             books.ID,
		CustomerID:     books.CustomerID,
		SubmissionDate: books.SubmissionDate,
		BillDate:       books.BillDate.String,
		DebtTotal:      books.DebtTotal,
		CreditTotal:    books.CreditTotal,
		Status:         books.Status.String,
		CreatedAt:      books.CreatedAt,
		UpdatedAt:      books.UpdatedAt.String,
		DeletedAt:      books.UpdatedAt.String,
	}

	return res, err
}

func (uc BooksDebtUseCase) Edit(input request.BooksDebtRequest, ID string, tx *sql.Tx) (err error) {
	model := actions.NewBooksDebtModel(uc.DB)
	now := time.Now().UTC()
	isExist, err := uc.IsDebtExist(ID, enums.Nunggak)
	if err != nil {
		return err
	}
	if !isExist {
		return errors.New(messages.DataNotFound)
	}

	body := viewmodel.BooksDebtVm{
		ID:             ID,
		CustomerID:     input.CustomerID,
		SubmissionDate: input.SubmissionDate,
		BillDate:       input.BillDate,
		DebtTotal:      input.DebtTotal,
		CreditTotal:    input.CreditTotal,
		Status:         input.Status,
		UpdatedAt:      now.Format(time.RFC3339),
	}
	_, err = model.Edit(body, tx)
	if err != nil {
		fmt.Println(5)
		return err
	}

	return nil
}

func (uc BooksDebtUseCase) Add(input request.BooksDebtRequest, userID string, tx *sql.Tx) (res string, err error) {
	model := actions.NewBooksDebtModel(uc.DB)
	now := time.Now().UTC()

	body := viewmodel.BooksDebtVm{
		CustomerID:     input.CustomerID,
		SubmissionDate: input.SubmissionDate,
		BillDate:       input.BillDate,
		DebtTotal:      input.DebtTotal,
		CreditTotal:    input.CreditTotal,
		Status:         input.Status,
		CreatedAt:      now.Format(time.RFC3339),
		UpdatedAt:      now.Format(time.RFC3339),
	}
	res, err = model.Add(body, tx)
	if err != nil {
		fmt.Println("here")
		return res, err
	}

	return res, nil
}

func (uc BooksDebtUseCase) Register(userID, bookName, createdAt, updatedAt string, tx *sql.Tx) (err error) {
	model := actions.NewBooksDebtModel(uc.DB)
	body := viewmodel.BooksDebtVm{
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	_, err = model.Add(body, tx)
	if err != nil {
		return err
	}

	return nil
}

func (uc BooksDebtUseCase) Delete(ID string) (err error) {
	fmt.Println(ID)
	model := actions.NewBooksDebtModel(uc.DB)
	now := time.Now().UTC()

	isExist, err := uc.IsDebtExist(ID, enums.Nunggak)
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

func (uc BooksDebtUseCase) DeleteByUser(userID string, tx *sql.Tx) (err error) {
	model := actions.NewBooksDebtModel(uc.DB)
	now := time.Now().UTC()

	err = model.DeleteByCustomer(userID, now.Format(time.RFC3339), now.Format(time.RFC3339), tx)
	if err != nil {
		return err
	}

	return nil
}

func (uc BooksDebtUseCase) IsDebtExist(ID, status string) (res bool, err error) {
	model := actions.NewBooksDebtModel(uc.DB)
	count, err := model.CountByPk(ID, status)
	if err != nil {
		return res, err
	}

	if err != nil {
		return res, err
	}

	return count > 0, err
}

func (uc BooksDebtUseCase) IsDebtCustomerExist(customerID, status string) (res bool, err error, data []models.BooksDebt) {
	model := actions.NewBooksDebtModel(uc.DB)
	count, err := model.CountByCustomer(customerID, status)
	if err != nil {
		return res, err, data
	}
	data, err = model.BrowseByCustomer(customerID, status)

	if err != nil {
		return res, err, data
	}

	return count > 0, err, data
}
