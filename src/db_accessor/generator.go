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
	this.CreateTable()
}

func (this *TTestTableGenerator) CreateTable() {
	var transaction, transactionOpeningResult = this.Connection.Begin()
	Assert(transactionOpeningResult)
	defer transaction.Commit()
	transaction.Exec("CREATE TABLE IF NOT EXISTS 'Notes' ('id' INT, 'text' TEXT);")
}
