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

func (uc TransactionUseCase) TransactionReport(shopID, search, name, amount, transDate, startDate, endDate string) (res viewmodel.TransactionListVm, err error) {
	model := actions.NewTransactionModel(uc.DB)
	var filter string
	if startDate != "" && endDate != "" {
		filter = `and (t."transaction_date" BETWEEN '` + startDate + `' and '` + endDate + `')`
	}
	if search != "" { //input nama
		filter = `and uc."full_name" ILIKE '%` + search + `%'` + filter
	}
	//sort hanya dipakai sekali
	if name == "ASC" || name == "asc" {
		filter = filter + ` order by uc."full_name" ` + name
	}
	if name == "DESC" || name == "desc" {
		filter = filter + ` order by uc."full_name" ` + name
	}
	if amount == "ASC" || amount == "asc" {
		filter = filter + ` order by t."amount" ` + amount
	}
	if amount == "DESC" || amount == "desc" {
		filter = filter + ` order by t."amount" ` + amount
	}
	if transDate == "ASC" || transDate == "asc" {
		filter = filter + ` order by t."transaction_date" ` + transDate
	}
	if transDate == "DESC" || transDate == "desc" {
		filter = filter + ` order by t."transaction_date" ` + transDate
	}
	Transactions, err := model.TransactionReport(shopID, filter)
	if err != nil {
		fmt.Println(1)
		return res, err
	}
	resultCount, err := model.CountDistinctBy("shop_id", shopID)
	if err != nil {
		fmt.Println(2)
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
					ID:          Transactions[i].ID,
					ReferenceID: Transactions[i].ReferenceID,
					Name:        Transactions[i].Name,
					Description: Transactions[i].Description.String,
					Amount:      Transactions[i].Amount.Int32,
					Type:        Transactions[i].Type,
				})

			} else {
				debtDetails = append(debtDetails, viewmodel.DataDetails{
					ID:          Transactions[i].ID,
					ReferenceID: Transactions[i].ReferenceID,
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
				ID:          Transactions[i].ID,
				ReferenceID: Transactions[i].ReferenceID,
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
			ShopID:      Transactions[i].IDShop,
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

//list transaksi
func (uc TransactionUseCase) TransactionList(shopID, name, timeFilter string) (res viewmodel.TransactionListVm, err error) {
	model := actions.NewTransactionModel(uc.DB)
	var filter string
	if name != "" {
		filter = `and uc."full_name" ilike '%` + name + `%'`
	}
	Transactions, err := model.TransactionBrowsByShop(shopID, filter)
	if err != nil {
		fmt.Println(1)
		return res, err
	}
	resultCount, err := model.CountDistinctBy("shop_id", shopID)
	if err != nil {
		fmt.Println(2)
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
					ID:          Transactions[i].ID,
					ReferenceID: Transactions[i].ReferenceID,
					Name:        Transactions[i].Name,
					Description: Transactions[i].Description.String,
					Amount:      Transactions[i].Amount.Int32,
					Type:        Transactions[i].Type,
				})

			} else {
				debtDetails = append(debtDetails, viewmodel.DataDetails{
					ID:          Transactions[i].ID,
					ReferenceID: Transactions[i].ReferenceID,
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
				ID:          Transactions[i].ID,
				ReferenceID: Transactions[i].ReferenceID,
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
			ShopID:      Transactions[i].IDShop,
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

func (uc TransactionUseCase) DebtReport(shopID, search, name, amount, transDate, startDate, endDate string) (res viewmodel.ReportHutangVm, err error) {
	model := actions.NewTransactionModel(uc.DB)
	booksDebtUC := BooksDebtUseCase{UcContract: uc.UcContract}

	books, err := booksDebtUC.BrowseByShop(shopID, "")
	if err != nil {
		return res, err
	}

	var debtTotal int
	var creditTotal int
	var customerID string
	var bookDebtID string
	var debtDate []viewmodel.DebtReport
	var debtDetails []viewmodel.DebtDetail

	if len(books) > 0 {
		//variabel total di sini gk berubah selama filter gk diset
		for i, book := range books {
			debtTotal = debtTotal + book.DebtTotal
			creditTotal = creditTotal + book.CreditTotal
			if i != len(books)-1 {
				customerID = customerID + "'" + book.CustomerID + "'" + ","
				bookDebtID = bookDebtID + "'" + book.ID + "'" + ","
			} else {
				customerID = customerID + "'" + book.CustomerID + "'"
				bookDebtID = bookDebtID + "'" + book.ID + "'"
			}
		}

		fmt.Println(bookDebtID)

		var filter string

		if startDate != "" && endDate != "" {
			filter = `and (t."transaction_date" BETWEEN '` + startDate + `' and '` + endDate + `')`
		}
		if search != "" { //input nama
			filter = `and uc."full_name" ILIKE '%` + search + `%'` + filter
		}
		//sort hanya dipakai sekali
		if name == "ASC" || name == "asc" {
			filter = filter + ` order by uc."full_name" ` + name
		}
		if name == "DESC" || name == "desc" {
			filter = filter + ` order by uc."full_name" ` + name
		}
		if amount == "ASC" || amount == "asc" {
			filter = filter + ` order by t."amount" ` + amount
		}
		if amount == "DESC" || amount == "desc" {
			filter = filter + ` order by t."amount" ` + amount
		}
		if transDate == "ASC" || transDate == "asc" {
			filter = filter + ` order by t."transaction_date" ` + transDate
		}
		if transDate == "DESC" || transDate == "desc" {
			filter = filter + ` order by t."transaction_date" ` + transDate
		}

		transactions, err := model.DebtReport(customerID, shopID, bookDebtID, filter)
		if err != nil {
			return res, err
		}

		if filter != "" {
			debtTotal = 0
			creditTotal = 0
		}

		for i := 0; i < len(transactions); i++ {

			if filter != "" {
				if transactions[i].Type == enums.Debet {
					debtTotal = debtTotal + int(transactions[i].Amount.Int32)
				} else {
					creditTotal = creditTotal + int(transactions[i].Amount.Int32)
				}
			}

			tempDate, err := time.Parse(time.RFC3339, transactions[i].TransactionDate.String)
			if err != nil {
				fmt.Println(err.Error())
			}

			var nextDate time.Time
			if i < len(transactions)-1 {
				nextDate, err = time.Parse(time.RFC3339, transactions[i+1].TransactionDate.String)
				if err != nil {
					fmt.Println(err.Error())
				}
				if tempDate == nextDate {
					debtDetails = append(debtDetails, viewmodel.DebtDetail{
						ID:          transactions[i].ID,
						ReferenceID: transactions[i].ReferenceID,
						Name:        transactions[i].Name,
						Description: transactions[i].Description.String,
						Amount:      transactions[i].Amount.Int32,
						Type:        transactions[i].Type,
					})

				} else {
					debtDetails = append(debtDetails, viewmodel.DebtDetail{
						ID:          transactions[i].ID,
						ReferenceID: transactions[i].ReferenceID,
						Name:        transactions[i].Name,
						Description: transactions[i].Description.String,
						Amount:      transactions[i].Amount.Int32,
						Type:        transactions[i].Type,
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
					ID:          transactions[i].ID,
					ReferenceID: transactions[i].ReferenceID,
					Name:        transactions[i].Name,
					Description: transactions[i].Description.String,
					Amount:      transactions[i].Amount.Int32,
					Type:        transactions[i].Type,
				})
				debtDate = append(debtDate, viewmodel.DebtReport{
					TransactionDate: tempDate.String(),
					Details:         debtDetails,
				})

				debtDetails = nil
				tempDate = nextDate
			}
		}
	}

	res = viewmodel.ReportHutangVm{
		TotalCredit: creditTotal,
		TotalDebit:  debtTotal,
		ListData:    debtDate,
	}

	return res, err
}

func (uc TransactionUseCase) BrowseByCustomer(customerID string) (res viewmodel.DetailsHutangVm, err error) {
	model := actions.NewTransactionModel(uc.DB)
	bookDebtUc := BooksDebtUseCase{UcContract: uc.UcContract}
	var transactionDate []viewmodel.DebtList
	var transactionDetails []viewmodel.Detail
	var debtBooks viewmodel.BooksDebtVm

	isBookDebtExist, err := bookDebtUc.IsDebtCustomerExist(customerID, enums.Nunggak)
	if err != nil {
		return res, err
	}

	transactionCount, err := uc.CountBy("reference_id", customerID)
	if err != nil {
		return res, err
	}

	if isBookDebtExist && transactionCount > 0 {
		Transactions, err := model.BrowseByCustomer(customerID) //only use it for details
		if err != nil {
			return res, errors.New(messages.DataNotFound)
		}

		booksDebtUC := BooksDebtUseCase{UcContract: uc.UcContract}

		debtBooks, err = booksDebtUC.BrowseByUser(customerID, enums.Nunggak)
		if err != nil {
			return res, errors.New(messages.DataNotFound)
		}

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
		ShopID:          Transaction.IDShop,
		BooksDebtID:     Transaction.BooksDeptID.String,
		BooksTransID:    Transaction.BooksTransID.String,
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

func (uc TransactionUseCase) DeleteDebt(ID string) (err error) {
	model := actions.NewTransactionModel(uc.DB)
	bookDebtUc := BooksDebtUseCase{UcContract: uc.UcContract}
	now := time.Now().UTC()
	var debtTotal int
	var creditTotal int

	isExist, err := uc.IsTransactionExist(ID)
	if err != nil {

		return err
	}
	if !isExist {

		return errors.New(messages.DataNotFound)
	}

	transactionData, err := uc.Read(ID)
	if err != nil {

		return err
	}

	transaction, err := uc.DB.Begin()
	if err != nil {

		return err
	}

	books, err := bookDebtUc.Read(transactionData.BooksDebtID, "")
	if err != nil {
		transaction.Rollback()
		return err
	}

	if transactionData.Type == enums.Debet {
		debtTotal = books.DebtTotal - int(transactionData.Amount)
		creditTotal = books.CreditTotal
	} else {
		creditTotal = books.CreditTotal - int(transactionData.Amount)
		debtTotal = books.DebtTotal
	}

	reqBooks := request.BooksDebtRequest{
		ID:             books.ID,
		CustomerID:     books.CustomerID,
		SubmissionDate: books.SubmissionDate,
		BillDate:       books.BillDate,
		DebtTotal:      debtTotal,
		CreditTotal:    creditTotal,
		Status:         books.Status,
		CreatedAt:      books.CreatedAt,
		UpdatedAt:      now.Format(time.RFC3339),
	}

	err = bookDebtUc.Edit(reqBooks, transactionData.BooksDebtID, transaction)
	if err != nil {
		fmt.Println(5)
		transaction.Rollback()
		return err
	}

	err = model.Delete(ID, now.Format(time.RFC3339), now.Format(time.RFC3339), transaction)
	if err != nil {
		fmt.Println(6)
		transaction.Rollback()

		return err
	}

	transaction.Commit()

	return nil
}

//karena transaksi itu bukan hutang jadi gk usah edit customer debt
func (uc TransactionUseCase) AddTransaksi(input request.TransactionRequest) (err error) {
	model := actions.NewTransactionModel(uc.DB)

	now := time.Now().UTC()

	if err != nil {
		return err
	}

	transaction, err := uc.DB.Begin()
	if err != nil {
		return err
	}

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

	_, err = model.Add(TransactionBody, transaction)
	if err != nil {
		fmt.Println(1)
		transaction.Rollback()
		return err
	}

	transaction.Commit()

	return nil
}

func (uc TransactionUseCase) EditTransaction(input request.TransactionRequest) (err error) {
	model := actions.NewTransactionModel(uc.DB)

	now := time.Now().UTC()

	if err != nil {
		return err
	}

	transaction, err := uc.DB.Begin()
	if err != nil {
		return err
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
		UpdatedAt:       now.Format(time.RFC3339),
		CreatedAt:       now.Format(time.RFC3339),
	}

	_, err = model.Edit(TransactionBody, transaction)
	if err != nil {
		fmt.Println(1)
		transaction.Rollback()
		return err
	}

	transaction.Commit()

	return nil
}

func (uc TransactionUseCase) DeleteTransactions(ID string) (err error) {
	model := actions.NewTransactionModel(uc.DB)
	now := time.Now().UTC()

	isExist, err := uc.IsTransactionExist(ID)
	if err != nil {

		return err
	}
	if !isExist {

		return errors.New(messages.DataNotFound)
	}

	transaction, err := uc.DB.Begin()
	if err != nil {

		return err
	}

	err = model.Delete(ID, now.Format(time.RFC3339), now.Format(time.RFC3339), transaction)
	if err != nil {
		fmt.Println(6)
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
	var newAmount int
	createNew := false
	//check if fcustomer already exist in books debt
	debtExist, err := booksDebtUC.IsDebtCustomerExist(customerData.ID, enums.Nunggak)
	if err != nil {
		return err
	}

	bookdebt, err := booksDebtUC.BrowseByUser(customerData.ID, enums.Nunggak)
	if err != nil {
		return err
	}

	if debtExist {
		//edit booksDebt, status akan terus nunggak baik itu user yang hutang atau customer yang hutang.

		getTrans, err := model.Read(input.ID)
		if err != nil {
			fmt.Println(err)
			return err
		}

		if input.TransactionType == enums.Debet {
			fmt.Println("disini")
			if bookdebt.DebtTotal > 0 {
				if int(input.Amount) > int(getTrans.Amount.Int32) {
					debtAmount = int(bookdebt.DebtTotal) + (int(input.Amount) - int(getTrans.Amount.Int32))
					creditAmount = bookdebt.CreditTotal
					status = enums.Nunggak
				} else {
					fmt.Println(bookdebt.DebtTotal)
					fmt.Println(getTrans)
					debtAmount = int(bookdebt.DebtTotal) + (int(getTrans.Amount.Int32) - int(input.Amount))
					fmt.Println(debtAmount)
					creditAmount = bookdebt.CreditTotal
					if debtAmount == 0 {
						status = enums.Nunggak
					}
				}
			}

			if bookdebt.CreditTotal > 0 {
				if int(input.Amount) > int(getTrans.Amount.Int32) {
					newAmount = (int(input.Amount) - int(getTrans.Amount.Int32)) - int(bookdebt.CreditTotal)
					debtAmount = bookdebt.DebtTotal + newAmount
					creditAmount = 0
					status = enums.Nunggak
					createNew = true
				} else {
					creditAmount = int(bookdebt.CreditTotal) - (int(getTrans.Amount.Int32) - int(input.Amount))
					debtAmount = bookdebt.DebtTotal
					if creditAmount == 0 {
						status = enums.Lunas
					} else {
						status = enums.Nunggak
					}
				}
			}
		} else {
			if bookdebt.CreditTotal > 0 {
				if int(input.Amount) > int(getTrans.Amount.Int32) {
					creditAmount = int(bookdebt.CreditTotal) + (int(input.Amount) - int(getTrans.Amount.Int32))
					debtAmount = bookdebt.DebtTotal
					status = enums.Nunggak
				} else {
					creditAmount = int(bookdebt.CreditTotal) - (int(getTrans.Amount.Int32) - int(input.Amount))
					debtAmount = bookdebt.DebtTotal
					if creditAmount == 0 {
						status = enums.Lunas
					} else {
						status = enums.Nunggak
					}
				}
			}

			if bookdebt.DebtTotal > 0 {
				if int(input.Amount) > int(getTrans.Amount.Int32) {
					newAmount = (int(input.Amount) - int(getTrans.Amount.Int32)) - bookdebt.DebtTotal
					creditAmount = int(bookdebt.CreditTotal) + newAmount
					debtAmount = 0
					status = enums.Nunggak
					createNew = true
				} else {
					debtAmount = int(bookdebt.DebtTotal) - (int(getTrans.Amount.Int32) - int(input.Amount))
					creditAmount = 0
					if debtAmount == 0 {
						status = enums.Lunas
					} else {
						status = enums.Nunggak
					}
				}
			}
		}

		if createNew {
			fmt.Println("new")
			booksInput := request.BooksDebtRequest{
				CustomerID:     input.ReferenceID,
				SubmissionDate: input.TransactionDate,
				DebtTotal:      debtAmount,
				CreditTotal:    creditAmount,
				BillDate:       input.BillDate,
				Status:         enums.Nunggak,
				CreatedAt:      now.Format(time.RFC3339),
				UpdatedAt:      now.Format(time.RFC3339),
			}
			booksID, err := booksDebtUC.Add(booksInput, input.CustomerID, transaction)
			if err != nil {
				transaction.Rollback()
				return err
			}

			bookEditInput := request.BooksDebtRequest{
				CustomerID:     input.ReferenceID,
				SubmissionDate: now.Format(time.RFC3339),
				DebtTotal:      0,
				CreditTotal:    0,
				Status:         enums.Lunas,
				UpdatedAt:      now.Format(time.RFC3339),
			}
			err = booksDebtUC.Edit(bookEditInput, bookdebt.ID, transaction)
			if err != nil {
				transaction.Rollback()
				return err
			}

			TransactionBody := viewmodel.TransactionVm{
				ReferenceID:     input.ReferenceID,
				ShopID:          input.ShopID,
				Amount:          int32(newAmount),
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
				transaction.Rollback()
				return err
			}
		} else {
			booksInput := request.BooksDebtRequest{
				CustomerID:     customerData.ID,
				SubmissionDate: bookdebt.SubmissionDate,
				DebtTotal:      debtAmount,
				CreditTotal:    creditAmount,
				Status:         status,
				CreatedAt:      bookdebt.CreatedAt,
				UpdatedAt:      now.Format(time.RFC3339),
			}
			err = booksDebtUC.Edit(booksInput, bookdebt.ID, transaction)
			if err != nil {
				transaction.Rollback()
				return err
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
				BooksDebtID:     bookdebt.ID,
				UpdatedAt:       now.Format(time.RFC3339),
				CreatedAt:       getTrans.CreatedAt,
			}

			_, err = model.Edit(TransactionBody, transaction)
			if err != nil {
				transaction.Rollback()
				return err
			}
		}
	}

	transaction.Commit()

	return nil
}

//ini untuk add customer debt
func (uc TransactionUseCase) AddDebt(input request.TransactionRequest) (err error) {
	model := actions.NewTransactionModel(uc.DB)
	booksDebtUC := BooksDebtUseCase{UcContract: uc.UcContract}
	now := time.Now().UTC()
	var debtAmount int
	var creditAmount int
	var status string
	var booksID string
	var createNew bool

	transaction, err := uc.DB.Begin()
	if err != nil {
		return err
	}

	//check if customer already exist in books debt
	debtExist, err := booksDebtUC.IsDebtCustomerExist(input.ReferenceID, enums.Nunggak)
	if err != nil {
		return err
	}

	if debtExist {
		bookdebts, err := booksDebtUC.BrowseByUser(input.ReferenceID, enums.Nunggak)
		if err != nil {
			return err
		}
		if input.TransactionType == enums.Credit {
			if bookdebts.CreditTotal > 0 {
				createNew = false
				creditAmount = bookdebts.CreditTotal + int(input.Amount)
				status = enums.Nunggak
			}

			if bookdebts.DebtTotal > 0 {
				if int(input.Amount) > bookdebts.DebtTotal {
					createNew = true
					creditAmount = int(input.Amount) - bookdebts.DebtTotal
					debtAmount = 0
					status = enums.Lunas
				} else {
					createNew = false
					debtAmount = bookdebts.DebtTotal - int(input.Amount)
					if debtAmount == 0 {
						status = enums.Lunas
					} else {
						status = enums.Nunggak
					}
				}
			}
		} else {
			fmt.Println("debet")
			if bookdebts.DebtTotal > 0 {
				createNew = false
				debtAmount = bookdebts.DebtTotal + int(input.Amount)
				status = enums.Nunggak
			}

			if bookdebts.CreditTotal > 0 {
				if int(input.Amount) > bookdebts.CreditTotal {
					createNew = true
					debtAmount = int(input.Amount) - bookdebts.CreditTotal
					creditAmount = 0
					status = enums.Lunas
				} else {
					createNew = false
					creditAmount = bookdebts.CreditTotal - int(input.Amount)
					if creditAmount == 0 {
						status = enums.Lunas
					} else {
						status = enums.Nunggak
					}
				}
			}
		}

		booksID = bookdebts.ID
		if createNew {
			fmt.Println("ini satu")
			booksInput := request.BooksDebtRequest{
				CustomerID:     input.ReferenceID,
				SubmissionDate: input.TransactionDate,
				DebtTotal:      debtAmount,
				CreditTotal:    creditAmount,
				BillDate:       input.BillDate,
				Status:         enums.Nunggak,
				CreatedAt:      now.Format(time.RFC3339),
				UpdatedAt:      now.Format(time.RFC3339),
			}
			booksID, err = booksDebtUC.Add(booksInput, input.CustomerID, transaction)
			if err != nil {
				transaction.Rollback()
				return err
			}

			bookEditInput := request.BooksDebtRequest{
				CustomerID:     input.ReferenceID,
				SubmissionDate: now.Format(time.RFC3339),
				DebtTotal:      0,
				CreditTotal:    0,
				Status:         status,
				UpdatedAt:      now.Format(time.RFC3339),
			}
			err = booksDebtUC.Edit(bookEditInput, bookdebts.ID, transaction)
			if err != nil {
				transaction.Rollback()
				return err
			}
		} else {
			fmt.Println("ini dua")
			bookEditInput := request.BooksDebtRequest{
				CustomerID:     input.ReferenceID,
				SubmissionDate: now.Format(time.RFC3339),
				DebtTotal:      debtAmount,
				CreditTotal:    creditAmount,
				Status:         status,
				UpdatedAt:      now.Format(time.RFC3339),
			}
			err = booksDebtUC.Edit(bookEditInput, bookdebts.ID, transaction)
			if err != nil {
				transaction.Rollback()
				return err
			}
		}
	} else {
		fmt.Println("ini tiga")
		//for adding new debt so adding on books debt
		if input.TransactionType == enums.Debet {
			debtAmount = debtAmount + int(input.Amount)
		} else {
			creditAmount = creditAmount + int(input.Amount)
		}

		booksInput := request.BooksDebtRequest{
			CustomerID:     input.ReferenceID,
			SubmissionDate: now.Format(time.RFC3339),
			DebtTotal:      debtAmount,
			CreditTotal:    creditAmount,
			Status:         enums.Nunggak,
			CreatedAt:      now.Format(time.RFC3339),
			UpdatedAt:      now.Format(time.RFC3339),
		}
		booksID, err = booksDebtUC.Add(booksInput, input.CustomerID, transaction)
		if err != nil {
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

func (uc TransactionUseCase) CountBy(column, value string) (res int, err error) {
	model := actions.NewTransactionModel(uc.DB)
	res, err = model.CountBy(column, value)

	return res, err
}
