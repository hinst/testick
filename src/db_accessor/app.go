package db_accessor

import "database/sql"

type TApp struct {
	Connection *sql.DB
}

func (this *TApp) Run() {
	this.Connect()
	this.Read()
	this.Read()
}

func (this *TApp) Connect() {
	var connection, connectionResult = sql.Open("mysql", "testor:testor@tcp(localhost:3306)/testick")
	Assert(connectionResult)
	this.Connection = connection
}

func (this *TApp) Generate() {
	var generator = CreateTestTableGenerator()
	generator.Connection = this.Connection
	generator.Generate()
}

func (this *TApp) Read() {
	var readTest = CreateReadTest()
	readTest.Connection = this.Connection
	readTest.Run()
}
