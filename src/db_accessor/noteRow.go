package db_accessor

import "database/sql"

type TNoteRow struct {
	Id   int32
	Text string
}

func (this *TNoteRow) Insert(tableName string, transaction *sql.Tx) {
	var _, executionResult = transaction.Exec(
		"insert into "+tableName+"(id, text) values(?, ?)",
		this.Id, this.Text)
	Assert(executionResult)
}
