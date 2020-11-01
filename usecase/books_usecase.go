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

type BooksUseCase struct {
	*UcContract
}

func (uc BooksUseCase) Read() (res []viewmodel.BooksVM, err error) {
	model := actions.NewBooksModel(uc.DB)
	books, err := model.Read()
	if err != nil {
		return res, err
	}

	for _, book := range books {
		res = append(res, viewmodel.BooksVM{
			ID:           book.Id,
			Tittle:       book.Title.String,
			Publisher_id: book.Publisher_id,
			Authors_id:   book.Authors_id,
			Book_img:     book.Book_img.String,
			Stock:        book.Stock.Int64,
			Created_at:   book.Created_at,
			UPdated_at:   book.Updated_at,
		})
	}

	return res, err
}

func (uc BooksUseCase) ReadById(ID string) (res []viewmodel.BooksVM, err error) {
	model := actions.NewBooksModel(uc.DB)
	books, err := model.ReadByID(ID)
	if err != nil {
		return res, err
	}

	for _, book := range books {
		res = append(res, viewmodel.BooksVM{
			ID:           book.Id,
			Tittle:       book.Title.String,
			Publisher_id: book.Publisher_id,
			Authors_id:   book.Authors_id,
			Book_img:     book.Book_img.String,
			Stock:        book.Stock.Int64,
			Created_at:   book.Created_at,
			UPdated_at:   book.Updated_at,
		})
	}

	return res, err
}

func (uc BooksUseCase) Edit(input *request.BookRequest, ID string) (err error) {
	model := actions.NewBooksModel(uc.DB)
	now := time.Now().UTC()
	isExist, err := uc.IsBookExist(ID)
	if err != nil {
		return err
	}
	if !isExist {
		return errors.New(messages.DataNotFound)
	}

	body := viewmodel.BooksVM{
		ID:           ID,
		Tittle:       input.Title,
		Publisher_id: input.Publisher_id,
		Authors_id:   input.Author_id,
		Book_img:     input.Book_img,
		Stock:        input.Stock,
		UPdated_at:   now.Format(time.RFC3339),
	}
	_, err = model.Edit(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc BooksUseCase) EditStockRequest(input *request.BookRequest, ID string) (err error) {
	model := actions.NewBooksModel(uc.DB)
	now := time.Now().UTC()
	isExist, err := uc.IsBookExist(ID)
	if err != nil {
		return err
	}
	if !isExist {
		return errors.New(messages.DataNotFound)
	}

	body := viewmodel.BooksVM{
		ID:         ID,
		Stock:      input.Stock,
		UPdated_at: now.Format(time.RFC3339),
	}
	_, err = model.EditStock(body, nil)
	if err != nil {
		return err
	}

	return nil
}

func (uc BooksUseCase) EditStock(input request.BookRequest, ID string, tx *sql.Tx) (err error) {
	model := actions.NewBooksModel(uc.DB)
	now := time.Now().UTC()
	isExist, err := uc.IsBookExist(ID)
	if err != nil {
		return err
	}
	if !isExist {
		return errors.New(messages.DataNotFound)
	}

	body := viewmodel.BooksVM{
		ID:         ID,
		Stock:      input.Stock,
		UPdated_at: now.Format(time.RFC3339),
	}
	_, err = model.EditStock(body, tx)
	if err != nil {
		print(1)
		return err
	}

	return nil
}

func (uc BooksUseCase) Add(input *request.BookRequest, tx *sql.Tx) (res viewmodel.AddBookResVM, err error) {
	model := actions.NewBooksModel(uc.DB)
	now := time.Now().UTC()

	body := viewmodel.BooksVM{
		Tittle:       input.Title,
		Publisher_id: input.Publisher_id,
		Authors_id:   input.Author_id,
		Book_img:     input.Book_img,
		Stock:        input.Stock,
		Created_at:   now.Format(time.RFC3339),
		UPdated_at:   now.Format(time.RFC3339),
	}
	BookID := ""
	if tx != nil {
		BookID, err = model.Add(body, tx)
	} else {
		BookID, err = model.Add(body, nil)
	}

	if err != nil {
		return res, err
	}
	res = viewmodel.AddBookResVM{
		ID:         BookID,
		Created_at: now.Format(time.RFC3339),
		UPdated_at: now.Format(time.RFC3339),
	}
	return res, nil
}

// func (uc BooksUseCase) AddNewBookPcs(input *request.BookRequest) (err error) {
// 	model := actions.NewBooksModel(uc.DB)
// 	now := time.Now().UTC()
// 	transaction, err := uc.DB.Begin()
// 	if err != nil {
// 		return err
// 	}
// 	body := viewmodel.BooksVM{
// 		Tittle:       input.Title,
// 		Publisher_id: input.Publisher_id,
// 		Authors_id:   input.Author_id,
// 		Book_img:     input.Book_img,
// 		Stock:        input.Stock,
// 		Created_at:   now.Format(time.RFC3339),
// 		UPdated_at:   now.Format(time.RFC3339),
// 	}
// 	BooksID, err = model.Add(body, transaction)
// 	if err != nil {
// 		transaction.Rollback()
// 		return err
// 	}
// 	body = viewmodel.BooksVM{
// 		ID:           BooksID,
// 		Tittle:       input.Title,
// 		Publisher_id: input.Publisher_id,
// 		Authors_id:   input.Author_id,
// 		Book_img:     input.Book_img,
// 		Stock:        input.Stock,
// 		Created_at:   now.Format(time.RFC3339),
// 		UPdated_at:   now.Format(time.RFC3339),
// 	}
// 	_, err = model.AddBarcode(body, transaction)
// 	if err != nil {
// 		transaction.Rollback()
// 		return err
// 	}
// 	transaction.Commit()
// 	return nil
// }

func (uc BooksUseCase) Delete(ID string) (err error) {
	fmt.Println(ID)
	model := actions.NewBooksModel(uc.DB)
	now := time.Now().UTC()

	isExist, err := uc.IsBookExist(ID)
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

func (uc BooksUseCase) IsBookExist(ID string) (res bool, err error) {
	model := actions.NewBooksModel(uc.DB)
	count, err := model.CountByPk(ID)
	if err != nil {
		return res, err
	}

	return count > 0, err
}
