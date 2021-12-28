package model

import (
	"log"
	"os"

	couchdb "github.com/leesper/couchdb-golang"
)

var (
	err           error
	userDB        *couchdb.Database
	applicationDB *couchdb.Database
	roleDB        *couchdb.Database
	weaponDB      *couchdb.Database
)

type Date struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
	Second int `json:"second"`
	Day    int `json:"day"`
	Month  int `json:"month"`
	Year   int `json:"year"`
	Week   int `json:"week"`
}

//classic contain method for arrays
//implemented for string arrays
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

/*
use this for file deletion. Look at task_handler.go EditTaskUpload() DeleteTask() for an example
param files should contain a string array which contains all currently used filepaths
param filePath should contain the filePath of the replaced file
*/
func CleanUpFiles(files []string, filePath string) {
	if filePath == "" {
		return
	}
	log.Println(filePath)
	if !contains(files, filePath) {
		filePath = filePath[1:]
		log.Println("no more references detected. deleting " + filePath)
		err := os.Remove(filePath) //deleting file
		if err != nil {
			log.Println(err)
		}
	}
}
