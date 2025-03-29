package main

import (
	"database/sql"
	"fmt"
	"strconv"

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

func getNotes()([]map[string]string,error){
    query := "SELECT * FROM note"
    rows, err := DB.Query(query)
    if err != nil {
        return nil, fmt.Errorf("failed to query note: %v",err)
    }
    defer rows.Close()
    var  notes []map[string]string
    for rows.Next(){
        var id int64
        var noteName, noteContent string
        if err := rows.Scan(&id,&noteName,&noteContent); err != nil {
            return nil, fmt.Errorf("failed to scan notes: %v",err)
        }
        note := map[string]string{
            "id": strconv.FormatInt(id,10),
            "note_name": noteName,
            "note_content": noteContent,
        }
        notes = append(notes, note)
    }
    if err != nil {
        return nil, fmt.Errorf("error iterating note: %v",err)
    }
    return notes,nil
}


func GetNoteById(id int64)(map[string]string,error){
    query := "SELECT id, note_name,note_content FROM note WHERE id = ?"
    row := DB.QueryRow(query,id)
    var noteId int64
    var noteName, noteContent string
    if err := row.Scan(&noteId, &noteName, &noteContent); err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("task not found")
        }
        return nil, fmt.Errorf("failed to scan task: %v", err)
    }
    targetNote := map[string]string{
        "noteId": strconv.FormatInt(id,10),
        "noteName": noteName,
        "noteContent": noteContent,
    }

    return targetNote, nil
}