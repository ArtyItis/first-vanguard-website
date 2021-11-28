package model

import (
	couchdb "github.com/leesper/couchdb-golang"
)

var (
	err           error
	userDB        *couchdb.Database
	recruitmentDB *couchdb.Database
)

type Date struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
	Second int `json:"second"`
	Day    int `json:"day"`
	Month  int `json:"month"`
	Year   int `json:"year"`
}
