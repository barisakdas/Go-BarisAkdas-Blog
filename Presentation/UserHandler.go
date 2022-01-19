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

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)
	data, err := app.UserApplication{}.GetAllUsers()
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

func GetAllActiveUsers(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)
	data, err := app.UserApplication{}.GetAllActiveUsers()
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

func GetUserById(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	data, err := app.UserApplication{}.GetUserById(id)
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

func CreateUser(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)

	var userDto = repo.UserDTO{}
	bodyText, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bodyText, &userDto)

	result, err := app.UserApplication{}.CreateUser(userDto)
	if err != nil {
		util.HttpNotFound(w)
		return
	}

	jsonResponse := make(map[string]interface{})
	jsonResponse["Status Code"] = 200
	jsonResponse["Status"] = "Success"
	jsonResponse["User"] = result.UserName + " is created succesfully."
	jsonResponse["Count"] = 1
	jsonResponse["Data"] = result

	response, _ := json.Marshal(&jsonResponse)
	w.Write(response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)

	var userDto = repo.UserDTO{}
	bodyText, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bodyText, &userDto)

	result, err := app.UserApplication{}.UpdateUser(userDto)
	if err != nil {
		util.HttpNotFound(w)
		return
	}

	jsonResponse := make(map[string]interface{})
	jsonResponse["Status Code"] = 204
	jsonResponse["Status"] = "Success"
	jsonResponse["User"] = result.UserName + " is updated succesfully."
	jsonResponse["Count"] = 1
	jsonResponse["Data"] = result

	response, _ := json.Marshal(&jsonResponse)
	w.Write(response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	result, err := app.UserApplication{}.DeleteUser(id)
	if err != nil {
		util.HttpNotFound(w)
		return
	}

	jsonResponse := make(map[string]interface{})
	jsonResponse["Status Code"] = 204
	jsonResponse["Status"] = "Success"
	jsonResponse["User"] = result.UserName + " is deleted succesfully."
	jsonResponse["Count"] = 1
	jsonResponse["Data"] = result

	response, _ := json.Marshal(&jsonResponse)
	w.Write(response)
}
