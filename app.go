package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	"os"
)

// main page route
func mainPage(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		req.ParseForm()
		name := req.Form["name"]
		subject := req.Form["subject"]
		email := req.Form["email"]
		msg := req.Form["message"]
		sendMail(name[0], email[0], subject[0], msg[0])
	}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)

}

// send
func sendMail(name string, email string, sub string, body string) {
	from := "tonykosh4@gmail.com"
	password := "swvkfhtazuofijpp"
	sendTo := []string{from}
	host := "smtp.gmail.com"
	port := "587"
	msg := []byte(name + "\n" + email + "\n" + sub + "\n" + body)
	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail(host+":"+port, auth, from, sendTo, msg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("message sent successfully")

}

func main() {
	port := os.Getenv("PORT")
	statics := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", statics))
	http.HandleFunc("/", mainPage)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
