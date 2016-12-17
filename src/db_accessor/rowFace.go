package db_accessor

import "database/sql"

type IRow interface {
	Load(rows *sql.Rows)
}
