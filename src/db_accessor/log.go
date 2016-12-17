package db_accessor

import "fmt"
import "time"

func WriteLog(s string) {
	s = TimeToString(time.Now()) + ": " + s
	fmt.Println(s)
}
