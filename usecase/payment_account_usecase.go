package usecase

import (
	"bukuduit-go/db/repositories/actions"
	"bukuduit-go/helpers/messages"
	request "bukuduit-go/server/requests"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
	"errors"
	"time"
)

type PaymentAccountUseCase struct {
	*UcContract
}

func (uc PaymentAccountUseCase) BrowseByShop(shopID string) (res []viewmodel.PaymentAccountVm, err error) {
	model := actions.NewPaymentAccountModel(uc.DB)
	paymentAccounts, err := model.BrowseByShop(shopID)
	if err != nil {
		return res, err
	}

	for _, paymentAccount := range paymentAccounts {
		res = append(res, viewmodel.PaymentAccountVm{
			ID:            paymentAccount.ID,
			AccountName:   paymentAccount.Name,
			OwnerName:     paymentAccount.OwnerName,
			ShopID:        paymentAccount.ShopID,
			PaymentNumber: paymentAccount.PaymentNumber,
			CreatedAt:     paymentAccount.CreatedAt,
			UpdatedAt:     paymentAccount.UpdatedAt.String,
			DeletedAt:     paymentAccount.DeletedAt.String,
		})
	}

	return res, err
}
func (uc PaymentAccountUseCase) Read(ID, lunas string) (res viewmodel.PaymentAccountVm, err error) {
	model := actions.NewPaymentAccountModel(uc.DB)
	payment, err := model.Read(ID)
	if err != nil {
		return res, err
	}

	res = viewmodel.PaymentAccountVm{
		ID:            payment.ID,
		ShopID:        payment.ShopID,
		AccountName:   payment.Name,
		OwnerName:     payment.OwnerName,
		PaymentNumber: payment.PaymentNumber,
		CreatedAt:     payment.CreatedAt,
		UpdatedAt:     payment.UpdatedAt.String,
		DeletedAt:     payment.DeletedAt.String,
	}

	return res, err
}

func (uc PaymentAccountUseCase) Edit(input *request.PaymentAccountRequest, ID string) (err error) {
	model := actions.NewPaymentAccountModel(uc.DB)
	now := time.Now().UTC()
	isExist, err := uc.IsAccountPaymentExist(ID)
	if err != nil {
		return err
	}
	if !isExist {
		return errors.New(messages.DataNotFound)
	}

	body := viewmodel.PaymentAccountVm{
		ID:            ID,
		OwnerName:     input.OwnerName,
		AccountName:   input.AccountName,
		ShopID:        input.ShopID,
		PaymentNumber: input.PaymentNumber,
		UpdatedAt:     now.Format(time.RFC3339),
	}
	_, err = model.Edit(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc PaymentAccountUseCase) Add(input *request.PaymentAccountRequest) (err error) {
	model := actions.NewPaymentAccountModel(uc.DB)
	now := time.Now().UTC()

	Body := viewmodel.PaymentAccountVm{
		ShopID:        input.ShopID,
		OwnerName:     input.OwnerName,
		AccountName:   input.AccountName,
		PaymentNumber: input.PaymentNumber,
		UpdatedAt:     now.Format(time.RFC3339),
		CreatedAt:     now.Format(time.RFC3339),
	}
	_, err = model.Add(Body)

	if err != nil {
		return err
	}
	return nil
}

func (uc PaymentAccountUseCase) Delete(ID string) (err error) {
	model := actions.NewPaymentAccountModel(uc.DB)
	now := time.Now().UTC()

	isExist, err := uc.IsAccountPaymentExist(ID)
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

func (uc PaymentAccountUseCase) DeleteByUser(ShopID string, tx *sql.Tx) (err error) {
	model := actions.NewPaymentAccountModel(uc.DB)
	now := time.Now().UTC()

	err = model.DeleteByShop(ShopID, now.Format(time.RFC3339), now.Format(time.RFC3339), tx)
	if err != nil {
		return err
	}

	return nil
}

func (uc PaymentAccountUseCase) IsAccountPaymentExist(ID string) (res bool, err error) {
	model := actions.NewPaymentAccountModel(uc.DB)
	count, err := model.CountByPk(ID)
	if err != nil {
		return res, err
	}

	return count > 0, err
}
