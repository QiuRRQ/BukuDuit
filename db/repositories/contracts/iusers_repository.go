package contracts

import (
	"bukuduit-go/db/models"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
)

type IUserRepository interface {
	Browse(search, order, sort string, limit, offset int) (data []models.Users, count int, err error)

	ReadByPk(ID string) (data models.Users, err error)

	ReadBy(column, value string) (data models.Users, err error)

	Edit(input viewmodel.UserVm) (res string, err error)

	EditPin(ID, pin, updatedAt string) (res string, err error)

	Add(input viewmodel.UserVm, pin string, tx *sql.Tx) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	CountBy(column, value string) (res int, err error)
}
