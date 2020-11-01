package usecase

import (
	"bukuduit-go/db/repositories/actions"
	"bukuduit-go/helpers/messages"
	request "bukuduit-go/server/requests"
	"bukuduit-go/usecase/viewmodel"
	"errors"
	"fmt"
	"time"
)

type BarcodeUseCase struct {
	*UcContract
}

func (uc BarcodeUseCase) Read() (res []viewmodel.BarcodeVM, err error) {
	model := actions.NewBarcodeModel(uc.DB)
	barcodes, err := model.Read()
	if err != nil {
		return res, err
	}

	for _, barcode := range barcodes {
		res = append(res, viewmodel.BarcodeVM{
			Barcode:    barcode.Barcode,
			BooksID:    barcode.BooksId,
			Created_at: barcode.Created_at,
			UPdated_at: barcode.Updated_at,
		})
	}

	return res, err
}

func (uc BarcodeUseCase) ReadById(ID string) (res []viewmodel.BarcodeVM, err error) {
	model := actions.NewBarcodeModel(uc.DB)
	barcodes, err := model.ReadByID(ID)
	if err != nil {
		return res, err
	}

	for _, barcode := range barcodes {
		res = append(res, viewmodel.BarcodeVM{
			Barcode:    barcode.Barcode,
			BooksID:    barcode.BooksId,
			Created_at: barcode.Created_at,
			UPdated_at: barcode.Updated_at,
		})
	}

	return res, err
}

func (uc BarcodeUseCase) Add(input *request.BarcodeRequest) (err error) {
	model := actions.NewBarcodeModel(uc.DB)
	now := time.Now().UTC()

	body := viewmodel.BarcodeVM{
		BooksID:    input.BooksID,
		UPdated_at: now.Format(time.RFC3339),
		Created_at: now.Format(time.RFC3339),
	}
	_, err = model.Add(body, nil)
	if err != nil {
		return err
	}

	return nil
}

func (uc BarcodeUseCase) Delete(ID string) (err error) {
	fmt.Println(ID)
	model := actions.NewBarcodeModel(uc.DB)
	now := time.Now().UTC()

	isExist, err := uc.IsBarcodeExist(ID)
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

func (uc BarcodeUseCase) IsBarcodeExist(ID string) (res bool, err error) {
	model := actions.NewBarcodeModel(uc.DB)
	count, err := model.CountByPk(ID)
	if err != nil {
		return res, err
	}

	return count > 0, err
}
