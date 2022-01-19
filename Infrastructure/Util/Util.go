package util

import (
	"encoding/json"
	"net/http"
)

func HttpNotFound(w http.ResponseWriter) {
	resp := make(map[string]interface{})
	resp["Message"] = "Resource not found."
	resp["Status Code"] = 404
	resp["Status"] = "Fail"
	respMessage, _ := json.Marshal(resp)
	w.Write(respMessage)
}
