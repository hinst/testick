package db_accessor

import "fmt"
import "time"

func WriteLog(s string) {
	s = time.Now().Format("") + "" + s
	fmt.Println(s)
}
