package main

import (
	"db_accessor"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var app = &db_accessor.TApp{}
	app.Run()
}
