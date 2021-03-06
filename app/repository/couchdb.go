package repository

import (
	"strconv"
	"net/http"
	"io/ioutil"
	"time"
	"encoding/json"
	"bytes"
	"github.com/Leomn138/Widget-Factory/app/common"
	"github.com/Leomn138/Widget-Factory/config"
)
const (
	allDocsSufix = "_all_docs?include_docs="
)

type CouchDbAllDocs struct {
	TotalRows int64 `json:"total_rows"`
	Offset int64 `json:"offset"`
	Rows []CouchDbAllDocsRow `json:"rows"`
}
type CouchDbAllDocsRow struct {
	Id string `json:"id"`
	Key string `json:"key"`
	Value interface {} `json:"value"`
	Doc map[string] interface {} `json:"doc"`
}

func GetAllDocs(dbConfig *config.DBConfig) ([] map[string] interface{}, common.HttpResponse){
	var docList [] map[string] interface{}
	includeDocs := true
	var allDocsUrl = BuildAllDocsUrl(dbConfig, includeDocs)
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := netClient.Get(allDocsUrl)
	if err != nil {
		return docList, common.GetInternalServerErrorResponse()
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var couchdbAllDocs CouchDbAllDocs
	json.Unmarshal(body, &couchdbAllDocs)

	docList = make([]map[string] interface {}, len(couchdbAllDocs.Rows))

	for i, couchdbDoc := range couchdbAllDocs.Rows {
		docList[i] = couchdbDoc.Doc
	}
	return docList, common.GetSuccessResponse()
}

func GetDocument(dbConfig *config.DBConfig, id string, deleteSettings bool) (map[string] interface{}, common.HttpResponse) {
	var documentMap map[string] interface{}

	url := BuildGetUrl(dbConfig, id)
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	response, err := netClient.Get(url)
	if err != nil {
		return documentMap, common.GetInternalServerErrorResponse()
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return documentMap, common.GetInternalServerErrorResponse()
	}

	json.Unmarshal(body, &documentMap)

	if (len(documentMap) == 0) {
		return documentMap, common.GetNotFoundResponse()
	}

	if deleteSettings == true {
		delete(documentMap, "_id")
		delete(documentMap, "_rev");
	}
	return documentMap, common.GetSuccessResponse()
}

func CreateDocument(dbConfig *config.DBConfig, id string, newDocumentMap map[string] interface{}) common.HttpResponse {
	url := BuildGetUrl(dbConfig, id)
	newDocumentMap["_id"] = id
	document, err := json.Marshal(newDocumentMap)
	if err != nil {
		return common.GetInternalServerErrorResponse()
	}
	return PutDocument(url, document, "")
}

func UpdateDocument(dbConfig *config.DBConfig, id string, newDocumentMap map[string] interface{}) common.HttpResponse {
	url := BuildGetUrl(dbConfig, id)
	deleteSettings := false
	oldDocumentMap, response := GetDocument(dbConfig, id, deleteSettings)
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

func CreateDatabaseIfNotExists(dbConfig *config.DBConfig){
	url := BuildUrl(dbConfig)
	PutDocument(url, []byte{}, "")
}

func PutDocument(url string, document []byte, revision string) common.HttpResponse {

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(document))
	if err != nil {
		return common.GetInternalServerErrorResponse()
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
		return common.GetInternalServerErrorResponse()
	}
	var documentMap map [string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return common.GetInternalServerErrorResponse()
	}

	json.Unmarshal(body, &documentMap)

	if documentMap["error"] != nil {
		if (documentMap["error"].(string) == "conflict"){
			return common.GetConflictResponse()
		}

		return common.GetInternalServerErrorResponse()
	}
	defer resp.Body.Close()
	return common.GetSuccessResponse()
}

func BuildAllDocsUrl(dbConfig *config.DBConfig, includeDocs bool) string {
	var allDocsUrl string
	allDocsUrl = BuildUrl(dbConfig) + "/" + allDocsSufix + strconv.FormatBool(includeDocs)
	return allDocsUrl
}

func BuildGetUrl(dbConfig *config.DBConfig, id string) string {
	var getDocUrl string
	getDocUrl = BuildUrl(dbConfig) + "/" + id
	return getDocUrl
}

func BuildUrl(dbConfig *config.DBConfig) string {
	var getDocUrl string
	getDocUrl = dbConfig.Protocol + "://" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + "/" + dbConfig.Name
	return getDocUrl
}