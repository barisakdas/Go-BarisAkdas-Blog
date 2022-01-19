package Presentation

import (
	"io/ioutil"
	"net/http"
)

func GetAllEndpoints(w http.ResponseWriter, r *http.Request) {
	file, _ := ioutil.ReadFile("endpoints.json")
	w.Write(file)
}
