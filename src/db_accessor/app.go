package db_accessor

import "database/sql"

type TApp struct {
	Connection *sql.DB
}

func (this *TApp) Run() {
	this.Connect()
	var generator = CreateTestTableGenerator()
	generator.Connection = this.Connection
	generator.Generate()
}

func (this *TApp) Connect() {
	var connection, connectionResult = sql.Open("mysql", "testor:testor@tcp(localhost:3306)/testick")
	Assert(connectionResult)
	this.Connection = connection
}
