package handler

import (
	"net/http"
	"encoding/json"
	"widgetFactory/app/repository"
	"widgetFactory/config"
	"github.com/gorilla/mux"
)

func GetAllUsers(dbConfig *config.DBConfig, w http.ResponseWriter, r *http.Request) {
	userList, response := repository.GetAllDocs(dbConfig)
	if response.Success == false {
		http.Error(w, http.StatusText(response.Code), response.Code)
		return
	}
	json.NewEncoder(w).Encode(userList)
}

func GetUser(dbConfig *config.DBConfig, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	searchId := params["id"]

	deleteSettings := true
	widgetMap, response := repository.GetDocument(dbConfig, searchId, deleteSettings)
	if response.Success == false {
		http.Error(w, http.StatusText(response.Code), response.Code)
		return
	}
	json.NewEncoder(w).Encode(widgetMap)
}