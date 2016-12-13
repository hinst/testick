package db_accessor

func Assert(e error) {
	if e != nil {
		panic(e)
	}
}
