package model

import (
	"encoding/json"
	"log"

	couchdb "github.com/leesper/couchdb-golang"
)

type User struct {
	Id               string    `json:"_id"`
	Rev              string    `json:"_rev"`
	Username         string    `json:"username"`
	Password         string    `json:"password"`
	Permission_Level int       `json:"permission_level"`
	Character        Character `json:"character"`
	Date             Date      `json:"date"`
}

func init() {
	userDB, err = couchdb.NewDatabase(DBURL + "/users")
	if err != nil {
		panic(err)
	}
}

func AddUser(u User) error {
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

func GetUserByName(name string) (user User, err error) {
	users, err := userDB.QueryJSON(`{"selector": {"username": {"$eq": "` + name + `"}}}`)
	if err == nil && len(users) > 0 {
		user := Map2User(users[0])
		return user, nil
	} else {
		return User{}, err
	}
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
