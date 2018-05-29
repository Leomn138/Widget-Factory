package repository

import (
	"strconv"
	"net/http"
	"io/ioutil"
	"time"
	"encoding/json"
	"bytes"
	"widgetFactory/utils"
)
const (
	couchdbHost = "http://127.0.0.1:5984"
	allDocsSufix = "/_all_docs?include_docs="
)

func GetAllDocs(databaseName string) ([] map[string] interface{}, rest.HttpResponse){
	var docList [] map[string] interface{}
	includeDocs := true
	var allDocsUrl = BuildAllDocsUrl(databaseName, includeDocs)
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := netClient.Get(allDocsUrl)
	if err != nil {
		return docList, rest.GetInternalServerErrorResponse()
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var couchdbAllDocs map[string] interface {}
	json.Unmarshal(body, &couchdbAllDocs)

	docList = make([] map[string] interface{}, len(couchdbAllDocs["rows"].([] map[string] interface{})))
	for i, couchdbDoc := range couchdbAllDocs["rows"].([] map[string] interface{}) {
		docList[i] = couchdbDoc["doc"].(map[string] interface{})
	}
	return docList, rest.GetSuccessResponse()
}

func GetDocument(databaseName string, documentType string, id string, deleteSettings bool) (map[string] interface{}, rest.HttpResponse) {
	var documentMap map[string] interface{}

	url := BuildGetUrl(databaseName, documentType, id)
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	response, err := netClient.Get(url)
	if err != nil {
		return documentMap, rest.GetInternalServerErrorResponse()
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return documentMap, rest.GetInternalServerErrorResponse()
	}

	json.Unmarshal(body, &documentMap)

	if (len(documentMap) == 0) {
		return documentMap, rest.GetNotFoundResponse()
	}

	if deleteSettings == true {
		delete(documentMap, "_id")
		delete(documentMap, "_rev");
	}
	return documentMap, rest.GetSuccessResponse()
}

func CreateDocument(databaseName string, documentType string, id string, newDocumentMap map[string] interface{}) rest.HttpResponse {
	url := BuildCreateUrl(databaseName)
	newDocumentMap["_id"] = documentType + "-" + id
	document, _ := json.Marshal(newDocumentMap)
	return PutDocument(url, document, "")
}

func UpdateDocument(databaseName string, documentType string, id string, newDocumentMap map[string] interface{}) rest.HttpResponse {
	url := BuildGetUrl(databaseName, documentType, id)
	deleteSettings := false
	oldDocumentMap, response := GetDocument(databaseName, documentType, id, deleteSettings)
	if response.Success == false {
		return response
	}

	for index, element := range newDocumentMap {
		oldDocumentMap[index] = element
	}
	newDocumentMap = oldDocumentMap

	document, _ := json.Marshal(newDocumentMap)
	revision := newDocumentMap["_rev"].(string)
	return PutDocument(url, document, revision)
}

func PutDocument(url string, document []byte, revision string) rest.HttpResponse {

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(document))
	if err != nil {
		return rest.GetInternalServerErrorResponse()
	}

	req.Header.Set("Content-Type", "application/json")
	if revision != "" {
		req.Header.Set("If-Match", revision)
	}
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := netClient.Do(req)
	if err != nil {
		return rest.GetInternalServerErrorResponse()
	}
	defer resp.Body.Close()
	return rest.GetSuccessResponse()
}

func BuildAllDocsUrl(databaseName string, includeDocs bool) string {
	var allDocsUrl string
	allDocsUrl = couchdbHost + "/" + databaseName + allDocsSufix + strconv.FormatBool(includeDocs)
	return allDocsUrl
}

func BuildGetUrl(databaseName string, documentType string, id string) string {
	var getDocUrl string
	getDocUrl = couchdbHost + "/" + databaseName + "/" + documentType + "-" + id
	return getDocUrl
}

func BuildCreateUrl(databaseName string) string {
	var getDocUrl string
	getDocUrl = couchdbHost + "/" + databaseName
	return getDocUrl
}