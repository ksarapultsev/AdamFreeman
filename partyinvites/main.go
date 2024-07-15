package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

var responses = make([]*Rsvp, 0, 10)
var templetes = make(map[string]*template.Template, 3)

func loadTempLates() {
	templateNames := [6]string{"welcome", "form", "thanks", "sorry", "list", "modallist2"}
	for index, name := range templateNames {
		t, err := template.ParseFiles("layout.html", name+".html")
		if err == nil {
			templetes[name] = t
			fmt.Println("loaded templete", index, name)
		} else {
			panic(err)
		}
	}
}

func welcomeHandler(writer http.ResponseWriter, request *http.Request) {
	templetes["welcome"].Execute(writer, nil)
}
func modallistHandler(writer http.ResponseWriter, request *http.Request) {
	templetes["modallist2"].Execute(writer, nil)
}

func listHandler(writer http.ResponseWriter, request *http.Request) {
	templetes["list"].Execute(writer, responses)
}

type formData struct {
	*Rsvp
	Errors []string
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		templetes["form"].Execute(writer, formData{
			Rsvp: &Rsvp{}, Errors: []string{},
		})
	} else if request.Method == http.MethodPost {
		request.ParseForm()
		responseData := Rsvp{
			Name:       request.Form["name"][0],
			Email:      request.Form["email"][0],
			Phone:      request.Form["phone"][0],
			WillAttend: request.Form["willattend"][0] == "true",
		}

		errors := []string{}

		if responseData.Name == "" {
			errors = append(errors, "Please enter your name")
		}
		if responseData.Email == "" {
			errors = append(errors, "Please enter your email address")
		}
		if responseData.Phone == "" {
			errors = append(errors, "Please enter your phone number")
		}
		if len(errors) > 0 {
			templetes["form"].Execute(writer, formData{
				Rsvp: &responseData, Errors: errors,
			})
		} else {
			responses = append(responses, &responseData)

			if responseData.WillAttend {
				templetes["thanks"].Execute(writer, responseData.Name)
			} else {
				templetes["sorry"].Execute(writer, responseData.Name)
			}
		}

	}
}

func main() {
	loadTempLates()

	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/modallist2", modallistHandler)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
	}

}
