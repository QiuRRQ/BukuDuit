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

type AuthorsUseCase struct {
	*UcContract
}

func (uc AuthorsUseCase) Read() (res []viewmodel.AuthorsVM, err error) {
	model := actions.NewAuthorsModel(uc.DB)
	authors, err := model.Read()
	if err != nil {
		return res, err
	}

	for _, author := range authors {
		res = append(res, viewmodel.AuthorsVM{
			ID:         author.Id,
			Name:       author.Name.String,
			Address:    author.Address,
			City:       author.City,
			Province:   author.Province.String,
			PostalCode: author.PostalCode.String,
			NoTelp:     author.NoTelp.String,
			Created_at: author.Created_at,
			UPdated_at: author.Updated_at,
		})
	}

	return res, err
}

func (uc AuthorsUseCase) ReadById(ID string) (res []viewmodel.AuthorsVM, err error) {
	model := actions.NewAuthorsModel(uc.DB)
	authors, err := model.ReadByID(ID)
	if err != nil {
		return res, err
	}

	for _, author := range authors {
		res = append(res, viewmodel.AuthorsVM{
			ID:         author.Id,
			Name:       author.Name.String,
			Address:    author.Address,
			City:       author.City,
			Province:   author.Province.String,
			PostalCode: author.PostalCode.String,
			NoTelp:     author.NoTelp.String,
			Created_at: author.Created_at,
			UPdated_at: author.Updated_at,
		})
	}

	return res, err
}

func (uc AuthorsUseCase) Edit(input *request.AuthorRequest, ID string) (err error) {
	model := actions.NewAuthorsModel(uc.DB)
	now := time.Now().UTC()
	isExist, err := uc.IsAuthorExist(ID)
	if err != nil {
		return err
	}
	if !isExist {
		return errors.New(messages.DataNotFound)
	}

	body := viewmodel.AuthorsVM{
		ID:         ID,
		Name:       input.AuthorName,
		Address:    input.AuthorAddress,
		City:       input.AuthorCity,
		Province:   input.AuthorProvince,
		PostalCode: input.PostalCode,
		NoTelp:     input.NoTelp,
		UPdated_at: now.Format(time.RFC3339),
	}
	_, err = model.Edit(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc AuthorsUseCase) Add(input *request.AuthorRequest) (err error) {
	model := actions.NewAuthorsModel(uc.DB)
	now := time.Now().UTC()

	body := viewmodel.AuthorsVM{
		Name:       input.AuthorName,
		Address:    input.AuthorAddress,
		City:       input.AuthorCity,
		Province:   input.AuthorProvince,
		PostalCode: input.PostalCode,
		NoTelp:     input.NoTelp,
		UPdated_at: now.Format(time.RFC3339),
		Created_at: now.Format(time.RFC3339),
	}
	_, err = model.Add(body, nil)
	if err != nil {
		return err
	}

	return nil
}

func (uc AuthorsUseCase) Delete(ID string) (err error) {
	fmt.Println(ID)
	model := actions.NewAuthorsModel(uc.DB)
	now := time.Now().UTC()

	isExist, err := uc.IsAuthorExist(ID)
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

func (uc AuthorsUseCase) IsAuthorExist(ID string) (res bool, err error) {
	model := actions.NewAuthorsModel(uc.DB)
	count, err := model.CountByPk(ID)
	if err != nil {
		return res, err
	}

	return count > 0, err
}
