package usecase

import (
	"bukuduit-go/db/repositories/actions"
	"bukuduit-go/helpers/messages"
	request "bukuduit-go/server/requests"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type BorrowCardUseCase struct {
	*UcContract
}

func (uc BorrowCardUseCase) Read() (res []viewmodel.BorrowCardVM, err error) {
	model := actions.NewBorrowCardModel(uc.DB)
	borrowcards, err := model.Read()
	if err != nil {
		return res, err
	}

	for _, borrowcard := range borrowcards {
		res = append(res, viewmodel.BorrowCardVM{
			Id:         borrowcard.Id,
			MembersId:  borrowcard.MembersId.String,
			TransDate:  borrowcard.TransDate,
			TransMonth: borrowcard.TransMonth,
			TransYear:  borrowcard.TransYear.String,
			Status:     borrowcard.Status.String,
			BookId:     borrowcard.BookId.String,
			Jumlah:     borrowcard.Jumlah,
			BorrowDone: borrowcard.BorrowDone,
			Created_at: borrowcard.Created_at,
			Updated_at: borrowcard.Updated_at,
		})
	}

	return res, err
}

func (uc BorrowCardUseCase) ReadById(ID string) (res []viewmodel.BorrowCardVM, err error) {
	model := actions.NewBorrowCardModel(uc.DB)
	borrowcards, err := model.ReadByID(ID)
	if err != nil {
		return res, err
	}

	for _, borrowcard := range borrowcards {
		res = append(res, viewmodel.BorrowCardVM{
			MembersId:  borrowcard.Id,
			TransDate:  borrowcard.TransDate,
			TransMonth: borrowcard.TransMonth,
			TransYear:  borrowcard.TransYear.String,
			Status:     borrowcard.Status.String,
			BookId:     borrowcard.BookId.String,
			Jumlah:     borrowcard.Jumlah,
			BorrowDone: borrowcard.BorrowDone,
			Created_at: borrowcard.Created_at,
			Updated_at: borrowcard.Updated_at,
		})
	}

	return res, err
}

func (uc BorrowCardUseCase) ReadByBookId(ID, memberID string) (res []viewmodel.BorrowCardVM, err error) {
	model := actions.NewBorrowCardModel(uc.DB)
	borrowcards, err := model.ReadByBookID(ID, memberID)
	if err != nil {
		return res, err
	}

	for _, borrowcard := range borrowcards {
		res = append(res, viewmodel.BorrowCardVM{
			Id:         borrowcard.Id,
			MembersId:  borrowcard.MembersId.String,
			TransDate:  borrowcard.TransDate,
			TransMonth: borrowcard.TransMonth,
			TransYear:  borrowcard.TransYear.String,
			Status:     borrowcard.Status.String,
			BookId:     borrowcard.BookId.String,
			Jumlah:     borrowcard.Jumlah,
			BorrowDone: borrowcard.BorrowDone,
			Created_at: borrowcard.Created_at,
			Updated_at: borrowcard.Updated_at,
		})
	}

	return res, err
}

func (uc BorrowCardUseCase) Edit(input *request.BorrowCardRequest, ID string) (err error) {
	model := actions.NewBorrowCardModel(uc.DB)
	now := time.Now().UTC()
	isExist, err := uc.IsBorrowCardExist(ID)
	if err != nil {
		return err
	}
	if !isExist {
		return errors.New(messages.DataNotFound)
	}

	body := viewmodel.BorrowCardVM{
		Id:         ID,
		MembersId:  input.MembersId,
		TransDate:  input.TransDate,
		TransMonth: input.TransMonth,
		TransYear:  input.TransYear,
		Status:     input.Status,
		BookId:     input.BookId,
		Jumlah:     input.Jumlah,
		BorrowDone: input.BorrowDone,
		Updated_at: now.Format(time.RFC3339),
	}
	_, err = model.Edit(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc BorrowCardUseCase) EditBorrowDone(ID string) (err error) {
	model := actions.NewBorrowCardModel(uc.DB)
	now := time.Now().UTC()
	print(ID)
	isExist, err := uc.IsBorrowCardExist(ID)
	if err != nil {
		return err
	}
	if !isExist {
		return errors.New(messages.DataNotFound)
	}

	body := viewmodel.BorrowCardVM{
		Id:         ID,
		BorrowDone: "1",
		Updated_at: now.Format(time.RFC3339),
	}
	_, err = model.EditBorrowDone(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc BorrowCardUseCase) AddPinjaman(input *request.BorrowCardRequest) (err error) {
	transaction, err := uc.DB.Begin()
	BookUC := BooksUseCase{UcContract: uc.UcContract}
	if err != nil {
		return err
	}
	err = uc.Add(input, transaction)
	if err != nil {
		transaction.Rollback()
		return err
	}

	books, err := BookUC.ReadById(input.BookId)
	if err != nil {
		return err
	}
	book := books[0]
	body := request.BookRequest{
		Stock: book.Stock - input.Jumlah,
	}
	err = BookUC.EditStock(body, book.ID, transaction)
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}

func (uc BorrowCardUseCase) LaporanPeminjaman() (res []viewmodel.RekapPinjamanVM, err error) {
	model := actions.NewBorrowCardModel(uc.DB)
	rekapPInjams, err := model.ReadBorrowedBook()
	if err != nil {
		return res, err
	}

	for _, rekapPinjam := range rekapPInjams {
		res = append(res, viewmodel.RekapPinjamanVM{
			NoMember:     rekapPinjam.NoMember,
			MemberName:   rekapPinjam.MemberName,
			BookTitle:    rekapPinjam.BookTitle,
			JumlahPinjam: rekapPinjam.JumlahPinjam,
			Status:       rekapPinjam.Status,
			Tgl_Pinjam:   rekapPinjam.Tgl_Pinjam,
			Bln_Pinjam:   rekapPinjam.Bln_Pinjam,
			Year_Pinjam:  rekapPinjam.Year_Pinjam,
		})
	}

	return res, err
}

func (uc BorrowCardUseCase) LaporanPengembalian() (res []viewmodel.RekapPengembalianVM, err error) {
	model := actions.NewBorrowCardModel(uc.DB)
	rekapPInjams, err := model.ReadKembalianBook()
	if err != nil {
		return res, err
	}

	for _, rekapPinjam := range rekapPInjams {
		res = append(res, viewmodel.RekapPengembalianVM{
			NoMember:     rekapPinjam.NoMember,
			MemberName:   rekapPinjam.MemberName,
			BookTitle:    rekapPinjam.BookTitle,
			JumlahPinjam: rekapPinjam.JumlahPinjam,
			Status:       "sudah dikembalian",
			Tgl_Pinjam:   rekapPinjam.Tgl_Pinjam,
			Bln_Pinjam:   rekapPinjam.Bln_Pinjam,
			Year_Pinjam:  rekapPinjam.Year_Pinjam,
		})
	}

	return res, err
}

func (uc BorrowCardUseCase) AddPengembalian(input *request.BorrowCardRequest) (err error) {
	transaction, err := uc.DB.Begin()
	BookUC := BooksUseCase{UcContract: uc.UcContract}
	if err != nil {
		return err
	}
	err = uc.Add(input, transaction)
	if err != nil {
		transaction.Rollback()
		return err
	}

	TransPinjaman, err := uc.ReadByBookId(input.BookId, input.MembersId)
	if err != nil {
		return err
	}
	err = uc.EditBorrowDone(TransPinjaman[0].Id)
	if err != nil {
		return err
	}
	books, err := BookUC.ReadById(input.BookId)
	if err != nil {
		return err
	}
	book := books[0]
	body := request.BookRequest{
		Stock: book.Stock + input.Jumlah,
	}
	err = BookUC.EditStock(body, book.ID, transaction)
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}

func (uc BorrowCardUseCase) Add(input *request.BorrowCardRequest, tx *sql.Tx) (err error) {
	model := actions.NewBorrowCardModel(uc.DB)
	now := time.Now().UTC()

	body := viewmodel.BorrowCardVM{
		MembersId:  input.MembersId,
		TransDate:  input.TransDate,
		TransMonth: input.TransMonth,
		TransYear:  input.TransYear,
		Status:     input.Status,
		BookId:     input.BookId,
		Jumlah:     input.Jumlah,
		BorrowDone: input.BorrowDone,
		Created_at: now.Format(time.RFC3339),
		Updated_at: now.Format(time.RFC3339),
	}
	if tx != nil {
		_, err = model.Add(body, tx)
	} else {
		_, err = model.Add(body, nil)
	}

	if err != nil {
		return err
	}

	return nil
}

func (uc BorrowCardUseCase) Delete(ID string) (err error) {
	fmt.Println(ID)
	model := actions.NewBorrowCardModel(uc.DB)
	now := time.Now().UTC()

	isExist, err := uc.IsBorrowCardExist(ID)
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

func (uc BorrowCardUseCase) IsBorrowCardExist(ID string) (res bool, err error) {
	model := actions.NewBorrowCardModel(uc.DB)
	count, err := model.CountByPk(ID)
	if err != nil {
		return res, err
	}

	return count > 0, err
}
