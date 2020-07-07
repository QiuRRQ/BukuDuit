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
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type ShopUseCase struct {
	*UcContract
}

func (uc ShopUseCase) ExportToFile(ID string) (res string, err error) {
	data, err := uc.Read(ID, "", "")
	if err != nil {
		return res, err
	}

	xlsx := excelize.NewFile()
	sheet1Name := "data utang/piutang"
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

	xlsx.SetCellValue(sheet1Name, "A1", "Nama")
	xlsx.SetCellValue(sheet1Name, "B1", "Nominal")
	xlsx.SetCellValue(sheet1Name, "C1", "Tipe")

	err = xlsx.AutoFilter(sheet1Name, "A1", "B1", "")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	var tipe string
	for i, each := range data.UserCustomers {
		if data.UserCustomers[i].Type == enums.Credit {
			tipe = "utang"
		} else {
			tipe = "piutang"
		}
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+2), each.FullName)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+2), each.Amount)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+2), tipe)
	}

	err = xlsx.SaveAs("./../file/data.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	return "./../file/data.xlsx", err
}

func (uc ShopUseCase) BrowseByUser(userID string) (res []viewmodel.ShopVm, err error) {
	model := actions.NewShopModel(uc.DB)
	shops, err := model.BrowseByUser(userID)
	if err != nil {
		return res, err
	}

	for _, shop := range shops {
		res = append(res, viewmodel.ShopVm{
			ID:          shop.ID,
			FullName:    shop.FullName.String,
			BookName:    shop.BookName,
			MobilePhone: shop.MobilePhone.String,
			TagLine:     shop.TagLine.String,
			Address:     shop.Address.String,
			Email:       shop.Email.String,
			Avatar:      shop.Avatar.String,
			CreatedAt:   shop.CreatedAt,
			UpdatedAt:   shop.UpdatedAt.String,
			DeletedAt:   shop.DeletedAt.String,
		})
	}

	return res, err
}

//fucntion for hutang list lunas
func (uc ShopUseCase) Read(ID, lunas, name string) (res viewmodel.ShopVm, err error) { //lunas = 1
	model := actions.NewShopModel(uc.DB)
	userCustomerUC := UserCustomerUseCase{UcContract: uc.UcContract}
	bookDebtUc := BooksDebtUseCase{UcContract: uc.UcContract}
	var debtTotal int
	var creditTotal int

	tempDataUserCustomer, err := userCustomerUC.BrowseByShop(ID, "")
	if err != nil {
		return res, err
	}

	var elems int
	var amount int32
	var typeBook string
	var dataUserCustomer = make([]viewmodel.UserCustomerDebetCreditVm, elems)
	for _, data := range tempDataUserCustomer {
		if lunas == "1" {
			bookdebtsLunas, err := bookDebtUc.BrowseByUser(data.ID, enums.Lunas)
			if err == nil {
				if bookdebtsLunas.CreditTotal == 0 {
					dataUserCustomer = append(dataUserCustomer, viewmodel.UserCustomerDebetCreditVm{
						ID:          data.ID,
						FullName:    data.FullName,
						Amount:      int32(bookdebtsLunas.CreditTotal),
						Type:        typeBook,
						MobilePhone: data.MobilePhone,
					})
				}
			}
		} else {
			bookdebtsNunggak, err := bookDebtUc.BrowseByUser(data.ID, enums.Nunggak)
			fmt.Println(bookdebtsNunggak)
			if err == nil {
				creditTotal = creditTotal + bookdebtsNunggak.CreditTotal
				debtTotal = debtTotal + bookdebtsNunggak.DebtTotal
				if bookdebtsNunggak.CreditTotal != 0 {
					typeBook = enums.Credit
					amount = int32(bookdebtsNunggak.CreditTotal)
				} else {
					typeBook = enums.Debet
					amount = int32(bookdebtsNunggak.DebtTotal)
				}
				dataUserCustomer = append(dataUserCustomer, viewmodel.UserCustomerDebetCreditVm{
					ID:          data.ID,
					FullName:    data.FullName,
					Amount:      amount,
					Type:        typeBook,
					MobilePhone: data.MobilePhone,
				})
			}
		}
	}

	if name != "" {
		for _, v := range dataUserCustomer {
			if strings.Contains(strings.ToLower(v.FullName), strings.ToLower(name)) {
				dataUserCustomer = nil
				dataUserCustomer = append(dataUserCustomer, viewmodel.UserCustomerDebetCreditVm{
					ID:          v.ID,
					FullName:    v.FullName,
					Amount:      v.Amount,
					Type:        v.Type,
					MobilePhone: v.MobilePhone,
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

	shop, err := model.Read(ID)
	if err != nil {
		return res, err
	}

	res = viewmodel.ShopVm{
		ID:                  shop.ID,
		FullName:            shop.FullName.String,
		BookName:            shop.BookName,
		MobilePhone:         shop.MobilePhone.String,
		TagLine:             shop.TagLine.String,
		Address:             shop.Address.String,
		Email:               shop.Email.String,
		UserCustomers:       dataUserCustomer,
		TotalCustomerCredit: int32(creditTotal),
		TotalOwnerCredit:    int32(debtTotal),
		CreatedAt:           shop.CreatedAt,
		UpdatedAt:           shop.UpdatedAt.String,
		DeletedAt:           shop.UpdatedAt.String,
	}

	return res, err
}

func (uc ShopUseCase) Edit(input *request.ShopRequest, ID string) (err error) {
	model := actions.NewShopModel(uc.DB)
	now := time.Now().UTC()
	isExist, err := uc.IsShopExist(ID)
	if err != nil {
		return err
	}
	if !isExist {
		return errors.New(messages.DataNotFound)
	}

	body := viewmodel.ShopVm{
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

func (uc ShopUseCase) Add(input *request.ShopRequest, userID string) (err error) {
	model := actions.NewShopModel(uc.DB)
	now := time.Now().UTC()

	body := viewmodel.ShopVm{
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

func (uc ShopUseCase) Register(userID, bookName, createdAt, updatedAt string, tx *sql.Tx) (err error) {
	model := actions.NewShopModel(uc.DB)
	body := viewmodel.ShopVm{
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

func (uc ShopUseCase) Delete(ID string) (err error) {
	fmt.Println(ID)
	model := actions.NewShopModel(uc.DB)
	now := time.Now().UTC()

	isExist, err := uc.IsShopExist(ID)
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

func (uc ShopUseCase) DeleteByUser(userID string, tx *sql.Tx) (err error) {
	model := actions.NewShopModel(uc.DB)
	now := time.Now().UTC()

	err = model.DeleteByUser(userID, now.Format(time.RFC3339), now.Format(time.RFC3339), tx)
	if err != nil {
		return err
	}

	return nil
}

func (uc ShopUseCase) IsShopExist(ID string) (res bool, err error) {
	model := actions.NewShopModel(uc.DB)
	count, err := model.CountByPk(ID)
	if err != nil {
		return res, err
	}

	return count > 0, err
}
