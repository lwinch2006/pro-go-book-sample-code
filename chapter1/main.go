package main

import (
	"fmt"
	"net/http"
	"partyinvites/constants"
	"partyinvites/controllers/homecontroller"
	"partyinvites/utils"
)

func main() {
	var pages = constants.GetHomePages()

	utils.LoadTemplates(&pages, &homecontroller.HomeTemplates)

	http.HandleFunc("/", homecontroller.Index)
	http.HandleFunc("/listguests", homecontroller.ListGuests)

	err := http.ListenAndServe(":5555", nil)

	if err != nil {
		fmt.Println(err)
	}
}
