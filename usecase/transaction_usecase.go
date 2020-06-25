package usecase

import (
	"bukuduit-go/db/repositories/actions"
	"bukuduit-go/helpers/messages"
	request "bukuduit-go/server/requests"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type TransactionUseCase struct {
	*UcContract
}

func (uc TransactionUseCase) BrowseByCustomer(customerID string) (res []viewmodel.TransactionVm, err error) {
	model := actions.NewTransactionModel(uc.DB)
	Transactions, err := model.BrowseByCustomer(customerID)
	if err != nil {
		return res, err
	}

	for _, Transaction := range Transactions {
		res = append(res, viewmodel.TransactionVm{
			ID:               Transaction.ID,
			Customer_Id:      Transaction.Customer_Id,
			Amount:           fmt.Sprint(Transaction.Amount),
			Description:      Transaction.Description,
			Image:            Transaction.Image,
			Type:             Transaction.Type,
			Transaction_Date: Transaction.Transaction_Date.String,
			Created_at:       Transaction.Created_at,
			Update_at:        Transaction.Update_at.String,
			Deleted_at:       Transaction.Deleted_at.String,
		})
	}

	return res, err
}

func (uc TransactionUseCase) Read(ID string) (res viewmodel.TransactionVm, err error) {
	model := actions.NewTransactionModel(uc.DB)
	Transaction, err := model.Read(ID)
	if err != nil {
		return res, err
	}

	res = viewmodel.TransactionVm{
		ID:               Transaction.ID,
		Customer_Id:      Transaction.Customer_Id,
		Amount:           fmt.Sprint(Transaction.Amount),
		Description:      Transaction.Description,
		Image:            Transaction.Image,
		Type:             Transaction.Type,
		Transaction_Date: Transaction.Transaction_Date.String,
		Created_at:       Transaction.Created_at,
		Update_at:        Transaction.Update_at.String,
		Deleted_at:       Transaction.Deleted_at.String,
	}

	return res, err
}

func (uc TransactionUseCase) Edit(input *request.TransactionRequest, ID string) (err error) {
	model := actions.NewTransactionModel(uc.DB)
	now := time.Now().UTC()
	isExist, err := uc.IsTransactionExist(ID)
	if err != nil {
		return err
	}
	if !isExist {
		return errors.New(messages.DataNotFound)
	}

	body := viewmodel.TransactionVm{
		ID:               ID,
		Customer_Id:      input.Customer_Id,
		Amount:           input.Amount,
		Description:      input.Description,
		Image:            input.Image,
		Type:             input.Type,
		Transaction_Date: input.Transaction_Date,
		Update_at:        now.Format(time.RFC3339),
	}
	_, err = model.Edit(body)
	if err != nil {
		return err
	}

	return nil
}

// func (uc TransactionUseCase) Add(input *request.TransactionRequest, customerID string) (err error) {
// 	model := actions.NewTransactionModel(uc.DB)
// 	now := time.Now().UTC()

// 	body := viewmodel.TransactionVm{
// 		Customer_Id:      input.Customer_Id,
// 		Amount:           input.Amount,
// 		Description:      input.Description,
// 		Image:            input.Image,
// 		Type:             input.Type,
// 		Transaction_Date: input.Transaction_Date,
// 		Update_at:        now.Format(time.RFC3339),
// 		Created_at:       now.Format(time.RFC3339),
// 	}
// 	_, err = model.Add(body)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (uc TransactionUseCase) Delete(ID string) (err error) {
	fmt.Println(ID)
	model := actions.NewTransactionModel(uc.DB)
	now := time.Now().UTC()

	isExist, err := uc.IsTransactionExist(ID)
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

func (uc TransactionUseCase) DebtPayment(CustomerID, DebtType string, UserCustomerDebt, amount int) error {
	TransactionModel := actions.NewTransactionModel(uc.DB)
	userCustomerUc := UserCustomerUseCase{UcContract: uc.UcContract}
	now := time.Now().UTC()
	Transaction, err := uc.DB.Begin()
	if err != nil {
		return err
	}
	TransactionBody := viewmodel.TransactionVm{
		Customer_Id:      CustomerID,
		Amount:           strconv.Itoa(amount),
		Type:             DebtType,
		Transaction_Date: now.Format(time.RFC3339),
		Update_at:        now.Format(time.RFC3339),
		Created_at:       now.Format(time.RFC3339),
	}

	if DebtType == "pay" {
		UserCustomerDebt = UserCustomerDebt - amount
	} else {
		UserCustomerDebt = UserCustomerDebt + amount
	}

	_, errr := TransactionModel.Add(TransactionBody, Transaction)
	if err != nil {
		Transaction.Rollback()
		return err
	}

	eror := userCustomerUc.EditDebt(CustomerID, int32(UserCustomerDebt), Transaction)
	if eror != nil {
		Transaction.Rollback()
		return eror
	}
	Transaction.Commit()
	return errr
}

func (uc TransactionUseCase) DeleteByCustomer(customerID string, tx *sql.Tx) (err error) {
	model := actions.NewTransactionModel(uc.DB)
	now := time.Now().UTC()

	err = model.DeleteByCustomer(customerID, now.Format(time.RFC3339), now.Format(time.RFC3339), tx)
	if err != nil {
		return err
	}

	return nil
}

func (uc TransactionUseCase) IsTransactionExist(ID string) (res bool, err error) {
	model := actions.NewTransactionModel(uc.DB)
	count, err := model.CountByPk(ID)
	if err != nil {
		return res, err
	}

	return count > 0, err
}
