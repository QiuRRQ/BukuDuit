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
	if err != nil {
		fmt.Println(1)
		return res, err
	}

	booksDebtUC := BooksDebtUseCase{UcContract: uc.UcContract}

	var debtBooks viewmodel.BooksDebtVm
	debtBooks, err = booksDebtUC.BrowseByUser(Transactions[0].ReferenceID, enums.Nunggak)
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
		TotalCredit: debtBooks.CreditTotal,
		TotalDebit:  debtBooks.DebtTotal,
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

	booksDebtUC := BooksDebtUseCase{UcContract: uc.UcContract}

	var debtBooks viewmodel.BooksDebtVm
	debtBooks, err = booksDebtUC.BrowseByUser(customerID, enums.Nunggak)
	if err != nil {
		return res, err
	}

	var transactionDate []viewmodel.DebtList
	var transactionDetails []viewmodel.Detail

	for i := 0; i < len(Transactions); i++ {

		tempDate, err := time.Parse(time.RFC3339, Transactions[i].TransactionDate.String)
		if err != nil {
			fmt.Println(err.Error())
		}

		var nextDate time.Time
		if i < len(Transactions)-1 {
			nextDate, err = time.Parse(time.RFC3339, Transactions[i+1].TransactionDate.String)
			if err != nil {
				fmt.Println(err.Error())
			}
			if tempDate == nextDate {
				transactionDetails = append(transactionDetails, viewmodel.Detail{
					ID:          Transactions[i].ID,
					Description: Transactions[i].Description.String,
					Amount:      Transactions[i].Amount.Int32,
					Type:        Transactions[i].Type,
				})

			} else {
				transactionDetails = append(transactionDetails, viewmodel.Detail{
					ID:          Transactions[i].ID,
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
				ID:          Transactions[i].ID,
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
		ReferenceID: Transactions[0].ReferenceID,
		Name:        Transactions[0].Name,
		TotalCredit: debtBooks.CreditTotal,
		TotalDebit:  debtBooks.DebtTotal,
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
	// userCustomerUc := UserCustomerUseCase{UcContract: uc.UcContract}
	booksTransUC := BooksTransactionUseCase{UcContract: uc.UcContract}
	now := time.Now().UTC()

	// customerData, err := userCustomerUc.Read(input.CustomerID)
	if err != nil {
		return err
	}

	transaction, err := uc.DB.Begin()
	if err != nil {
		return err
	}

	var debtAmount int
	var creditAmount int
	var status string
	var booksID string
	status = enums.Nunggak
	//check if fcustomer already exist in books debt
	debtExist, err, data := booksTransUC.IsShopExist(input.ShopID)
	if err != nil {
		return err
	}

	if debtExist {
		//edit booksDebt, status akan terus nunggak baik itu user yang hutang atau customer yang hutang.
		books, err := booksTransUC.Read(data[0].ID, status)

		if err != nil {
			return err
		}

		if input.TransactionType == enums.Debet {
			debtAmount = int(input.Amount) - int(books.DebtTotal)
			creditAmount = books.CreditTotal
			fmt.Println(debtAmount)
			fmt.Println(creditAmount)
		} else {
			creditAmount = int(input.Amount) + int(books.CreditTotal)
			debtAmount = books.DebtTotal
			fmt.Println(debtAmount)
			fmt.Println(creditAmount)
		}

		booksInput := request.BooksTransactionRequest{
			ID:          books.ID,
			ShopID:      books.ShopID,
			DebtTotal:   debtAmount,
			CreditTotal: creditAmount,
			CreatedAt:   now.Format(time.RFC3339),
			UpdatedAt:   now.Format(time.RFC3339),
		}
		err = booksTransUC.Edit(booksInput, books.ID)
		if err != nil {
			fmt.Println(4)
			transaction.Rollback()
			return err
		}

	} else {
		fmt.Println("create new")
		//for adding new debt so adding on books debt
		if input.TransactionType == enums.Debet {
			debtAmount = debtAmount - int(input.Amount)
		} else {
			creditAmount = creditAmount + int(input.Amount)
		}

		// booksInput := request.BooksTransactionRequest{
		// 	ShopID:      input.ShopID,
		// 	DebtTotal:   debtAmount,
		// 	CreditTotal: creditAmount,
		// 	CreatedAt:   now.Format(time.RFC3339),
		// 	UpdatedAt:   now.Format(time.RFC3339),
		// }
		// booksID, err = booksTransUC.Add(booksInput, input.CustomerID, transaction)
		if err != nil {
			fmt.Println(5)
			transaction.Rollback()
			return err
		}
	}

	TransactionBody := viewmodel.TransactionVm{
		ReferenceID:     input.ReferenceID,
		ShopID:          input.ShopID,
		Amount:          input.Amount,
		Description:     input.Description,
		Type:            input.TransactionType,
		CustomerID:      input.CustomerID,
		TransactionDate: input.TransactionDate,
		BooksDebtID:     booksID,
		UpdatedAt:       now.Format(time.RFC3339),
		CreatedAt:       now.Format(time.RFC3339),
	}

	_, err = model.Add(TransactionBody, transaction)
	if err != nil {
		fmt.Println(1)
		transaction.Rollback()
		return err
	}

	transaction.Commit()

	return nil
}

//edit debt customer
func (uc TransactionUseCase) EditDebt(input request.TransactionRequest) (err error) {
	model := actions.NewTransactionModel(uc.DB)
	userCustomerUc := UserCustomerUseCase{UcContract: uc.UcContract}
	booksDebtUC := BooksDebtUseCase{UcContract: uc.UcContract}
	now := time.Now().UTC()

	customerData, err := userCustomerUc.Read(input.ReferenceID)
	if err != nil {
		return err
	}

	transaction, err := uc.DB.Begin()
	if err != nil {
		return err
	}

	var debtAmount int
	var creditAmount int
	var status string
	var books viewmodel.BooksDebtVm
	var getTrans models.Transactions
	status = enums.Nunggak
	//check if fcustomer already exist in books debt
	debtExist, err, data := booksDebtUC.IsDebtCustomerExist(customerData.ID, enums.Nunggak)
	if err != nil {
		return err
	}

	if debtExist {
		//edit booksDebt, status akan terus nunggak baik itu user yang hutang atau customer yang hutang.
		books, err := booksDebtUC.Read(data[0].ID, status)
		if err != nil {
			return err
		}

		getTrans, err := model.Read(input.ID)
		if err != nil {
			return err
		}

		if input.TransactionType == enums.Debet {
			if int(input.Amount) > int(getTrans.Amount.Int32) {
				debtAmount = int(books.DebtTotal) + (int(input.Amount) - int(getTrans.Amount.Int32))
				creditAmount = books.CreditTotal

			} else {
				debtAmount = int(books.DebtTotal) - (int(getTrans.Amount.Int32) - int(input.Amount))
				creditAmount = books.CreditTotal
			}

			fmt.Println(debtAmount)
			fmt.Println(creditAmount)
		} else {
			if int(input.Amount) > int(getTrans.Amount.Int32) {
				creditAmount = int(books.CreditTotal) + (int(input.Amount) - int(getTrans.Amount.Int32))
				debtAmount = books.DebtTotal

			} else {
				creditAmount = int(books.CreditTotal) - (int(getTrans.Amount.Int32) - int(input.Amount))
				debtAmount = books.DebtTotal
			}

			fmt.Println(debtAmount)
			fmt.Println(creditAmount)
		}

		booksInput := request.BooksDebtRequest{
			CustomerID:     customerData.ID,
			SubmissionDate: books.SubmissionDate,
			DebtTotal:      debtAmount,
			CreditTotal:    creditAmount,
			Status:         books.Status,
			CreatedAt:      books.CreatedAt,
			UpdatedAt:      now.Format(time.RFC3339),
		}
		err = booksDebtUC.Edit(booksInput, books.ID, transaction)
		if err != nil {
			transaction.Rollback()
			return err
		}

	}

	TransactionBody := viewmodel.TransactionVm{
		ID:              input.ID,
		ReferenceID:     input.ReferenceID,
		ShopID:          input.ShopID,
		Amount:          input.Amount,
		Description:     input.Description,
		Type:            input.TransactionType,
		CustomerID:      input.CustomerID,
		TransactionDate: input.TransactionDate,
		BooksDebtID:     books.ID,
		UpdatedAt:       now.Format(time.RFC3339),
		CreatedAt:       getTrans.CreatedAt,
	}

	_, err = model.Edit(TransactionBody, transaction)
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()

	return nil
}

//ini untuk add customer debt
func (uc TransactionUseCase) DebtPayment(input request.TransactionRequest) (err error) {
	model := actions.NewTransactionModel(uc.DB)
	userCustomerUc := UserCustomerUseCase{UcContract: uc.UcContract}
	booksDebtUC := BooksDebtUseCase{UcContract: uc.UcContract}
	now := time.Now().UTC()

	customerData, err := userCustomerUc.Read(input.ReferenceID)
	if err != nil {
		return err
	}

	transaction, err := uc.DB.Begin()
	if err != nil {
		return err
	}

	var debtAmount int
	var creditAmount int
	var status string
	var booksID string
	status = enums.Nunggak
	//check if fcustomer already exist in books debt
	debtExist, err, data := booksDebtUC.IsDebtCustomerExist(customerData.ID, enums.Nunggak)
	if err != nil {
		return err
	}

	if debtExist {
		//edit booksDebt, status akan terus nunggak baik itu user yang hutang atau customer yang hutang.
		books, err := booksDebtUC.Read(data[0].ID, status)

		if err != nil {
			return err
		}

		if input.TransactionType == enums.Debet {
			debtAmount = int(input.Amount) - int(books.DebtTotal)
			creditAmount = books.CreditTotal
			fmt.Println(debtAmount)
			fmt.Println(creditAmount)
		} else {
			creditAmount = int(input.Amount) + int(books.CreditTotal)
			debtAmount = books.DebtTotal
			fmt.Println(debtAmount)
			fmt.Println(creditAmount)
		}

		booksInput := request.BooksDebtRequest{
			CustomerID:     customerData.ID,
			SubmissionDate: now.Format(time.RFC3339),
			DebtTotal:      debtAmount,
			CreditTotal:    creditAmount,
			Status:         status,
			CreatedAt:      now.Format(time.RFC3339),
			UpdatedAt:      now.Format(time.RFC3339),
		}
		err = booksDebtUC.Edit(booksInput, books.ID, transaction)
		if err != nil {
			fmt.Println(4)
			transaction.Rollback()
			return err
		}

	} else {
		fmt.Println("create new")
		//for adding new debt so adding on books debt
		if input.TransactionType == enums.Debet {
			debtAmount = debtAmount - int(input.Amount)
		} else {
			creditAmount = creditAmount + int(input.Amount)
		}

		booksInput := request.BooksDebtRequest{
			CustomerID:     customerData.ID,
			SubmissionDate: now.Format(time.RFC3339),
			DebtTotal:      debtAmount,
			CreditTotal:    creditAmount,
			Status:         status,
			CreatedAt:      now.Format(time.RFC3339),
			UpdatedAt:      now.Format(time.RFC3339),
		}
		booksID, err = booksDebtUC.Add(booksInput, input.CustomerID, transaction)
		if err != nil {
			fmt.Println(5)
			transaction.Rollback()
			return err
		}
	}

	TransactionBody := viewmodel.TransactionVm{
		ReferenceID:     input.ReferenceID,
		ShopID:          input.ShopID,
		Amount:          input.Amount,
		Description:     input.Description,
		Type:            input.TransactionType,
		CustomerID:      input.CustomerID,
		TransactionDate: input.TransactionDate,
		BooksDebtID:     booksID,
		UpdatedAt:       now.Format(time.RFC3339),
		CreatedAt:       now.Format(time.RFC3339),
	}

	_, err = model.Add(TransactionBody, transaction)
	if err != nil {
		fmt.Println(1)
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
