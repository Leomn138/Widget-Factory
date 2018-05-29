package service

import (
	"net/http"
	"encoding/json"
	"widgetFactory/repository"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	userList, response := repository.GetAllDocs(databaseName)
	if response.Success == false {
		http.Error(w, http.StatusText(response.Code), response.Code)
		return
	}
	json.NewEncoder(w).Encode(userList)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	searchId := params["id"]

	deleteSettings := true
	widgetMap, response := repository.GetDocument(databaseName, documentType, searchId, deleteSettings)
	if response.Success == false {
		http.Error(w, http.StatusText(response.Code), response.Code)
		return
	}
	json.NewEncoder(w).Encode(widgetMap)
}