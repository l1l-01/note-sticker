package main

import (
	"fmt"
	"net/http"
)

func pageHandler(w http.ResponseWriter,r *http.Request){
	http.ServeFile(w,r,"index.html")
}

func main(){
	fs := http.FileServer(http.Dir("static/css"))
	http.Handle("/static/css/", http.StripPrefix("/static/css/", fs))

	fjs := http.FileServer(http.Dir("static/javascript"))
	http.Handle("/static/javascript/",http.StripPrefix("/static/javascript/",fjs))

	fimg := http.FileServer(http.Dir("static/images"))
	http.Handle("/static/images/",http.StripPrefix("/static/images/",fimg))

	http.HandleFunc("/",pageHandler)
	fmt.Println("Server is runing on http//:localhost8080")
	if err := http.ListenAndServe(":8080",nil); err != nil {
		fmt.Println("Error starting server: ",err)
	}
}