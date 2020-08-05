package issue

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func SetupRoutes(apiBasePath string) {
	handleIssue := http.HandlerFunc(getIssueHandler)
	handleInertIssue := http.HandlerFunc(insertIssueHandler)

	http.Handle(fmt.Sprintf("%s/", apiBasePath), handleIssue)
	http.Handle(fmt.Sprintf("%s", apiBasePath), handleInertIssue)
}

func getIssueHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "/")
	issueID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	issue, err := getIssue(issueID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if issue == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if r.Method == http.MethodGet {
		issueJSON, err := json.Marshal(issue)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(issueJSON)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func insertIssueHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var newIssue Issue
		err = json.Unmarshal(bodyBytes, &newIssue)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = insertIssue(newIssue)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
