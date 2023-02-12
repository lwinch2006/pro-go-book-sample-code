package homecontroller

import (
	"html/template"
	"log"
	"net/http"
	"partyinvites/constants"
	"partyinvites/data/models"
	"partyinvites/data/viewmodels/homecontroller"
)

var HomeTemplates = make(map[string]*template.Template, 3)

func Index(writer http.ResponseWriter, _ *http.Request) {
	err := HomeTemplates[constants.HomeWelcomePage].Execute(writer, nil)

	if err != nil {
		log.Printf("%v", err)
	}
}

func ListGuests(writer http.ResponseWriter, _ *http.Request) {
	err := HomeTemplates[constants.HomeListPage].Execute(writer, models.Responses)

	if err != nil {
		log.Printf("%v", err)
	}
}

func ReplyHandler(writer http.ResponseWriter, request *http.Request) {
	var err error

	if request.Method == http.MethodGet {
		err = HomeTemplates[constants.HomeFormPage].Execute(writer, homecontroller.ReplyViewModel{Reply: &models.Rvsp{}, Errors: []string{}})
	} else if request.Method == http.MethodPost {
		err = request.ParseForm()
		responseData := models.Rvsp{
			Name:       request.Form["name"][0],
			Email:      request.Form["email"][0],
			Phone:      request.Form["phone"][0],
			WillAttend: request.Form["willattend"][0] == "true",
		}

		responseErrors := []string{}

		if responseData.Name == "" {
			responseErrors = append(responseErrors, "Please enter your name")
		}

		if responseData.Email == "" {
			responseErrors = append(responseErrors, "Please enter your email")
		}

		if responseData.Phone == "" {
			responseErrors = append(responseErrors, "Please enter your phone")
		}

		if len(responseErrors) > 0 {
			err = HomeTemplates[constants.HomeFormPage].Execute(writer, homecontroller.ReplyViewModel{Reply: &responseData, Errors: responseErrors})
		} else {
			models.Responses = append(models.Responses, &responseData)

			if responseData.WillAttend {
				err = HomeTemplates[constants.HomeThanksPage].Execute(writer, responseData.Name)
			} else {
				err = HomeTemplates[constants.HomeSorryPage].Execute(writer, responseData.Name)
			}
		}
	}

	if err != nil {
		log.Printf("%v", err)
	}
}
