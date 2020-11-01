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

type BookHasCategoryUseCase struct {
	*UcContract
}

func (uc BookHasCategoryUseCase) Read() (res []viewmodel.BooksHasCategoryVM, err error) {
	model := actions.NewBookCategoryModel(uc.DB)
	bookcats, err := model.Read()
	if err != nil {
		return res, err
	}

	for _, bookcat := range bookcats {
		res = append(res, viewmodel.BooksHasCategoryVM{
			ID:         bookcat.Id,
			BookID:     bookcat.BooksId,
			CategoryId: bookcat.CategoryID,
			Created_at: bookcat.Created_at,
			UPdated_at: bookcat.Updated_at,
		})
	}

	return res, err
}

func (uc BookHasCategoryUseCase) ReadById(ID string) (res []viewmodel.BooksHasCategoryVM, err error) {
	model := actions.NewBookCategoryModel(uc.DB)
	bookcats, err := model.ReadByID(ID)
	if err != nil {
		return res, err
	}

	for _, bookcat := range bookcats {
		res = append(res, viewmodel.BooksHasCategoryVM{
			ID:         bookcat.Id,
			BookID:     bookcat.BooksId,
			CategoryId: bookcat.CategoryID,
			Created_at: bookcat.Created_at,
			UPdated_at: bookcat.Updated_at,
		})
	}

	return res, err
}

func (uc BookHasCategoryUseCase) Edit(input *request.BookHasCategoryRequest, ID string) (err error) {
	model := actions.NewBookCategoryModel(uc.DB)
	now := time.Now().UTC()
	isExist, err := uc.IsBookHasCategoryExist(ID)
	if err != nil {
		return err
	}
	if !isExist {
		return errors.New(messages.DataNotFound)
	}

	body := viewmodel.BooksHasCategoryVM{
		ID:         ID,
		BookID:     input.BookId,
		CategoryId: input.CategoryId,
		UPdated_at: now.Format(time.RFC3339),
	}
	_, err = model.Edit(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc BookHasCategoryUseCase) Add(input *request.BookHasCategoryRequest) (err error) {
	model := actions.NewBookCategoryModel(uc.DB)
	now := time.Now().UTC()

	body := viewmodel.BooksHasCategoryVM{
		BookID:     input.BookId,
		CategoryId: input.CategoryId,
		UPdated_at: now.Format(time.RFC3339),
		Created_at: now.Format(time.RFC3339),
	}
	_, err = model.Add(body, nil)
	if err != nil {
		return err
	}

	return nil
}

func (uc BookHasCategoryUseCase) Delete(ID string) (err error) {
	fmt.Println(ID)
	model := actions.NewBookCategoryModel(uc.DB)
	now := time.Now().UTC()

	isExist, err := uc.IsBookHasCategoryExist(ID)
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

func (uc BookHasCategoryUseCase) IsBookHasCategoryExist(ID string) (res bool, err error) {
	model := actions.NewBookCategoryModel(uc.DB)
	count, err := model.CountByPk(ID)
	if err != nil {
		return res, err
	}

	return count > 0, err
}
