package db_accessor

import "database/sql"

type TTableLoader struct {
	Query      string
	Connection *sql.DB
	GroupSize  int
}

func (this *TTableLoader) Load() {
}
