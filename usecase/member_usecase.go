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

type MembersUseCase struct {
	*UcContract
}

func (uc MembersUseCase) Read() (res []viewmodel.MembersVM, err error) {
	model := actions.NewMembersModel(uc.DB)
	members, err := model.Read()
	if err != nil {
		return res, err
	}

	for _, member := range members {
		res = append(res, viewmodel.MembersVM{
			ID:         member.Id,
			NoMember:   member.NoMember,
			Name:       member.Name,
			NoTelp:     member.NoTelp,
			Address:    member.Address,
			City:       member.City,
			Province:   member.Province,
			Member_IMG: member.Member_img,
			Gender:     member.Gender,
			BirthDate:  member.BirthDate,
			BirthMonth: member.BirthMonth,
			BirthYear:  member.BirthYear,
			CreatedAt:  member.CreatedAt,
			UpdatedAt:  member.UpdatedAt,
		})
	}

	return res, err
}

func (uc MembersUseCase) ReadById(ID string) (res []viewmodel.MembersVM, err error) {
	model := actions.NewMembersModel(uc.DB)
	members, err := model.ReadByID(ID)
	if err != nil {
		return res, err
	}

	for _, member := range members {
		res = append(res, viewmodel.MembersVM{
			ID:         member.Id,
			NoMember:   member.NoMember,
			Name:       member.Name,
			NoTelp:     member.NoTelp,
			Address:    member.Address,
			City:       member.City,
			Province:   member.Province,
			Member_IMG: member.Member_img,
			Gender:     member.Gender,
			BirthDate:  member.BirthDate,
			BirthMonth: member.BirthMonth,
			BirthYear:  member.BirthYear,
			CreatedAt:  member.CreatedAt,
			UpdatedAt:  member.UpdatedAt,
		})
	}

	return res, err
}

func (uc MembersUseCase) Edit(input *request.MemberRequest, ID string) (err error) {
	model := actions.NewMembersModel(uc.DB)
	now := time.Now().UTC()
	isExist, err := uc.IsMemberExist(ID)
	if err != nil {
		return err
	}
	if !isExist {
		return errors.New(messages.DataNotFound)
	}

	body := viewmodel.MembersVM{
		ID:         ID,
		NoMember:   input.NoMember,
		Name:       input.Name,
		NoTelp:     input.NoTelp,
		Address:    input.Address,
		City:       input.City,
		Province:   input.NoTelp,
		Member_IMG: input.MemberImg,
		Gender:     input.Gender,
		BirthDate:  input.BirthDate,
		BirthMonth: input.BirthMonth,
		BirthYear:  input.BirthYear,
		UpdatedAt:  now.Format(time.RFC3339),
	}
	_, err = model.Edit(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc MembersUseCase) Add(input *request.MemberRequest) (err error) {
	model := actions.NewMembersModel(uc.DB)
	now := time.Now().UTC()

	body := viewmodel.MembersVM{
		NoMember:   input.NoMember,
		Name:       input.Name,
		NoTelp:     input.NoTelp,
		Address:    input.Address,
		City:       input.City,
		Province:   input.NoTelp,
		Member_IMG: input.MemberImg,
		Gender:     input.Gender,
		BirthDate:  input.BirthDate,
		BirthMonth: input.BirthMonth,
		BirthYear:  input.BirthYear,
		UpdatedAt:  now.Format(time.RFC3339),
		CreatedAt:  now.Format(time.RFC3339),
	}
	_, err = model.Add(body, nil)
	if err != nil {
		return err
	}

	return nil
}

func (uc MembersUseCase) Delete(ID string) (err error) {
	fmt.Println(ID)
	model := actions.NewMembersModel(uc.DB)
	now := time.Now().UTC()

	isExist, err := uc.IsMemberExist(ID)
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

func (uc MembersUseCase) IsMemberExist(ID string) (res bool, err error) {
	model := actions.NewMembersModel(uc.DB)
	count, err := model.CountByPk(ID)
	if err != nil {
		return res, err
	}

	return count > 0, err
}
