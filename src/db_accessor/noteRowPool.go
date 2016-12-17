package db_accessor

import "sync"

type TNoteRowPool struct {
	sync.Pool
}

func CreateNoteRowPool() *TNoteRowPool {
	var result TNoteRowPool
	result.New = func() interface{} {
		return CreateNoteRow()
	}
	return &result
}

func (this *TNoteRowPool) Get() *TNoteRow {
	return this.Pool.Get().(*TNoteRow)
}

func (this *TNoteRowPool) Put(a *TNoteRow) {
	this.Pool.Put(a)
}
