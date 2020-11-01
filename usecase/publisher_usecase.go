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

type PublishersUseCase struct {
	*UcContract
}

func (uc PublishersUseCase) Read() (res []viewmodel.PublishersVM, err error) {
	model := actions.NewPublishersModel(uc.DB)
	publishers, err := model.Read()
	if err != nil {
		return res, err
	}

	for _, publisher := range publishers {
		res = append(res, viewmodel.PublishersVM{
			ID:         publisher.Id,
			Name:       publisher.Name.String,
			Address:    publisher.Address,
			City:       publisher.City,
			Province:   publisher.Province.String,
			PostalCode: publisher.PostalCode.String,
			NoTelp:     publisher.NoTelp.String,
			Created_at: publisher.Created_at,
			UPdated_at: publisher.Updated_at,
		})
	}

	return res, err
}

func (uc PublishersUseCase) ReadById(ID string) (res []viewmodel.PublishersVM, err error) {
	model := actions.NewPublishersModel(uc.DB)
	publishers, err := model.ReadByID(ID)
	if err != nil {
		return res, err
	}

	for _, publisher := range publishers {
		res = append(res, viewmodel.PublishersVM{
			ID:         publisher.Id,
			Name:       publisher.Name.String,
			Address:    publisher.Address,
			City:       publisher.City,
			Province:   publisher.Province.String,
			PostalCode: publisher.PostalCode.String,
			NoTelp:     publisher.NoTelp.String,
			Created_at: publisher.Created_at,
			UPdated_at: publisher.Updated_at,
		})
	}

	return res, err
}

func (uc PublishersUseCase) Edit(input *request.PublisherRequest, ID string) (err error) {
	model := actions.NewPublishersModel(uc.DB)
	now := time.Now().UTC()
	isExist, err := uc.IsPublisherExist(ID)
	if err != nil {
		return err
	}
	if !isExist {
		return errors.New(messages.DataNotFound)
	}

	body := viewmodel.PublishersVM{
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

func (uc PublishersUseCase) Add(input *request.PublisherRequest) (err error) {
	model := actions.NewPublishersModel(uc.DB)
	now := time.Now().UTC()

	body := viewmodel.PublishersVM{
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

func (uc PublishersUseCase) Delete(ID string) (err error) {
	fmt.Println(ID)
	model := actions.NewPublishersModel(uc.DB)
	now := time.Now().UTC()

	isExist, err := uc.IsPublisherExist(ID)
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

func (uc PublishersUseCase) IsPublisherExist(ID string) (res bool, err error) {
	model := actions.NewPublishersModel(uc.DB)
	count, err := model.CountByPk(ID)
	if err != nil {
		return res, err
	}

	return count > 0, err
}
