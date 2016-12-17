package db_accessor

import (
	"database/sql"
	"sync"
)

type TTableLoader struct {
	Query          string
	Connection     *sql.DB
	GroupSize      int
	Pool           sync.Pool
	CreateRow      func() IRow
	ReceiveRow     func(row IRow)
	RowChannel     chan IRow
	ReceiverWaiter sync.WaitGroup
}

func CreateTableLoader() (result *TTableLoader) {
	result = &TTableLoader{}
	return
}

func (this *TTableLoader) RollPrepare() {
	this.Pool.New = this.GetCreateRowForPool()
	this.RowChannel = make(chan IRow, 64)
	go this.RollReceive()
	this.ReceiverWaiter.Add(1)
}

func (this *TTableLoader) RollReceive() {
	for row := range this.RowChannel {
		this.ReceiveRow(row)
	}
	this.ReceiverWaiter.Done()
}

func (this *TTableLoader) Roll() {
	this.RollPrepare()
	var transaction, transactionBeginResult = this.Connection.Begin()
	Assert(transactionBeginResult)
	defer transaction.Commit()
	this.RollGroup(transaction)
	this.RollFinalize()
}

func (this *TTableLoader) RollFinalize() {
	close(this.RowChannel)
	this.ReceiverWaiter.Wait()
}

func (this *TTableLoader) RollGroup(transaction *sql.Tx) bool {
	var result = false
	WriteLog(this.Query)
	var rows, queryResult = transaction.Query(this.Query)
	Assert(queryResult)
	defer rows.Close()
	for rows.Next() {
		var row = this.Pool.Get().(IRow)
		row.Load(rows)
		this.RowChannel <- row
		result = true
	}
	return result
}

func (this *TTableLoader) GetCreateRowForPool() func() interface{} {
	return func() interface{} {
		return this.CreateRow()
	}
}
