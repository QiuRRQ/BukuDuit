package models

type TransByMonth struct {
	Sum     int `db:"sum"`
	Monthly int `db:"monthly"`
}
