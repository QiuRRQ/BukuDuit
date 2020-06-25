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
			ID:              Transaction.ID,
			ReferenceID:     Transaction.ReferenceID,
			Amount:          Transaction.Amount.Int32,
			Description:     Transaction.Description,
			Image:           Transaction.Image,
			Type:            Transaction.Type,
			TransactionDate: Transaction.TransactionDate.String,
			CreatedAt:       Transaction.CreatedAt,
			UpdatedAt:       Transaction.UpdatedAt.String,
			DeletedAt:      Transaction.DeletedAt.String,
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
		ID:              Transaction.ID,
		ReferenceID:     Transaction.ReferenceID,
		Amount:          Transaction.Amount.Int32,
		Description:     Transaction.Description,
		Image:           Transaction.Image,
		Type:            Transaction.Type,
		TransactionDate: Transaction.TransactionDate.String,
		CreatedAt:       Transaction.CreatedAt,
		UpdatedAt:       Transaction.UpdatedAt.String,
		DeletedAt:      Transaction.DeletedAt.String,
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
		ID:              ID,
		ReferenceID:     input.CustomerID,
		Amount:          input.Amount,
		Description:     input.Description,
		Image:           input.Image,
		Type:            input.Type,
		TransactionDate: input.TransactionDate,
		UpdatedAt:       now.Format(time.RFC3339),
	}
	_, err = model.Edit(body)
	if err != nil {
		return err
	}

	return nil
}

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

func (uc TransactionUseCase) DebtPayment(referenceID, DebtType, shopID string, amount int32) (err error) {
	model := actions.NewTransactionModel(uc.DB)
	userCustomerUc := UserCustomerUseCase{UcContract: uc.UcContract}
	now := time.Now().UTC()

	customerData, err := userCustomerUc.Read(referenceID)
	if err != nil {
		return err
	}
	userDebtAmount := customerData.Debt
	TransactionBody := viewmodel.TransactionVm{
		ReferenceID:     referenceID,
		ShopID:          shopID,
		Amount:          amount,
		Type:            DebtType,
		TransactionDate: now.Format(time.RFC3339),
		UpdatedAt:       now.Format(time.RFC3339),
		CreatedAt:       now.Format(time.RFC3339),
	}

	if DebtType == enums.Debt {
		userDebtAmount = userDebtAmount - amount
	} else {
		userDebtAmount = userDebtAmount + amount
	}

	transaction, err := uc.DB.Begin()
	if err != nil {
		return err
	}
	_, err = model.Add(TransactionBody, transaction)
	if err != nil {
		transaction.Rollback()
		return err
	}

	err = userCustomerUc.EditDebt(referenceID, userDebtAmount, transaction)
	if err != nil {
		transaction.Rollback()
		return err
	}
	transaction.Commit()

	return nil
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
