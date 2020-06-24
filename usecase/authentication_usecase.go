package usecase

import (
	"bukuduit-go/usecase/viewmodel"
	"os"
	"time"

	uuid "github.com/satori/go.uuid"
)

type AuthenticationUseCase struct {
	*UcContract
}

// UpdateSessionLogin ...
func (uc AuthenticationUseCase) UpdateSessionLogin(ID string) (res string, err error) {
	value := uuid.NewV4().String()
	exp := os.Getenv("SESSION_EXP")
	key := "session-" + ID
	resSession := viewmodel.UserSessionVm{}
	resSession.Session = value

	uc.StoreToRedistWithExpired(key, resSession, exp)

	return value, err
}

// GenerateJwtToken ...
func (uc AuthenticationUseCase) GenerateJwtToken(jwePayload, mobilePhone, session string) (token, refreshToken, expTokenAt, expRefreshTokenAt string, err error) {
	token, expTokenAt, err = uc.JwtCred.GetToken(session, mobilePhone, jwePayload)
	if err != nil {
		return token, refreshToken, expTokenAt, expRefreshTokenAt, err
	}

	refreshToken, expRefreshTokenAt, err = uc.JwtCred.GetRefreshToken(session, mobilePhone, jwePayload)
	if err != nil {
		return token, refreshToken, expTokenAt, expRefreshTokenAt, err
	}

	return token, refreshToken, expTokenAt, expRefreshTokenAt, err
}

func (uc AuthenticationUseCase) GenerateTokenByOtp(key, otp string) (res viewmodel.UserJwtTokenVm, err error) {
	otpUc := OtpUseCase{UcContract: uc.UcContract}
	otpRes, err := otpUc.SubmitOtp(key, otp)
	if err != nil {
		return res, err
	}

	jwePayload, _ := uc.Jwe.GenerateJwePayload(otpRes.MobilePhone)
	session, _ := uc.UpdateSessionLogin(otpRes.MobilePhone)
	token, refreshToken, tokenExpiredAt, refreshTokenExpiredAt, err := uc.GenerateJwtToken(jwePayload, otpRes.MobilePhone, session)

	res = viewmodel.UserJwtTokenVm{
		Token:           token,
		ExpTime:         tokenExpiredAt,
		RefreshToken:    refreshToken,
		ExpRefreshToken: refreshTokenExpiredAt,
	}

	return res, err
}

func (uc AuthenticationUseCase) Register(mobilePhone, pin, shopName string) (err error) {
	userUc := UserUseCase{UcContract: uc.UcContract}
	businessCardUc := BusinessCardUseCase{UcContract: uc.UcContract}
	now := time.Now().UTC()
	transaction, err := uc.DB.Begin()
	if err != nil {
		return err
	}
	userID, err := userUc.Add(mobilePhone, pin, transaction)
	if err != nil {
		transaction.Rollback()

		return err
	}
	err = businessCardUc.Register(userID, shopName, now.Format(time.RFC3339), now.Format(time.RFC3339), transaction)
	if err != nil {
		transaction.Rollback()

		return err
	}
	transaction.Commit()

	return nil
}
