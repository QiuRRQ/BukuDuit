package usecase

type AuthenticationUseCase struct {
	*UcContract
}

func (uc AuthenticationUseCase) Register(mobilePhone,pin string) (err error){
	userUc := UserUseCase{UcContract:uc.UcContract}
	transaction,err := uc.DB.Begin()
	if err != nil {
		return err
	}
	err = userUc.Add(mobilePhone,pin,transaction)
	if err != nil {
		transaction.Rollback()

		return err
	}

	return nil
}
