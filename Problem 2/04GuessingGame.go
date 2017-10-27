package main

import (													//Imports
	"fmt"
	"net/http"
	"text/template"
)

type TemplateData struct {
	Message string
}

func templateHandler(w http.ResponseWriter, r *http.Request) {	//Handle Http requests

	t, _ := template.ParseFiles("template/guess.html")											//Parse the template File
	t.Execute(w, TemplateData{Message: "Guess a Number Between 1 and 20"})						// Execute the Tmpl file

}

func cookieHandler(w http.ResponseWriter, r *http.Request) {
	
	randomNum := 0

	// Try to read the cookie.
	var cookie, err = r.Cookie("randomNum")
	
	if err == nil {
		// If we could read it, try to convert its value to an int.
		randomNum = 10
	}

	// Create a cookie instance and set the cookie.
	cookie = &http.Cookie{

		Name: "randomNum",
		Value: "Test",
	}

	http.SetCookie(w, cookie)

	fmt.Fprintf(w, "Random Number %d .", randomNum)

}

func main() {

	http.Handle("/", http.FileServer(http.Dir("./")))		//Handle http request
	http.HandleFunc("/guess",templateHandler)				//handle requests for templates
	//http.HandleFunc("/guess", cookieHandler)
	http.ListenAndServe(":8080", nil)						//Listen and report from port 8080

}