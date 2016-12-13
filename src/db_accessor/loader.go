package db_accessor

import "database/sql"

type TableLoader struct {
	Query      string
	Connection *sql.DB
}

func (this *TableLoader) Load() {
}
