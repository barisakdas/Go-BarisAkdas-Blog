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

func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)
	data, err := app.ArticleApplication{}.GetAllArticles()
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

func GetAllActiveArticles(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)
	data, err := app.ArticleApplication{}.GetAllActiveArticles()
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

func GetArticleById(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	data, err := app.ArticleApplication{}.GetArticleById(id)
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

func GetArticleBySlug(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)

	slug := r.URL.Query().Get("slug")
	data, err := app.ArticleApplication{}.GetArticleBySlug(slug)
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

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)

	var articleDto = repo.ArticleDTO{}
	bodyText, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bodyText, &articleDto)

	result, err := app.ArticleApplication{}.CreateArticle(&articleDto)
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

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)

	var articleDto = repo.ArticleDTO{}
	bodyText, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bodyText, &articleDto)

	result, err := app.ArticleApplication{}.UpdateArticle(&articleDto)
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

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	result, err := app.ArticleApplication{}.DeleteArticle(id)
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
