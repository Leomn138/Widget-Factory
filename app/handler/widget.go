package handler

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"widgetFactory/app/repository"
	"widgetFactory/config"
	"github.com/gorilla/mux"
)

func GetAllWidgets(dbConfig *config.DBConfig, w http.ResponseWriter, r *http.Request) {
	widgetList, response := repository.GetAllDocs(dbConfig)
	if response.Success == false {
		http.Error(w, http.StatusText(response.Code), response.Code)
		return
	}
	json.NewEncoder(w).Encode(widgetList)
}

func GetWidget(dbConfig *config.DBConfig, w http.ResponseWriter, r *http.Request) {
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


func CreateWidget(dbConfig *config.DBConfig, w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var widgetMap map[string] interface {}
	err = json.Unmarshal(requestBody, &widgetMap)
	// Todo Bater tipos
	if (widgetMap["id"] == "" || widgetMap["color"] == "" || widgetMap["name"] == "" || widgetMap["price"] == "" || widgetMap["inventory"] == "" || err != nil) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	response := repository.CreateDocument(dbConfig, widgetMap["id"].(string), widgetMap)
	if response.Success == false {
		http.Error(w, http.StatusText(response.Code), response.Code)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func UpdateWidget(dbConfig *config.DBConfig, w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var widgetMap map[string] interface {}
	err = json.Unmarshal(requestBody, &widgetMap)
	// Todo Bater tipos
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	id := params["id"]
	response := repository.UpdateDocument(dbConfig, id, widgetMap)
	if response.Success == false {
		http.Error(w, http.StatusText(response.Code), response.Code)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}