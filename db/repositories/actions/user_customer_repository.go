package actions

import (
	"bukuduit-go/db/models"
	"bukuduit-go/db/repositories/contracts"
	"bukuduit-go/helpers/datetime"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
	"time"
)

type UserCustomerRepository struct {
	DB *sql.DB
}

func NewUserCustomerModel(DB *sql.DB) contracts.IUserCustomerRepository {
	return UserCustomerRepository{DB: DB}
}

func (repository UserCustomerRepository) BrowseByBusiness(businessID string) (data []models.UserCustomers, err error) {
	statement := `select * from "user_customers" where "business_id"=$1 and "deleted_at" is null`
	rows, err := repository.DB.Query(statement, businessID)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.UserCustomers{}
		err = rows.Scan(&dataTemp.ID, &dataTemp.FullName, &dataTemp.MobilePhone, &dataTemp.BusinessID, &dataTemp.PaymentDate, &dataTemp.CreatedAt, &dataTemp.UpdatedAt, &dataTemp.DeletedAt)
		if err != nil {
			return data, err
		}

		data = append(data, dataTemp)
	}

	return data, err
}

func (repository UserCustomerRepository) Read(ID string) (data models.UserCustomers, err error) {
	statement := `select * from "user_customers" where "id"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, ID).Scan(&data.ID, &data.FullName, &data.MobilePhone, &data.BusinessID, &data.PaymentDate, &data.CreatedAt, &data.UpdatedAt, &data.DeletedAt)
	if err != nil {
		return data, err
	}

	return data, err
}

func (repository UserCustomerRepository) ReadByPhone(phone string) (data models.UserCustomers, err error) {
	statement := `select * from "user_customers" where "mobile_phone"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, phone).Scan(&data.ID, &data.FullName, &data.MobilePhone, &data.BusinessID, &data.PaymentDate, &data.CreatedAt, &data.UpdatedAt, &data.DeletedAt)
	if err != nil {
		return data, err
	}

	return data, err
}

func (repository UserCustomerRepository) EditDebt(ID, updatedAt string, debt int32) (res string, err error) {
	statement := `update "user_customers" set "debt"=$1, "updated_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(
		statement,
		debt, datetime.StrParseToTime(updatedAt, time.RFC3339), ID,
	).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository UserCustomerRepository) Add(body viewmodel.UserCustomerVm, businessID string) (res string, err error) {
	statement := `insert into "user_customers" ("full_name","mobile_phone","business_id","created_at","updated_at") values($1,$2,$3,$4,$5) returning "id"`
	err = repository.DB.QueryRow(statement, body.FullName, body.MobilePhone, businessID, datetime.StrParseToTime(body.CreatedAt, time.RFC3339), datetime.StrParseToTime(body.DeletedAt, time.RFC3339)).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository UserCustomerRepository) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "user_customers" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository UserCustomerRepository) CountByPk(ID string) (res int, err error) {
	statement := `select count("id") from "user_customers" where "id"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository UserCustomerRepository) CountBy(column, value string) (res int, err error) {
	statement := `select count("id") from "user_customers" where ` + column + `=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
