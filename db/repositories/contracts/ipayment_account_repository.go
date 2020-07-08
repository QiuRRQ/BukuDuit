package contracts

import (
	"bukuduit-go/db/models"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
)

type IPaymentAccountRepository interface {
	BrowseByShop(userID string) (data []models.PaymentAccount, err error)

	Read(ID string) (data models.PaymentAccount, err error)

	Edit(body viewmodel.PaymentAccountVm) (res string, err error)

	Add(body viewmodel.PaymentAccountVm) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	DeleteByShop(userID, updatedAt, deletedAt string, tx *sql.Tx) (err error)

	CountByPk(ID string) (res int, err error)

	CountBy(column, value string) (res int, err error)
}
