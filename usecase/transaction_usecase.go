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

func (uc TransactionUseCase) DebtReport(shopID, search, name, amount, transDate, startDate, endDate string) (res viewmodel.ReportHutangVm, err error) {
	model := actions.NewTransactionModel(uc.DB)
	booksDebtUC := BooksDebtUseCase{UcContract: uc.UcContract}

	books, err := booksDebtUC.Browse(enums.Nunggak)

	if err != nil {
		return res, err
	}

	var debtTotal int
	var creditTotal int
	var customerID string
	var debtDate []viewmodel.DebtReport
	var debtDetails []viewmodel.DebtDetail

	//variabel total di sini gk berubah selama filter gk diset
	for i, book := range books {
		debtTotal = debtTotal + book.DebtTotal
		creditTotal = creditTotal + book.CreditTotal
		if i != len(books)-1 {
			customerID = customerID + "'" + book.CustomerID + "'" + ","
		} else {
			customerID = customerID + "'" + book.CustomerID + "'"
		}

	}

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

	transactions, err := model.DebtReport(customerID, shopID, filter)
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

	res = viewmodel.ReportHutangVm{
		TotalCredit: creditTotal,
		TotalDebit:  debtTotal,
		ListData:    debtDate,
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
	var getTrans models.Transactions
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
			if int(input.Amount) > int(getTrans.Amount.Int32) {
				debtAmount = int(bookdebt.DebtTotal) + (int(input.Amount) - int(getTrans.Amount.Int32))
				creditAmount = bookdebt.CreditTotal

			} else {
				debtAmount = int(bookdebt.DebtTotal) - (int(getTrans.Amount.Int32) - int(input.Amount))
				creditAmount = bookdebt.CreditTotal
			}

			fmt.Println(debtAmount)
			fmt.Println(creditAmount)
		} else {
			if int(input.Amount) > int(getTrans.Amount.Int32) {
				creditAmount = int(bookdebt.CreditTotal) + (int(input.Amount) - int(getTrans.Amount.Int32))
				debtAmount = bookdebt.DebtTotal

			} else {
				creditAmount = int(bookdebt.CreditTotal) - (int(getTrans.Amount.Int32) - int(input.Amount))
				debtAmount = bookdebt.DebtTotal
			}

			fmt.Println(debtAmount)
			fmt.Println(creditAmount)
		}

		booksInput := request.BooksDebtRequest{
			CustomerID:     customerData.ID,
			SubmissionDate: bookdebt.SubmissionDate,
			DebtTotal:      debtAmount,
			CreditTotal:    creditAmount,
			Status:         bookdebt.Status,
			CreatedAt:      bookdebt.CreatedAt,
			UpdatedAt:      now.Format(time.RFC3339),
		}
		err = booksDebtUC.Edit(booksInput, bookdebt.ID, transaction)
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
		BooksDebtID:     bookdebt.ID,
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
func (uc TransactionUseCase) AddDebt(input request.TransactionRequest) (err error) {
	model := actions.NewTransactionModel(uc.DB)
	booksDebtUC := BooksDebtUseCase{UcContract: uc.UcContract}
	userCustomerUc := UserCustomerUseCase{UcContract: uc.UcContract}
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
		fmt.Println("nunggak")
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

		if createNew {
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
		//for adding new debt so adding on books debt
		if input.TransactionType == enums.Debet {
			debtAmount = debtAmount - int(input.Amount)
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
