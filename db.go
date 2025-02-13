package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
func InitDB(dsn string)error{
    var err error
    DB, err = sql.Open("mysql",dsn)
    if err != nil{
        return fmt.Errorf("failed to open database: %v",err)
    }

    if err = DB.Ping(); err != nil{
        return fmt.Errorf("failed to ping database: %v",err)
    }
    fmt.Println("connected successfully to database")
    return nil
}