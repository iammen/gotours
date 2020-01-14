package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "root:ultraman7@tcp(localhost:3306)/4xbuddy_db?parseTime=true&charset=utf8")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	var myTime time.Time
	rows, err := db.Query("select created_at from users where user_id = 1")
	defer rows.Close()
	if rows.Next() {
		if err = rows.Scan(&myTime); err != nil {
			panic(err)
		}
	}

	fmt.Println(myTime)
}
