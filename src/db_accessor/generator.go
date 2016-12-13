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

}
