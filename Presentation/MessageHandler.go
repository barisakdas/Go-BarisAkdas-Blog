package Presentation

import (
	app "Go-BarisAkdas-Blog/Application"
	repo "Go-BarisAkdas-Blog/Domain/Repository"
	util "Go-BarisAkdas-Blog/Infrastructure/Util"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetAllMessages(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)
	data, err := app.MessageApplication{}.GetAllMessages()
	if err != nil {
		util.HttpNotFound(w)
		return
	}

	jsonResponse := make(map[string]interface{})
	jsonResponse["Status Code"] = 200
	jsonResponse["Status"] = "Success"
	jsonResponse["Count"] = len(data)
	jsonResponse["Data"] = data

	response, _ := json.Marshal(&jsonResponse)
	w.Write(response)
}

func GetAllActiveMessages(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)
	data, err := app.MessageApplication{}.GetAllActiveMessages()
	if err != nil {
		util.HttpNotFound(w)
		return
	}

	jsonResponse := make(map[string]interface{})
	jsonResponse["Status Code"] = 200
	jsonResponse["Status"] = "Success"
	jsonResponse["Count"] = len(data)
	jsonResponse["Data"] = data

	response, _ := json.Marshal(&jsonResponse)
	w.Write(response)
}

func GetMessageById(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	data, err := app.MessageApplication{}.GetMessageById(id)
	if err != nil {
		util.HttpNotFound(w)
		return
	}

	jsonResponse := make(map[string]interface{})
	jsonResponse["Status Code"] = 200
	jsonResponse["Status"] = "Success"
	jsonResponse["Count"] = 1
	jsonResponse["Data"] = data

	response, _ := json.Marshal(&jsonResponse)
	w.Write(response)
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)

	var messageDto = repo.MessageDTO{}
	bodyText, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bodyText, &messageDto)

	result, err := app.MessageApplication{}.CreateMessage(messageDto)
	if err != nil {
		util.HttpNotFound(w)
		return
	}

	jsonResponse := make(map[string]interface{})
	jsonResponse["Status Code"] = 200
	jsonResponse["Status"] = "Success"
	jsonResponse["Message"] = result.Name + " is created succesfully."
	jsonResponse["Count"] = 1
	jsonResponse["Data"] = result

	response, _ := json.Marshal(&jsonResponse)
	w.Write(response)
}

func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)

	var messageDto = repo.MessageDTO{}
	bodyText, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bodyText, &messageDto)

	result, err := app.MessageApplication{}.UpdateMessage(messageDto)
	if err != nil {
		util.HttpNotFound(w)
		return
	}

	jsonResponse := make(map[string]interface{})
	jsonResponse["Status Code"] = 204
	jsonResponse["Status"] = "Success"
	jsonResponse["Message"] = result.Name + " is updated succesfully."
	jsonResponse["Count"] = 1
	jsonResponse["Data"] = result

	response, _ := json.Marshal(&jsonResponse)
	w.Write(response)
}

func ReplyMessage(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	answer := r.URL.Query().Get("answer")
	result, err := app.MessageApplication{}.ReplyMessage(id, answer)
	if err != nil {
		util.HttpNotFound(w)
		return
	}

	jsonResponse := make(map[string]interface{})
	jsonResponse["Status Code"] = 204
	jsonResponse["Status"] = "Success"
	jsonResponse["Message"] = result.Name + " is replied succesfully."
	jsonResponse["Count"] = 1
	jsonResponse["Data"] = result

	response, _ := json.Marshal(&jsonResponse)
	w.Write(response)
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	result, err := app.MessageApplication{}.DeleteMessage(id)
	if err != nil {
		util.HttpNotFound(w)
		return
	}

	jsonResponse := make(map[string]interface{})
	jsonResponse["Status Code"] = 204
	jsonResponse["Status"] = "Success"
	jsonResponse["Message"] = result.Name + " is deleted succesfully."
	jsonResponse["Count"] = 1
	jsonResponse["Data"] = result

	response, _ := json.Marshal(&jsonResponse)
	w.Write(response)
}
