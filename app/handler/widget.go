package handler

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/Leomn138/Widget-Factory/app/repository"
	"github.com/Leomn138/Widget-Factory/config"
	"github.com/gorilla/mux"
	"math"
	"strconv"
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
	if (!EnsureWidgetType("price", widgetMap["price"]) ||
		!EnsureWidgetType("melts", widgetMap["melts"]) ||
		!EnsureWidgetType("color", widgetMap["color"]) ||
		!EnsureWidgetType("name", widgetMap["name"]) ||
		!EnsureWidgetType("id", widgetMap["id"]) ||
		!EnsureWidgetType("inventory", widgetMap["inventory"]) || err != nil) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	response := repository.CreateDocument(dbConfig, strconv.FormatInt(int64(widgetMap["id"].(float64)), 10), widgetMap)
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
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	for index, element := range widgetMap {
		if (!EnsureWidgetType(index, element)) {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
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

func EnsureWidgetType(key string, value interface {}) bool {
	if (key == "price") {
		return EnsurePriceType(value)
	}
	if (key == "melts") {
		return EnsureMeltsType(value)
	}
	if (key == "color") {
		return EnsureColorType(value)
	}
	if (key == "inventory") {
		return EnsureInventoryType(value)
	}
	if (key == "name") {
		return EnsureNameType(value)
	}
	if (key == "id") {
		return EnsureIdType(value)
	}
	return false
}

func EnsurePriceType(value interface {}) bool {
	_, isPriceTypeCorrect := value.(float64)
	return isPriceTypeCorrect
}

func EnsureMeltsType(value interface {}) bool {
	_, isMeltsTypeCorrect := value.(bool)
	return isMeltsTypeCorrect
}

func EnsureColorType(value interface {}) bool {
	_, isColorTypeCorrect := value.(string)
	return isColorTypeCorrect
}
func EnsureNameType(value interface {}) bool {
	_, isNameTypeCorrect := value.(string)
	return isNameTypeCorrect
}
func EnsureIdType(value interface {}) bool {
	id, isIdTypeCorrect := value.(float64)
	isIdTypeCorrect = isIdTypeCorrect && id == math.Trunc(id)
	return isIdTypeCorrect
}
func EnsureInventoryType(value interface {}) bool {
	inventory, isInventoryTypeCorrect := value.(float64)
	isInventoryTypeCorrect = isInventoryTypeCorrect && inventory == math.Trunc(inventory)
	return isInventoryTypeCorrect
}