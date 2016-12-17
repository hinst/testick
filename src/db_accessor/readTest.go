package db_accessor

import (
	"database/sql"
	"fmt"
	"strings"
)

type TReadTest struct {
	Connection *sql.DB
}

func (this *TReadTest) Run() {
	fmt.Println("Starting test...")
	var loader = CreateTableLoader()
	loader.Connection = this.Connection
	loader.CreateRow = func() IRow { return CreateNoteRow() }
	loader.Query = "select * from Notes"
	var totalCount = 0
	var goldCount = 0
	loader.ReceiveRow = func(iRow IRow) {
		totalCount++
		var row = iRow.(*TNoteRow)
		if strings.Contains(row.Text, "gold") {
			goldCount++
		}
	}
}
