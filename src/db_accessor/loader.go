package db_accessor

import (
	"database/sql"
	"strconv"
)

type TTableLoader struct {
	Query      string
	Connection *sql.DB
	GroupSize  int
}

func CreateTableLoader() (result *TTableLoader) {
	result = &TTableLoader{}
	result.GroupSize = 100
	return
}

func (this *TTableLoader) Roll() {
	var transaction, transactionBeginResult = this.Connection.Begin()
	Assert(transactionBeginResult)
	defer transaction.Commit()
}

func (this *TTableLoader) RollGroup(offset int, transaction *sql.Tx) bool {
	var result = false
	var rows, queryResult = transaction.Query(this.Query +
		" offset " + strconv.Itoa(offset) + " limit " + strconv.Itoa(this.GroupSize))
	Assert(queryResult)
	defer rows.Close()
	for rows.Next() {
		result = true
	}
	return result
}
