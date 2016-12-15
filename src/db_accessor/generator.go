package db_accessor

import (
	"bytes"
	"database/sql"
	"math/rand"
	"time"
)

type TTestTableGenerator struct {
	RowCount   int
	TextSize   int
	Connection *sql.DB
	Random     *rand.Rand
}

func CreateTestTableGenerator() *TTestTableGenerator {
	var result TTestTableGenerator
	result.RowCount = 100 * 1000
	result.TextSize = 1000
	result.Random = rand.New(rand.NewSource(time.Now().UnixNano()))
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

func (this *TTestTableGenerator) GenerateText() string {
	var text bytes.Buffer
	if this.Random.Int31n(10) == 0 {
		text.WriteString("gold ")
	}
	for text.Len() < this.TextSize {
		text.WriteString("text ")
	}
	return text.String()
}

func (this *TTestTableGenerator) GenerateRow(id int32) *TNoteRow {
	var result = CreateNoteRow()
	result.Id = id
	result.Text = this.GenerateText()
	return result
}
