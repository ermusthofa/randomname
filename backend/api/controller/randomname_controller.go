package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ermusthofa/randomname/backend/api/model"
	"github.com/ermusthofa/randomname/backend/api/response"
)

func GetRandomName(w http.ResponseWriter, r *http.Request) {

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

	return body, nil

}
