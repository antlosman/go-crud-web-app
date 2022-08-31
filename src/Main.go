package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func choice(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("assets/index.html"))
	tmpl.Execute(w, nil)
}
func getChoice(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	option := r.Form["Choice"][0] // map of string and []string
	if option == "Create" {
		tmpl := template.Must(template.ParseFiles("assets/Form.html"))
		//tmpl = template.Must()
		tmpl.Execute(w, nil)
	} else if option == "Search" {
		tmpl := template.Must(template.ParseFiles("assets/searchForm.html"))
		tmpl.Execute(w, nil)
	} else if option == "Show" {
		app, e := getAllData()
		if e != nil {
			fmt.Println("Error Occurred in form parsing")
			tmpl := template.Must(template.ParseFiles("assets/Error.html"))
			tmpl.Execute(w, e)
		}
		tmpl := template.Must(template.ParseFiles("assets/showAll.html"))
		fmt.Println(app)
		tmpl.Execute(w, app)
	} else {
		tmpl := template.Must(template.ParseFiles("assets/index.html"))
		tmpl.Execute(w, nil)
	}
}

func getDob(str string) time.Time {
	if str == "" {
		return time.Now()
	}
	dobstr := strings.Split(str, "-")
	yob, _ := strconv.Atoi(dobstr[0]) //YYYY-MM-DD
	mob, _ := strconv.Atoi(dobstr[1])
	dob, _ := strconv.Atoi(dobstr[2])
	dobr := time.Date(yob, time.Month(mob), dob, 0, 0, 0, 0, time.UTC)
	return dobr
}
func createNewCustomer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("Form parsed")
	dob := getDob(r.Form["dob"][0])
	var err error
	if r.Form["userid"] != nil {
		_, err = Update(r.Form["userid"][0], r.Form["fname"][0], r.Form["lname"][0], dob, r.Form["gender"][0], r.Form["mail"][0], r.Form["address"][0])
	} else {
		_, err = Create(r.Form["fname"][0], r.Form["lname"][0], dob, r.Form["gender"][0], r.Form["mail"][0], r.Form["address"][0])
	}
	if err != nil {
		fmt.Println("Error Occurred in form parsing")
		tmpl := template.Must(template.ParseFiles("assets/Error.html"))
		tmpl.Execute(w, err.Error())
	} else {
		fmt.Println("Form parsed successfully")
		tmpl := template.Must(template.ParseFiles("assets/index.html"))
		tmpl.Execute(w, nil)
	}
}

func updateAction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	tmpl := template.Must(template.ParseFiles("assets/Form.html"))
	//tmpl = template.Must()
	tmpl.Execute(w, r.Form["userid"][0])
}

func search(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	app, e := searchData(r.Form["fname"][0], r.Form["lname"][0])
	if e != nil {
		fmt.Println("Error Occurred in form parsing")
		tmpl := template.Must(template.ParseFiles("assets/Error.html"))
		tmpl.Execute(w, e)
	}
	tmpl := template.Must(template.ParseFiles("assets/showAll.html"))
	fmt.Println(app)
	tmpl.Execute(w, app)
}

func main() {
	fmt.Println("Application started...")
	http.HandleFunc("/", choice)
	http.HandleFunc("/select", getChoice)
	http.HandleFunc("/submit", createNewCustomer)
	http.HandleFunc("/retry", choice)
	http.HandleFunc("/update", updateAction)
	http.HandleFunc("/search", search)
	http.ListenAndServe(":8080", nil)
}
