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

func GetAllRoles(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)
	data, err := app.RoleApplication{}.GetAllRoles()
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

func GetAllActiveRoles(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)
	data, err := app.RoleApplication{}.GetAllActiveRoles()
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

func GetRoleById(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	data, err := app.RoleApplication{}.GetRoleById(id)
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

func CreateRole(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)

	var roleDto = repo.RoleDTO{}
	bodyText, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bodyText, &roleDto)

	result, err := app.RoleApplication{}.CreateRole(roleDto)
	if err != nil {
		util.HttpNotFound(w)
		return
	}

	jsonResponse := make(map[string]interface{})
	jsonResponse["Status Code"] = 200
	jsonResponse["Status"] = "Success"
	jsonResponse["Role"] = result.Name + " is created succesfully."
	jsonResponse["Count"] = 1
	jsonResponse["Data"] = result

	response, _ := json.Marshal(&jsonResponse)
	w.Write(response)
}

func UpdateRole(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)

	var roleDto = repo.RoleDTO{}
	bodyText, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bodyText, &roleDto)

	result, err := app.RoleApplication{}.UpdateRole(roleDto)
	if err != nil {
		util.HttpNotFound(w)
		return
	}

	jsonResponse := make(map[string]interface{})
	jsonResponse["Status Code"] = 204
	jsonResponse["Status"] = "Success"
	jsonResponse["Role"] = result.Name + " is updated succesfully."
	jsonResponse["Count"] = 1
	jsonResponse["Data"] = result

	response, _ := json.Marshal(&jsonResponse)
	w.Write(response)
}

func DeleteRole(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	result, err := app.RoleApplication{}.DeleteRole(id)
	if err != nil {
		util.HttpNotFound(w)
		return
	}

	jsonResponse := make(map[string]interface{})
	jsonResponse["Status Code"] = 204
	jsonResponse["Status"] = "Success"
	jsonResponse["Role"] = result.Name + " is deleted succesfully."
	jsonResponse["Count"] = 1
	jsonResponse["Data"] = result

	response, _ := json.Marshal(&jsonResponse)
	w.Write(response)
}
