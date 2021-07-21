package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/einnar82/net-http-api-sample/app/helpers"
)

func GetPosts(writer http.ResponseWriter, request *http.Request) {
	client := http.Client{Timeout: time.Duration(2) * time.Second}
	req, err := http.NewRequest("GET", "https://jsonplaceholders.typicode.com/posts", nil)
	if err != nil {
		writer.Header().Set("Content-type", "application/json")
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(helpers.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	req.Header.Add("Accept", `application/json`)
	response, err := client.Do(req)

	if err != nil {
		writer.Header().Set("Content-type", "application/json")
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(helpers.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		writer.Header().Set("Content-type", "application/json")
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(helpers.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(200)
	writer.Write(body)
}
