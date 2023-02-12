package homecontroller

import (
	"html/template"
	"log"
	"net/http"
	"partyinvites/constants"
	"partyinvites/data"
)

var HomeTemplates = make(map[string]*template.Template, 3)

func Index(writer http.ResponseWriter, request *http.Request) {
	err := HomeTemplates[constants.HomeWelcomePage].Execute(writer, nil)

	if err != nil {
		log.Printf("%v", err)
	}
}

func ListGuests(writer http.ResponseWriter, request *http.Request) {
	err := HomeTemplates[constants.HomeListPage].Execute(writer, data.Responses)

	if err != nil {
		log.Printf("%v", err)
	}
}
