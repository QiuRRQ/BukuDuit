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

type CategoriesUseCase struct {
	*UcContract
}

func (uc CategoriesUseCase) Read() (res []viewmodel.CategoriesVM, err error) {
	model := actions.NewCategoryModel(uc.DB)
	categories, err := model.Read()
	if err != nil {
		return res, err
	}

	for _, category := range categories {
		res = append(res, viewmodel.CategoriesVM{
			ID:         category.Id,
			Name:       category.Name,
			Created_at: category.Created_at,
			UPdated_at: category.Updated_at,
		})
	}

	return res, err
}

func (uc CategoriesUseCase) ReadById(ID string) (res []viewmodel.CategoriesVM, err error) {
	model := actions.NewCategoryModel(uc.DB)
	categories, err := model.ReadByID(ID)
	if err != nil {
		return res, err
	}

	for _, category := range categories {
		res = append(res, viewmodel.CategoriesVM{
			ID:         category.Id,
			Name:       category.Name,
			Created_at: category.Created_at,
			UPdated_at: category.Updated_at,
		})
	}

	return res, err
}

func (uc CategoriesUseCase) Edit(input *request.CategoryRequest, ID string) (err error) {
	model := actions.NewCategoryModel(uc.DB)
	now := time.Now().UTC()
	isExist, err := uc.IsCategoryExist(ID)
	if err != nil {
		return err
	}
	if !isExist {
		return errors.New(messages.DataNotFound)
	}

	body := viewmodel.CategoriesVM{
		ID:         ID,
		Name:       input.Name,
		UPdated_at: now.Format(time.RFC3339),
	}
	_, err = model.Edit(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc CategoriesUseCase) Add(input *request.CategoryRequest) (err error) {
	model := actions.NewCategoryModel(uc.DB)
	now := time.Now().UTC()

	body := viewmodel.CategoriesVM{
		Name:       input.Name,
		UPdated_at: now.Format(time.RFC3339),
		Created_at: now.Format(time.RFC3339),
	}
	_, err = model.Add(body, nil)
	if err != nil {
		return err
	}

	return nil
}

func (uc CategoriesUseCase) Delete(ID string) (err error) {
	fmt.Println(ID)
	model := actions.NewCategoryModel(uc.DB)
	now := time.Now().UTC()

	isExist, err := uc.IsCategoryExist(ID)
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

func (uc CategoriesUseCase) IsCategoryExist(ID string) (res bool, err error) {
	model := actions.NewCategoryModel(uc.DB)
	count, err := model.CountByPk(ID)
	if err != nil {
		return res, err
	}

	return count > 0, err
}
