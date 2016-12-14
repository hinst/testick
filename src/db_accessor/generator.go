package db_accessor

import "database/sql"

type TTestTableGenerator struct {
	RowCount   int
	Connection *sql.DB
}

func CreateTestTableGenerator() *TTestTableGenerator {
	var result TTestTableGenerator
	result.RowCount = 100 * 1000
	return &result
}

func (this *TTestTableGenerator) Generate() {
	var transaction, transactionOpeningResult = this.Connection.Begin()
	Assert(transactionOpeningResult)
	defer transaction.Commit()
	this.CreateTable(transaction)
}

func (this *TTestTableGenerator) CreateTable(transaction *sql.Tx) {
	var _, executionResult = transaction.Exec("CREATE TABLE IF NOT EXISTS Notes (id INT, text TEXT);")
	Assert(executionResult)
}
