package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func mainPage(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

func main() {
	port := os.Getenv("PORT")
	statics := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", statics))
	http.HandleFunc("/", mainPage)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
