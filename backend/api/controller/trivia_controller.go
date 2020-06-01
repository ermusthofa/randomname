package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/ermusthofa/randomname/backend/api/model"
	"github.com/ermusthofa/randomname/backend/api/response"
)

func GetHealthz(w http.ResponseWriter, r *http.Request) {

	backendStatus := struct {
		Status string `json:"status"`
	}{
		Status: "healthy",
	}

	response.JSON(w, http.StatusOK, backendStatus)

}

func GetTrivia(w http.ResponseWriter, r *http.Request) {

	var trivia *model.Trivia

	x, err := fetchTrivia()
	if err != nil {
		response.ERROR(w, http.StatusGatewayTimeout, err)
	}

	y := "{\"trivia\":\"" + string(x) + "\"}"
	z := []byte(y)

	err = json.Unmarshal(z, &trivia)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	response.JSON(w, http.StatusOK, trivia)

}

func fetchTrivia() ([]byte, error) {

	_, month, date := time.Now().Date()

	url := "http://numbersapi.com/" + strconv.Itoa(int(month)) + "/" + strconv.Itoa(date) + "/date"

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
