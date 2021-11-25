package model

import (
	"encoding/json"
	"log"

	couchdb "github.com/leesper/couchdb-golang"
)

type User struct {
	Id        string    `json:"_id"`
	Rev       string    `json:"_rev"`
	Password  string    `json:"password"`
	Character Character `json:"character"`
	Date      Date      `json:"date"`
}

func init() {
	userDB, err = couchdb.NewDatabase(DBURL + "/users")
	if err != nil {
		panic(err)
	}
}

func Add(u User) error {
	user := User2Map(u)
	delete(user, "_id")
	delete(user, "_rev")
	_, _, err = userDB.Save(user, nil)
	if err != nil {
		log.Println("Error in AddUser: ", err)
	} else {
		log.Println("added User")
	}
	return err
}

func User2Map(u User) (user map[string]interface{}) {
	jJSON, _ := json.Marshal(u)
	json.Unmarshal(jJSON, &user)
	return user
}

func Map2User(user map[string]interface{}) (u User) {
	jJSON, _ := json.Marshal(user)
	json.Unmarshal(jJSON, &u)
	return u
}
