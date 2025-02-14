package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func HomeHandler(w http.ResponseWriter,r *http.Request){
	notes, err := getNotes()
	if err != nil {
		http.Error(w,"failed to fetch notes",http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w,notes)
}

func addNoteHandler(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w,"invalid request method",http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w,"failed to parse form: %v",http.StatusBadRequest)
		return
	}
	noteName := r.FormValue("note-name")
	noteContent := r.FormValue("note-content")
	if err := addNote(noteName,noteContent); err != nil {
		http.Error(w,"failed to add note",http.StatusInternalServerError)
		return
	}
	http.Redirect(w,r,"/msg=?note+was+added+successfully!",http.StatusSeeOther)
}

func main(){
	dsn := "root:@(localhost:3306)/notes?parseTime=true"
	if err := InitDB(dsn); err != nil {
		log.Fatal(err)
	}
	defer DB.Close()
	fs := http.FileServer(http.Dir("static/css"))
	http.Handle("/static/css/", http.StripPrefix("/static/css/", fs))

	fjs := http.FileServer(http.Dir("static/javascript"))
	http.Handle("/static/javascript/",http.StripPrefix("/static/javascript/",fjs))

	fimg := http.FileServer(http.Dir("static/images"))
	http.Handle("/static/images/",http.StripPrefix("/static/images/",fimg))

	http.HandleFunc("/add-note",addNoteHandler)

	http.HandleFunc("/",HomeHandler)
	fmt.Println("Server is runing on http//:localhost8080")
	if err := http.ListenAndServe(":8080",nil); err != nil {
		fmt.Println("Error starting server: ",err)
	}
}