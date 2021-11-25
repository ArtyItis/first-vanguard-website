package model

import (
	"time"

	couchdb "github.com/leesper/couchdb-golang"
)

var (
	err           error
	userDB        *couchdb.Database
	recruitmentDB *couchdb.Database
)

type Date struct {
	Date   time.Time `json:"Date"`
	Hour   int       `json:"hour"`
	Minute int       `json:"minute"`
	Second int       `json:"second"`
	Day    int       `json:"day"`
	Month  int       `json:"month"`
	Year   int       `json:"year"`
}
