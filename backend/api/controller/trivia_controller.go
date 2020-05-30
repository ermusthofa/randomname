package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ermusthofa/triviadate/backend/api/model"
	"github.com/ermusthofa/triviadate/backend/api/response"
)

func GetTrivia(w http.ResponseWriter, r *http.Request) {

	var randomUser *model.RandomUser

	x, err := fetchRandomUser()
	if err != nil {
		response.ERROR(w, http.StatusGatewayTimeout, err)
	}

	err = json.Unmarshal(x, &randomUser)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	response.JSON(w, http.StatusOK, randomUser)

}

func fetchRandomUser() ([]byte, error) {

	url := "https://api.randomuser.me"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Println(body)

	return body, nil

}
