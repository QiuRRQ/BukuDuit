package models

import "database/sql"

type Members struct {
	Id         string         `db:"id"`
	NoMember   string         `db:"no_member"`
	Name       string         `db:"name"`
	NoTelp     string         `db:"no_telp"`
	Address    string         `db:"address"`
	City       string         `db:"city"`
	Province   string         `db:"province"`
	Member_img string         `db:"member_img"`
	Gender     string         `db:"gender"`
	BirthDate  string         `db:"birth_date"`
	BirthMonth string         `db:"birth_month'`
	BirthYear  string         `db:"birth_year"`
	CreatedAt  string         `db:"created_at"`
	UpdatedAt  string         `db:"updated_at"`
	DeletedAt  sql.NullString `db:"deleted_at"`
}
