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
	"log"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
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
	if err != nil { //data not found
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

	if Transactions != nil {
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
						Name:        Transactions[i].Name.String,
						Description: Transactions[i].Description.String,
						Amount:      Transactions[i].Amount.Int32,
						Type:        Transactions[i].Type,
					})

				} else {
					debtDetails = append(debtDetails, viewmodel.DataDetails{
						ID:          Transactions[i].ID,
						ReferenceID: Transactions[i].ReferenceID,
						Name:        Transactions[i].Name.String,
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
					Name:        Transactions[i].Name.String,
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
	}

	if resultCount > 0 && Transactions != nil {
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
	}

	return res, err
}

//list transaksi by week
func (uc TransactionUseCase) TransactionListByWeeks(shopID, search, name, amount, transDate, timeGroup, startDate, endDate string) (res viewmodel.TransactionListVm, err error) {
	model := actions.NewTransactionModel(uc.DB)

	var filter string
	var debtDate []viewmodel.DataList
	var debtDetails []viewmodel.DataDetails
	var dateCreditAmount int
	var dateDebetAmount int
	var debtTotal int
	var creditTotal int
	weekIndex := 0
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
	if timeGroup == enums.Week {
		filter = filter + `order by weekly desc`
	}

	Transactions, err := model.TransactionBrowsByShop(shopID, filter)
	if err != nil {
		return res, err
	}
	weekly, err := model.GroubByWeeksMonth(enums.Week)
	if err != nil {
		return res, err
	}

	if Transactions != nil {
		for i := 0; i < len(Transactions); i++ {
			weeks, err := strconv.Atoi(Transactions[i].Weekly.String)
			if err != nil {
				return res, err
			}
			if Transactions[i].Type == enums.Debet {
				debtTotal = debtTotal + int(Transactions[i].Amount.Int32)
			} else {
				creditTotal = creditTotal + int(Transactions[i].Amount.Int32)
			}
			if weekly[weekIndex].Weekly.Int32 == int32(weeks) {
				if Transactions[i].Type == enums.Debet {
					dateDebetAmount = dateDebetAmount + int(Transactions[i].Amount.Int32)
				} else {
					dateCreditAmount = dateCreditAmount + int(Transactions[i].Amount.Int32)
				}

				debtDetails = append(debtDetails, viewmodel.DataDetails{
					ID:          Transactions[i].ID,
					ReferenceID: Transactions[i].ReferenceID,
					Name:        Transactions[i].Name.String,
					Description: Transactions[i].Description.String,
					Amount:      Transactions[i].Amount.Int32,
					Type:        Transactions[i].Type,
				})
			} else {
				weekIndex++
				debtDate = append(debtDate, viewmodel.DataList{
					TransactionDate:  "minggu - sabtu",
					DateAmountCredit: dateCreditAmount,
					DateAmountDebet:  dateDebetAmount,
					Details:          debtDetails,
				})
				debtDetails = nil
				dateDebetAmount = 0
				dateCreditAmount = 0
				if weekly[weekIndex].Weekly.Int32 == int32(weeks) {
					if Transactions[i].Type == enums.Debet {
						dateDebetAmount = dateDebetAmount + int(Transactions[i].Amount.Int32)
					} else {
						dateCreditAmount = dateCreditAmount + int(Transactions[i].Amount.Int32)
					}

					debtDetails = append(debtDetails, viewmodel.DataDetails{
						ID:          Transactions[i].ID,
						ReferenceID: Transactions[i].ReferenceID,
						Name:        Transactions[i].Name.String,
						Description: Transactions[i].Description.String,
						Amount:      Transactions[i].Amount.Int32,
						Type:        Transactions[i].Type,
					})
				}

			}
		}

		debtDate = append(debtDate, viewmodel.DataList{
			TransactionDate:  "minggu - sabtu",
			DateAmountCredit: dateCreditAmount,
			DateAmountDebet:  dateDebetAmount,
			Details:          debtDetails,
		})
		res = viewmodel.TransactionListVm{
			ShopID:      Transactions[0].IDShop,
			TotalCredit: creditTotal,
			TotalDebit:  debtTotal,
			ListData:    debtDate,
			CreatedAt:   Transactions[0].CreatedAt,
			UpdatedAt:   Transactions[0].UpdatedAt.String,
			DeletedAt:   Transactions[0].DeletedAt.String,
		}

	}
	return res, err
}

//list transaksi by months
func (uc TransactionUseCase) TransactionListMonth(shopID, searching, name, amount, transDate, startDate, endDate string) (res viewmodel.TransactionListVm, err error) {

	model := actions.NewTransactionModel(uc.DB)
	transDate = "asc"

	data, err := uc.TransactionList(shopID, searching, name, amount, transDate, "", startDate, endDate)
	if err != nil {
		return res, err
	}

	monthly, err := model.GroubByWeeksMonth(enums.Month)
	fmt.Println(monthly)
	if err != nil {
		return res, err
	}

	var monthlyIndex = 0
	var debtDate []viewmodel.DataList
	var debtDetails []viewmodel.DataDetails
	var dateCreditAmount int
	var dateDebetAmount int
	var prefMonth time.Month
	if data.ListData != nil {
		for i := 0; i < len(data.ListData); i++ {
			t, err := time.Parse("2006-01-02 00:00:00 +0000 UTC", data.ListData[i].TransactionDate)
			if err != nil {
				fmt.Println(err)
			}

			_, currentMonth, _ := t.Date()

			if monthly[monthlyIndex].Monthly.Int32 == int32(currentMonth) {
				for j := 0; j < len(data.ListData[i].Details); j++ {
					fmt.Println(data.ListData[i].TransactionDate)
					prefMonth = currentMonth
					if data.ListData[i].Details[j].Type == enums.Debet {
						dateDebetAmount = dateDebetAmount + int(data.ListData[i].Details[j].Amount)
					} else {
						dateCreditAmount = dateCreditAmount + int(data.ListData[i].Details[j].Amount)
					}
					debtDetails = append(debtDetails, data.ListData[i].Details[j])
				}
			} else {
				monthlyIndex++
				fmt.Println(currentMonth)
				debtDate = append(debtDate, viewmodel.DataList{
					TransactionDate:  prefMonth.String(),
					DateAmountCredit: dateCreditAmount,
					DateAmountDebet:  dateDebetAmount,
					Details:          debtDetails,
				})
				debtDetails = nil
				dateDebetAmount = 0
				dateCreditAmount = 0
				if monthly[monthlyIndex].Monthly.Int32 == int32(currentMonth) {
					for j := 0; j < len(data.ListData[i].Details); j++ {
						fmt.Println(data.ListData[i].TransactionDate)
						if data.ListData[i].Details[j].Type == enums.Debet {
							dateDebetAmount = dateDebetAmount + int(data.ListData[i].Details[j].Amount)
						} else {
							dateCreditAmount = dateCreditAmount + int(data.ListData[i].Details[j].Amount)
						}
						debtDetails = append(debtDetails, data.ListData[i].Details[j])
					}
				}

				if i == len(data.ListData)-1 {
					prefMonth = currentMonth
				}
			}
		}
		debtDate = append(debtDate, viewmodel.DataList{
			TransactionDate:  prefMonth.String(),
			DateAmountCredit: dateCreditAmount,
			DateAmountDebet:  dateDebetAmount,
			Details:          debtDetails,
		})
		res = viewmodel.TransactionListVm{
			ShopID:      data.ShopID,
			TotalCredit: data.TotalCredit,
			TotalDebit:  data.TotalDebit,
			ListData:    debtDate,
			CreatedAt:   data.CreatedAt,
			UpdatedAt:   data.UpdatedAt,
			DeletedAt:   data.DeletedAt,
		}

	}
	return res, err

}

//list transaksi by days
func (uc TransactionUseCase) TransactionList(shopID, search, name, amount, transDate, timeGroup, startDate, endDate string) (res viewmodel.TransactionListVm, err error) {
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

	Transactions, err := model.TransactionBrowsByShop(shopID, filter)
	if err != nil {
		return res, err
	}
	resultCount, err := model.CountDistinctBy("shop_id", shopID)
	if err != nil {
		return res, err
	}

	var debtTotal int
	var creditTotal int

	var dateCreditAmount int
	var dateDebetAmount int

	var debtDate []viewmodel.DataList
	var debtDetails []viewmodel.DataDetails

	if Transactions != nil {
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
						Name:        Transactions[i].Name.String,
						Description: Transactions[i].Description.String,
						Amount:      Transactions[i].Amount.Int32,
						Type:        Transactions[i].Type,
					})

				} else {
					debtDetails = append(debtDetails, viewmodel.DataDetails{
						ID:          Transactions[i].ID,
						ReferenceID: Transactions[i].ReferenceID,
						Name:        Transactions[i].Name.String,
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
					Name:        Transactions[i].Name.String,
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
	}

	if resultCount > 0 && Transactions != nil {

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
	}

	return res, err
}

func (uc TransactionUseCase) BrowseByBookDebtID(bookDebtID string, status int) (res []viewmodel.TransactionVm, err error) {
	model := actions.NewTransactionModel(uc.DB)
	transactions, err := model.BrowseByBookDebtID(bookDebtID, status)

	if err != nil {
		return res, err
	}

	for _, transaction := range transactions {
		res = append(res, viewmodel.TransactionVm{
			ID:              transaction.ID,
			ReferenceID:     transaction.ReferenceID,
			Name:            transaction.Name.String,
			ShopID:          transaction.IDShop,
			CustomerID:      transaction.CustomerID.String,
			Amount:          transaction.Amount.Int32,
			Description:     transaction.Description.String,
			Image:           transaction.Image.String,
			Type:            transaction.Type,
			BooksDebtID:     transaction.BooksDeptID.String,
			BooksTransID:    transaction.BooksTransID.String,
			TransactionDate: transaction.TransactionDate.String,
			CreatedAt:       transaction.CreatedAt,
			UpdatedAt:       transaction.UpdatedAt.String,
			DeletedAt:       transaction.DeletedAt.String,
		})
	}

	return res, err
}

//export file untuk laporan transaksi
func (uc TransactionUseCase) TransactionReportExportFile(shopID, searching, name, amount, date, startDate, endDate string) (res string, err error) {
	data, err := uc.TransactionReport(shopID, searching, name, amount, date, startDate, endDate)
	if err != nil {
		return res, err
	}

	xlsx := excelize.NewFile()
	sheet1Name := "data Laporan Transaksi"
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

	xlsx.SetCellValue(sheet1Name, "A1", "Range Tanggal :")
	xlsx.SetCellValue(sheet1Name, "B1", startDate+" - "+endDate)

	xlsx.SetCellValue(sheet1Name, "A4", "No")
	xlsx.SetCellValue(sheet1Name, "B4", "Tanggal Transaksi")
	xlsx.SetCellValue(sheet1Name, "C4", "Kategori")
	xlsx.SetCellValue(sheet1Name, "D4", "Nama Pelanggan")
	xlsx.SetCellValue(sheet1Name, "E4", "Deskripsi")
	xlsx.SetCellValue(sheet1Name, "F4", "Uang Keluar")
	xlsx.SetCellValue(sheet1Name, "G4", "Uang Masuk")

	err = xlsx.AutoFilter(sheet1Name, "A1", "B1", "")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	var debtTotal int
	var creditTotal int
	no := 1
	var displayData []viewmodel.DataDetails
	for _, each := range data.ListData {
		for _, row := range each.Details {
			displayData = append(displayData, viewmodel.DataDetails{
				ID:              row.ID,
				TransactionDate: each.TransactionDate,
				ReferenceID:     row.ReferenceID,
				Name:            row.Name,
				Description:     row.Description,
				Amount:          row.Amount,
				Type:            row.Type,
			})
		}
	}

	for i, each := range displayData {
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+5), no)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+5), each.TransactionDate)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", i+5), each.Name)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", i+5), each.Description)
		if each.Type == enums.Credit {
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("F%d", i+5), each.Amount)
			creditTotal = creditTotal + int(each.Amount)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("G%d", i+5), "-")
		} else {
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("F%d", i+5), "-")
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("G%d", i+5), each.Amount)
			debtTotal = debtTotal + int(each.Amount)
		}

		if i == len(data.ListData)-1 {
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", i+7), "Total")
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("F%d", i+7), creditTotal)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("G%d", i+7), debtTotal)
		}
		no++
	}
	if debtTotal > creditTotal {
		xlsx.SetCellValue(sheet1Name, "A2", "Jumlah Transaksi:")
		xlsx.SetCellValue(sheet1Name, "B2", debtTotal-creditTotal)
	} else {
		xlsx.SetCellValue(sheet1Name, "A2", "Jumlah Transaksi:")
		xlsx.SetCellValue(sheet1Name, "B2", creditTotal-debtTotal)
	}

	err = xlsx.SaveAs("./../file/LaporanTransaksi.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	return "./../file/LaporanTransaksi.xlsx", err
}

//export file untuk list piutang per pelanggan
func (uc TransactionUseCase) DebtDetailExportFile(customerID string) (res string, err error) {
	data, err := uc.BrowseByCustomer(customerID)
	if err != nil {
		return res, err
	}
	xlsx := excelize.NewFile()
	sheet1Name := "data Laporan Hutang Per Pelanggan"
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

	xlsx.SetCellValue(sheet1Name, "A1", "Nama Pelanggan:")
	xlsx.SetCellValue(sheet1Name, "B1", data.Name)

	xlsx.SetCellValue(sheet1Name, "A4", "No")
	xlsx.SetCellValue(sheet1Name, "B4", "Tanggal Transaksi")
	xlsx.SetCellValue(sheet1Name, "C4", "Catatan")
	xlsx.SetCellValue(sheet1Name, "D4", "Uang Masuk")
	xlsx.SetCellValue(sheet1Name, "E4", "Uang Keluar")

	err = xlsx.AutoFilter(sheet1Name, "A1", "B1", "")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	var debtTotal int
	var creditTotal int
	no := 1
	var displayData []viewmodel.DataDetails
	for _, each := range data.ListData {
		for _, row := range each.Details {
			displayData = append(displayData, viewmodel.DataDetails{
				ID:              row.ID,
				TransactionDate: each.TransactionDate,
				Description:     row.Description,
				Amount:          row.Amount,
				Type:            row.Type,
			})
		}
	}

	for i, each := range displayData {
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+5), no)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+5), each.TransactionDate)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+5), each.Description)
		if each.Type == enums.Credit {
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", i+5), each.Amount)
			creditTotal = creditTotal + int(each.Amount)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", i+5), "-")
		} else {
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", i+5), "-")
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", i+5), each.Amount)
			debtTotal = debtTotal + int(each.Amount)
		}

		if i == len(data.ListData)-1 {
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+7), "Total")
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", i+7), creditTotal)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", i+7), debtTotal)
		}
		no++
	}
	if debtTotal > creditTotal {
		xlsx.SetCellValue(sheet1Name, "A2", "Jumlah Utang Piutang:")
		xlsx.SetCellValue(sheet1Name, "B2", debtTotal-creditTotal)
	} else {
		xlsx.SetCellValue(sheet1Name, "A2", "Jumlah Utang Piutang:")
		xlsx.SetCellValue(sheet1Name, "B2", creditTotal-debtTotal)
	}

	err = xlsx.SaveAs("./../file/HutangPerPelanggan.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	return "./../file/HutangPerPelanggan.xlsx", err
}

//export file untuk laporan hutang
func (uc TransactionUseCase) DebtReportExportFile(shopID, searching, name, amount, date, startDate, endDate string) (res string, err error) {
	data, err := uc.DebtReport(shopID, searching, name, amount, date, startDate, endDate)
	if err != nil {
		return res, err
	}

	xlsx := excelize.NewFile()
	sheet1Name := "data Laporan Hutang"
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

	xlsx.SetCellValue(sheet1Name, "A1", "Tanggal Penagihan/Pembayaran")
	xlsx.SetCellValue(sheet1Name, "B1", "Nama Pelanggan")
	xlsx.SetCellValue(sheet1Name, "C1", "Deskripsi")
	xlsx.SetCellValue(sheet1Name, "D1", "Uang Masuk")
	xlsx.SetCellValue(sheet1Name, "E1", "Uang Keluar")

	err = xlsx.AutoFilter(sheet1Name, "A1", "B1", "")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	var debtTotal int
	var creditTotal int
	var displayData []viewmodel.DataDetails
	for _, each := range data.ListData {
		for _, row := range each.Details {
			displayData = append(displayData, viewmodel.DataDetails{
				TransactionDate: each.TransactionDate,
				Name:            row.Name,
				Description:     row.Description,
				Amount:          row.Amount,
				Type:            row.Type,
			})

		}
	}

	for i, each := range displayData {
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+2), each.TransactionDate)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+2), each.Name)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+2), each.Description)
		if each.Type == enums.Credit {
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", i+2), each.Amount)
			creditTotal = creditTotal + int(each.Amount)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", i+2), "-")
		} else {
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", i+2), "-")
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", i+2), each.Amount)
			debtTotal = debtTotal + int(each.Amount)
		}

		if i == len(displayData)-1 {

			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+4), "Total")
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", i+4), creditTotal)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", i+4), debtTotal)
		}
	}

	err = xlsx.SaveAs("./../file/laporanHutang.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	return "./../file/laporanHutang.xlsx", err
}

//untuk laporan hutang
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
						Name:        transactions[i].Name.String,
						Description: transactions[i].Description.String,
						Amount:      transactions[i].Amount.Int32,
						Type:        transactions[i].Type,
					})

				} else {
					debtDetails = append(debtDetails, viewmodel.DebtDetail{
						ID:          transactions[i].ID,
						ReferenceID: transactions[i].ReferenceID,
						Name:        transactions[i].Name.String,
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
					Name:        transactions[i].Name.String,
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
			Name:        Transactions[0].Name.String,
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
	var debitAmount int32
	var creditAmount int32
	var status string

	customerData, err := userCustomerUc.Read(input.ReferenceID)
	if err != nil {
		fmt.Println(1)
		fmt.Println(err)
		return err
	}

	transaction, err := uc.DB.Begin()
	if err != nil {
		return err
	}

	//check if fcustomer already exist in books debt
	debtExist, err := booksDebtUC.IsDebtCustomerExist(customerData.ID, enums.Nunggak)
	if err != nil {
		fmt.Println(2)
		fmt.Println(err)
		return err
	}

	if debtExist {
		bookdebt, err := booksDebtUC.BrowseByUser(customerData.ID, enums.Nunggak)
		if err != nil {
			fmt.Println(3)
			fmt.Println(err)
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
		}

		_, err = model.Edit(TransactionBody, transaction)
		if err != nil {
			fmt.Println(4)
			fmt.Println(err)
			transaction.Rollback()
			return err
		}

		transaction.Commit()

		transaction2, err := uc.DB.Begin()

		transactions, err := uc.BrowseByBookDebtID(bookdebt.ID, 0)
		if err != nil {
			fmt.Println(5)
			fmt.Println(err)
			return err
		}

		for _, transaction := range transactions {
			if transaction.Type == enums.Debet {
				debitAmount = debitAmount + transaction.Amount
			} else {
				creditAmount = creditAmount + transaction.Amount
			}
		}

		fmt.Println(debitAmount)
		fmt.Println(creditAmount)
		if debitAmount == creditAmount {
			debitAmount = 0
			creditAmount = 0
			status = enums.Lunas
		}

		if debitAmount > creditAmount {
			debitAmount = debitAmount - creditAmount
			status = enums.Nunggak
			creditAmount = 0
		}

		if creditAmount > debitAmount {
			creditAmount = creditAmount - debitAmount
			debitAmount = 0
			status = enums.Nunggak
		}

		bookEditInput := request.BooksDebtRequest{
			CustomerID:     input.ReferenceID,
			SubmissionDate: now.Format(time.RFC3339),
			DebtTotal:      int(debitAmount),
			CreditTotal:    int(creditAmount),
			Status:         status,
			UpdatedAt:      now.Format(time.RFC3339),
		}
		err = booksDebtUC.Edit(bookEditInput, bookdebt.ID, transaction2)
		if err != nil {
			fmt.Println(6)
			fmt.Println(err)
			transaction.Rollback()
			return err
		}
		transaction2.Commit()
	}

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
