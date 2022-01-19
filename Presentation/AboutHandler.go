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

func GetAllAbouts(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)
	data, err := app.AboutApplication{}.GetAllAbouts()
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

func GetAllActiveAbouts(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)
	data, err := app.AboutApplication{}.GetAllActiveAbouts()
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

func GetAboutById(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	data, err := app.AboutApplication{}.GetAboutById(id)
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

func CreateAbout(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)

	var aboutDto = repo.AboutDTO{}
	bodyText, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bodyText, &aboutDto)

	result, err := app.AboutApplication{}.CreateAbout(&aboutDto)
	if err != nil {
		util.HttpNotFound(w)
		return
	}

	jsonResponse := make(map[string]interface{})
	jsonResponse["Status Code"] = 200
	jsonResponse["Status"] = "Success"
	jsonResponse["Message"] = result.Title + " is created succesfully."
	jsonResponse["Count"] = 1
	jsonResponse["Data"] = result

	response, _ := json.Marshal(&jsonResponse)
	w.Write(response)
}

func UpdateAbout(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)

	var aboutDto = repo.AboutDTO{}
	bodyText, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bodyText, &aboutDto)

	result, err := app.AboutApplication{}.UpdateAbout(&aboutDto)
	if err != nil {
		util.HttpNotFound(w)
		return
	}

	jsonResponse := make(map[string]interface{})
	jsonResponse["Status Code"] = 204
	jsonResponse["Status"] = "Success"
	jsonResponse["Message"] = result.Title + " is updated succesfully."
	jsonResponse["Count"] = 1
	jsonResponse["Data"] = result

	response, _ := json.Marshal(&jsonResponse)
	w.Write(response)
}

func DeleteAbout(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	result, err := app.AboutApplication{}.DeleteAbout(id)
	if err != nil {
		util.HttpNotFound(w)
		return
	}

	jsonResponse := make(map[string]interface{})
	jsonResponse["Status Code"] = 204
	jsonResponse["Status"] = "Success"
	jsonResponse["Message"] = result.Title + " is deleted succesfully."
	jsonResponse["Count"] = 1
	jsonResponse["Data"] = result

	response, _ := json.Marshal(&jsonResponse)
	w.Write(response)
}
