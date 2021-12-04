package model

import (
	"encoding/json"
	"log"
	"time"

	couchdb "github.com/leesper/couchdb-golang"
)

type User struct {
	Id               string    `json:"_id"`
	Rev              string    `json:"_rev"`
	Discord_tag      string    `json:"discord_tag"`
	Name             string    `json:"name"`
	Password         string    `json:"password"`
	Password_tmp     string    `json:"password_tmp"`
	Company          string    `json:"company"`
	Permission_level int       `json:"permission_level"`
	Character        Character `json:"character"`
	Date             Date      `json:"date"`
	Taxes            Taxes     `json:"taxes"`
}

type Taxes struct {
	Previous_week    Tax `json:"previous_week"`
	Current_week     Tax `json:"current_week"`
	Next_week        Tax `json:"next_week"`
	Second_next_week Tax `json:"second_next_week"`
}

type Tax struct {
	Week   int  `json:"week"`
	Amount int  `json:"amount"`
	Payed  bool `json:"payed"`
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
		log.Println("added User: " + u.Name)
	}
	return err
}

func UpdateUser(user User) {
	r := User2Map(user)
	userDB.Set(user.Id, r)
	log.Println("updated " + user.Character.Name)
}

func DeleteUser(user User) {
	err := userDB.Delete(user.Id)
	if err != nil {
		log.Println("Error in DeleteUser", err)
	}
}

func GetUserById(id string) (User, error) {
	user, err := userDB.Get(id, nil)
	if err != nil {
		return User{}, err
	} else {
		return Map2User(user), nil
	}
}

func GetUserByName(name string) (user User, err error) {
	users, err := userDB.QueryJSON(`{"selector": {"name": {"$eq": "` + name + `"}}}`)
	if err == nil && len(users) > 0 {
		user := Map2User(users[0])
		return user, nil
	} else {
		return User{}, err
	}
}

func GetAllUsers() ([]map[string]interface{}, error) {
	users, err := userDB.QueryJSON(`{"selector": {"name": {"$ne": "admin"}}}`)
	if err != nil {
		return nil, err
	} else {
		return users, nil
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

func ShiftTaxes(user User) {
	user.Taxes.Previous_week = user.Taxes.Current_week
	user.Taxes.Current_week = user.Taxes.Next_week
	user.Taxes.Next_week = user.Taxes.Second_next_week
	_, week := time.Now().ISOWeek()
	tax := Tax{
		Week: week + 2,
	}
	user.Taxes.Second_next_week = tax
}
