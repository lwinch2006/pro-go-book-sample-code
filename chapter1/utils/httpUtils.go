package utils

import (
	"fmt"
	"html/template"
	"partyinvites/constants"
)

func LoadTemplates(pagesNames *[]string, pages *map[string]*template.Template) {
	for index, name := range *pagesNames {
		pageTemplate, err := template.ParseFiles(constants.MainLayout+".html", name+".html")

		if err != nil {
			panic(err)
		}

		(*pages)[name] = pageTemplate
		fmt.Println("Loaded template", index, name)
	}
}
