package contracts

import (
	"bukuduit-go/db/models"
	"bukuduit-go/usecase/viewmodel"
	// "database/sql"
)

type IUserCustomerRepository interface {
	BrowseByBusiness(businessID string) (data []models.UserCustomers, err error)

	Read(ID string) (data models.UserCustomers, err error)

	EditDebt(ID, updatedAt string, debt int32) (res string, err error)

	Add(body viewmodel.UserCustomerVm, businessID string) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	CountByPk(ID string) (res int, err error)
}
