package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main(){
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test")
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = db.Exec("select  * from cc")
	if err != nil {
		fmt.Println(err.Error())
	}
}
