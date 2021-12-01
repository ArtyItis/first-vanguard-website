package model

import (
	"encoding/json"
	"log"

	couchdb "github.com/leesper/couchdb-golang"
)

type Weapon struct {
	Id   string `json:"_id"`
	Rev  string `json:"_rev"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func init() {
	weaponDB, err = couchdb.NewDatabase(DBURL + "/weapons")
	if err != nil {
		panic(err)
	}
}

func AddWeapon(w Weapon) error {
	weapon := Weapon2Map(w)
	delete(weapon, "_id")
	delete(weapon, "_rev")
	_, _, err := weaponDB.Save(weapon, nil)
	if err != nil {
		log.Println("Error in AddWeapon: ", err)
	} else {
		log.Println("added Weapon")
	}
	return err
}

func UpdateWeapon(weapon Weapon) {
	r := Weapon2Map(weapon)
	weaponDB.Set(weapon.Id, r)
	log.Println("updated Weapon: " + weapon.Name)
}

func GetWeaponById(id string) (Weapon, error) {
	weapon, err := weaponDB.Get(id, nil)
	if err != nil {
		return Weapon{}, err
	} else {
		return Map2Weapon(weapon), nil
	}
}
func GetWeaponByName(name string) (weapon Weapon, err error) {
	weapons, err := userDB.QueryJSON(`{"selector": {"name": {"$eq": "` + name + `"}}}`)
	if err == nil && len(weapons) > 0 {
		weapon := Map2Weapon(weapons[0])
		return weapon, nil
	} else {
		return Weapon{}, err
	}
}

func GetAllWeapons() ([]map[string]interface{}, error) {
	weapons, err := weaponDB.QueryJSON(`{"selector": {"_id": {"$ne": ""}}}`)
	if err != nil {
		return nil, err
	} else {
		return weapons, nil
	}
}

func Weapon2Map(w Weapon) (weapon map[string]interface{}) {
	jJSON, _ := json.Marshal(w)
	json.Unmarshal(jJSON, &weapon)
	return weapon
}

func Map2Weapon(weapon map[string]interface{}) (w Weapon) {
	jJSON, _ := json.Marshal(weapon)
	json.Unmarshal(jJSON, &w)
	return w
}
