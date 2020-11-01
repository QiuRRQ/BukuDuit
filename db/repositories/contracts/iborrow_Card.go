package contracts

import (
	"bukuduit-go/db/models"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
)

type IBorrowCard interface {
	Read() (data []models.BorrowCard, err error)

	ReadByID(ID string) (data []models.BorrowCard, err error)

	ReadByBookID(ID, memberID string) (data []models.BorrowCard, err error)

	Edit(body viewmodel.BorrowCardVM) (res string, err error)

	EditBorrowDone(body viewmodel.BorrowCardVM) (res string, err error)

	Add(body viewmodel.BorrowCardVM, tx *sql.Tx) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	CountByPk(ID string) (res int, err error)

	CountBy(column, value string) (res int, err error)

	ReadBorrowedBook() (data []models.RekapPinjaman, err error)

	ReadKembalianBook() (data []models.RekapPengembalian, err error)
}
