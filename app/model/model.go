package model

import (
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
