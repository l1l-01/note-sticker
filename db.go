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

func addNote(noteName, noteContent string)error{
    query := "INSERT INTO note (note_name,note_content) VALUES (?,?)"
    _, err := DB.Exec(query,noteName,noteContent)
    if err != nil {
        return fmt.Errorf("failed to insert note: %v",err)
    }
    fmt.Println("note was inserted successfully")
    return nil
}