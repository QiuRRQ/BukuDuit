package contracts

import (
	"bukuduit-go/db/models"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
)

type IMembers interface {
	Read() (data []models.Members, err error)

	ReadByID(ID string) (data []models.Members, err error)

	Edit(body viewmodel.MembersVM) (res string, err error)

	Add(body viewmodel.MembersVM, tx *sql.Tx) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	CountByPk(ID string) (res int, err error)

	CountBy(column, value string) (res int, err error)
}
