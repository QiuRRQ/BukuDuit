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

//list transaksi
func (uc TransactionUseCase) TransactionList(shopID string) (res viewmodel.TransactionListVm, err error) {
	model := actions.NewTransactionModel(uc.DB)
	Transactions, err := model.TransactionBrowsByShop(shopID)
	resultCount, err := model.CountDistinctBy("shop_id", shopID)
	fmt.Println(resultCount)
	if err != nil {
		return res, err
	}

	var debtTotal int
	var creditTotal int

	var dateCreditAmount int
	var dateDebetAmount int

	var debtDate []viewmodel.DataList
	var debtDetails []viewmodel.DataDetails

	for i := 0; i < len(Transactions); i++ {

		tempDate, err := time.Parse(time.RFC3339, Transactions[i].TransactionDate.String)
		if err != nil {
			fmt.Println(err.Error())
		}
		if Transactions[i].Type == enums.Debet {
			debtTotal = debtTotal + int(Transactions[i].Amount.Int32)
		} else {
			creditTotal = creditTotal + int(Transactions[i].Amount.Int32)
		}

		if Transactions[i].Type == enums.Debet {
			dateDebetAmount = dateDebetAmount + int(Transactions[i].Amount.Int32)
		} else {
			dateCreditAmount = dateCreditAmount + int(Transactions[i].Amount.Int32)
		}

		var nextDate time.Time
		if i < len(Transactions)-1 {
			nextDate, err = time.Parse(time.RFC3339, Transactions[i+1].TransactionDate.String)
			if err != nil {
				fmt.Println(err.Error())
			}
			if tempDate == nextDate {

				debtDetails = append(debtDetails, viewmodel.DataDetails{
					Name:        Transactions[i].Name,
					Description: Transactions[i].Description.String,
					Amount:      Transactions[i].Amount.Int32,
					Type:        Transactions[i].Type,
				})

			} else {
				debtDetails = append(debtDetails, viewmodel.DataDetails{
					Name:        Transactions[i].Name,
					Description: Transactions[i].Description.String,
					Amount:      Transactions[i].Amount.Int32,
					Type:        Transactions[i].Type,
				})
				debtDate = append(debtDate, viewmodel.DataList{
					TransactionDate:  tempDate.String(),
					DateAmountCredit: dateCreditAmount,
					DateAmountDebet:  dateDebetAmount,
					Details:          debtDetails,
				})

				debtDetails = nil
				dateDebetAmount = 0
				dateCreditAmount = 0
				tempDate = nextDate
			}
		} else {
			debtDetails = append(debtDetails, viewmodel.DataDetails{
				Name:        Transactions[i].Name,
				Description: Transactions[i].Description.String,
				Amount:      Transactions[i].Amount.Int32,
				Type:        Transactions[i].Type,
			})
			debtDate = append(debtDate, viewmodel.DataList{
				TransactionDate:  tempDate.String(),
				DateAmountCredit: dateCreditAmount,
				DateAmountDebet:  dateDebetAmount,
				Details:          debtDetails,
			})

			debtDetails = nil
			dateDebetAmount = 0
			dateCreditAmount = 0
			tempDate = nextDate
		}

	}

	for i := 0; i < resultCount; i++ {
		res = viewmodel.TransactionListVm{
			ID:          Transactions[i].ID,
			ReferenceID: Transactions[i].ReferenceID,
			TotalCredit: creditTotal,
			TotalDebit:  debtTotal,
			ListData:    debtDate,
			CreatedAt:   Transactions[i].CreatedAt,
			UpdatedAt:   Transactions[i].UpdatedAt.String,
			DeletedAt:   Transactions[i].DeletedAt.String,
		}
	}

	return res, err
}

//laporan hutang
func (uc TransactionUseCase) BrowseByShop(shopID string) (res viewmodel.ReportHutangVm, err error) {
	model := actions.NewTransactionModel(uc.DB)
	Transactions, err := model.BrowseByShop(shopID)
	resultCount, err := model.CountDistinctBy("shop_id", shopID)
	fmt.Println(resultCount)
	if err != nil {
		return res, err
	}

	var debtTotal int
	var creditTotal int

	var debtDate []viewmodel.DebtReport
	var debtDetails []viewmodel.DebtDetail

	for i := 0; i < len(Transactions); i++ {

		tempDate, err := time.Parse(time.RFC3339, Transactions[i].TransactionDate.String)
		if err != nil {
			fmt.Println(err.Error())
		}
		if Transactions[i].Type == enums.Debet {
			debtTotal = debtTotal + int(Transactions[i].Amount.Int32)
		} else {
			creditTotal = creditTotal + int(Transactions[i].Amount.Int32)
		}

		var nextDate time.Time
		if i < len(Transactions)-1 {
			nextDate, err = time.Parse(time.RFC3339, Transactions[i+1].TransactionDate.String)
			if err != nil {
				fmt.Println(err.Error())
			}
			if tempDate == nextDate {
				debtDetails = append(debtDetails, viewmodel.DebtDetail{
					Name:        Transactions[i].Name,
					Description: Transactions[i].Description.String,
					Amount:      Transactions[i].Amount.Int32,
					Type:        Transactions[i].Type,
				})

			} else {
				debtDetails = append(debtDetails, viewmodel.DebtDetail{
					Name:        Transactions[i].Name,
					Description: Transactions[i].Description.String,
					Amount:      Transactions[i].Amount.Int32,
					Type:        Transactions[i].Type,
				})
				debtDate = append(debtDate, viewmodel.DebtReport{
					TransactionDate: tempDate.String(),
					Details:         debtDetails,
				})

				debtDetails = nil
				tempDate = nextDate
			}
		} else {
			debtDetails = append(debtDetails, viewmodel.DebtDetail{
				Name:        Transactions[i].Name,
				Description: Transactions[i].Description.String,
				Amount:      Transactions[i].Amount.Int32,
				Type:        Transactions[i].Type,
			})
			debtDate = append(debtDate, viewmodel.DebtReport{
				TransactionDate: tempDate.String(),
				Details:         debtDetails,
			})

			debtDetails = nil
			tempDate = nextDate
		}

	}

	res = viewmodel.ReportHutangVm{
		ID:          Transactions[0].ID,
		ReferenceID: Transactions[0].ReferenceID,
		TotalCredit: creditTotal,
		TotalDebit:  debtTotal,
		ListData:    debtDate,
		CreatedAt:   Transactions[0].CreatedAt,
		UpdatedAt:   Transactions[0].UpdatedAt.String,
		DeletedAt:   Transactions[0].DeletedAt.String,
	}

	return res, err

}

func (uc TransactionUseCase) BrowseByCustomer(customerID string) (res viewmodel.DetailsHutangVm, err error) {
	model := actions.NewTransactionModel(uc.DB)
	Transactions, err := model.BrowseByCustomer(customerID) //only use it for details

	if err != nil {
		return res, err
	}

	var debtTotal int
	var creditTotal int

	var transactionDate []viewmodel.DebtList
	var transactionDetails []viewmodel.Detail

	for i := 0; i < len(Transactions); i++ {

		tempDate, err := time.Parse(time.RFC3339, Transactions[i].TransactionDate.String)
		if err != nil {
			fmt.Println(err.Error())
		}
		if Transactions[i].Type == enums.Debet {
			debtTotal = debtTotal + int(Transactions[i].Amount.Int32)
		} else {
			creditTotal = creditTotal + int(Transactions[i].Amount.Int32)
		}

		var nextDate time.Time
		if i < len(Transactions)-1 {
			nextDate, err = time.Parse(time.RFC3339, Transactions[i+1].TransactionDate.String)
			if err != nil {
				fmt.Println(err.Error())
			}
			if tempDate == nextDate {
				transactionDetails = append(transactionDetails, viewmodel.Detail{
					Description: Transactions[i].Description.String,
					Amount:      Transactions[i].Amount.Int32,
					Type:        Transactions[i].Type,
				})

			} else {
				transactionDetails = append(transactionDetails, viewmodel.Detail{
					Description: Transactions[i].Description.String,
					Amount:      Transactions[i].Amount.Int32,
					Type:        Transactions[i].Type,
				})
				transactionDate = append(transactionDate, viewmodel.DebtList{
					TransactionDate: Transactions[i].TransactionDate.String,
					Details:         transactionDetails,
				})

				transactionDetails = nil
				tempDate = nextDate
			}
		} else {
			transactionDetails = append(transactionDetails, viewmodel.Detail{
				Description: Transactions[i].Description.String,
				Amount:      Transactions[i].Amount.Int32,
				Type:        Transactions[i].Type,
			})
			transactionDate = append(transactionDate, viewmodel.DebtList{
				TransactionDate: Transactions[i].TransactionDate.String,
				Details:         transactionDetails,
			})

			transactionDetails = nil
			tempDate = nextDate
		}

	}

	res = viewmodel.DetailsHutangVm{
		ID:          Transactions[0].ID,
		ReferenceID: Transactions[0].ReferenceID,
		Name:        Transactions[0].Name,
		TotalCredit: creditTotal,
		TotalDebit:  debtTotal,
		ListData:    transactionDate,
		CreatedAt:   Transactions[0].CreatedAt,
		UpdatedAt:   Transactions[0].UpdatedAt.String,
		DeletedAt:   Transactions[0].DeletedAt.String,
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
		Description:     Transaction.Description.String,
		Image:           Transaction.Image.String,
		Type:            Transaction.Type,
		TransactionDate: Transaction.TransactionDate.String,
		CreatedAt:       Transaction.CreatedAt,
		UpdatedAt:       Transaction.UpdatedAt.String,
		DeletedAt:       Transaction.DeletedAt.String,
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
		ID:          ID,
		ReferenceID: input.ReferenceID,
		Amount:      input.Amount,
		//Description:     input.Description,
		Type:            input.TransactionType,
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

//karena transaksi itu bukan hutang jadi gk usah edit customer debt
func (uc TransactionUseCase) AddTransaksi(input request.TransactionRequest) (err error) {
	model := actions.NewTransactionModel(uc.DB)
	userCustomerUc := UserCustomerUseCase{UcContract: uc.UcContract}
	now := time.Now().UTC()

	customerData, err := userCustomerUc.Read(input.CustomerID)
	if err != nil {
		fmt.Println(1)
		return err
	}
	userDebtAmount := customerData.Debt
	TransactionBody := viewmodel.TransactionVm{
		ReferenceID:     input.ReferenceID,
		ShopID:          input.ShopID,
		Amount:          input.Amount,
		Description:     input.Description,
		Type:            input.TransactionType,
		CustomerID:      input.CustomerID,
		TransactionDate: input.TransactionDate,
		UpdatedAt:       now.Format(time.RFC3339),
		CreatedAt:       now.Format(time.RFC3339),
	}

	if input.TransactionType == enums.Debet {
		userDebtAmount = userDebtAmount - input.Amount
	} else {
		userDebtAmount = userDebtAmount + input.Amount
	}

	transaction, err := uc.DB.Begin()
	if err != nil {
		fmt.Println(2)
		return err
	}
	_, err = model.Add(TransactionBody, transaction)
	if err != nil {
		fmt.Println(3)
		transaction.Rollback()
		return err
	}

	transaction.Commit()

	return nil
}

//ini untuk pembayaran customer debt
func (uc TransactionUseCase) DebtPayment(input request.TransactionRequest) (err error) {
	model := actions.NewTransactionModel(uc.DB)
	userCustomerUc := UserCustomerUseCase{UcContract: uc.UcContract}
	now := time.Now().UTC()

	customerData, err := userCustomerUc.Read(input.ReferenceID)
	if err != nil {
		return err
	}
	userDebtAmount := customerData.Debt
	TransactionBody := viewmodel.TransactionVm{
		ReferenceID:     input.ReferenceID,
		ShopID:          input.ShopID,
		Amount:          input.Amount,
		Description:     input.Description,
		Type:            input.TransactionType,
		CustomerID:      input.CustomerID,
		TransactionDate: input.TransactionDate,
		UpdatedAt:       now.Format(time.RFC3339),
		CreatedAt:       now.Format(time.RFC3339),
	}

	if input.TransactionType == enums.Debet {
		userDebtAmount = userDebtAmount - input.Amount
	} else {
		userDebtAmount = userDebtAmount + input.Amount
	}

	transaction, err := uc.DB.Begin()
	if err != nil {
		return err
	}
	_, err = model.Add(TransactionBody, transaction)
	if err != nil {
		fmt.Println(1)
		transaction.Rollback()
		return err
	}

	err = userCustomerUc.EditDebt(input.ReferenceID, userDebtAmount, transaction)
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
