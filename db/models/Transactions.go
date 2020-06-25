package models

import (
	"database/sql"
)

type Transactions struct {
	ID               string         `db:"id"`
	Reference_Id     string         `db:"reference_id"`
	IDShop           string         `db:"shop_id"`
	Amount           sql.NullInt32  `db:"amount"`
	Description      string         `db:"description"`
	Image            string         `db:"image"`
	Type             string         `db:"type"`
	Transaction_Date sql.NullString `db:"transaction_date"`
	Created_at       string         `db:"created_at"`
	Update_at        sql.NullString `db:"update_at"`
	Deleted_at       sql.NullString `db:"deleted_at"`
}
