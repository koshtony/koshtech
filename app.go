package main

import (
	"html/template"
	"log"
	"net/http"
)

func mainPage(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

func main() {
	statics := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", statics))
	http.HandleFunc("/", mainPage)
	log.Fatal(http.ListenAndServe(":200", nil))
}
