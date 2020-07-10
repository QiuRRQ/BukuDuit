package models

import "database/sql"

type TransByMonth struct {
	Sum     int           `db:"sum"`
	Monthly sql.NullInt32 `db:"monthly"`
	Weekly  sql.NullInt32 `db:"weekly"`
}
