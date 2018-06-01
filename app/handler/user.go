package handler

import (
	"net/http"
	"encoding/json"
	"github.com/Leomn138/Widget-Factory/app/repository"
	"github.com/Leomn138/Widget-Factory/config"
	"github.com/gorilla/mux"
)

func GetAllUsers(dbConfig *config.DBConfig, w http.ResponseWriter, r *http.Request) {
	userList, response := repository.GetAllDocs(dbConfig)
	if response.Success == false {
		w.WriteHeader(http.response.Code)
		w.Write([]byte(http.StatusText(response.Code)))
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
		w.WriteHeader(http.response.Code)
		w.Write([]byte(http.StatusText(response.Code)))
		return
	}
	json.NewEncoder(w).Encode(widgetMap)
}