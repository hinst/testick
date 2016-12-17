package db_accessor

import (
	"database/sql"
	"strconv"
	"sync"
)

type TTableLoader struct {
	Query      string
	Connection *sql.DB
	GroupSize  int
	Pool       sync.Pool
	CreateRow  func() IRow
	ReceiveRow func(row IRow)
	RowChannel chan IRow
}

func CreateTableLoader() (result *TTableLoader) {
	result = &TTableLoader{}
	result.GroupSize = 100
	return
}

func (this *TTableLoader) RollPrepare() {
	this.Pool.New = this.GetCreateRowForPool()
	this.RowChannel = make(chan IRow, 8)
	go this.RollReceive()
}

func (this *TTableLoader) RollReceive() {
	for row := range this.RowChannel {
		this.ReceiveRow(row)
	}
}

func (this *TTableLoader) Roll() {
	this.RollPrepare()
	var transaction, transactionBeginResult = this.Connection.Begin()
	Assert(transactionBeginResult)
	defer transaction.Commit()
	var offset = 0
	for this.RollGroup(offset, transaction) {
		offset += this.GroupSize
	}
	this.RollFinalize()
}

func (this *TTableLoader) RollFinalize() {
	close(this.RowChannel)
}

func (this *TTableLoader) RollGroup(offset int, transaction *sql.Tx) bool {
	var result = false
	var rows, queryResult = transaction.Query(this.Query +
		" offset " + strconv.Itoa(offset) + " limit " + strconv.Itoa(this.GroupSize))
	Assert(queryResult)
	defer rows.Close()
	for rows.Next() {
		var row = this.Pool.Get().(IRow)
		row.Load(rows)
		result = true
	}
	return result
}

func (this *TTableLoader) GetCreateRowForPool() func() interface{} {
	return func() interface{} {
		return this.CreateRow()
	}
}
