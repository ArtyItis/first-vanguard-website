package model

import (
	"encoding/json"
	"log"

	couchdb "github.com/leesper/couchdb-golang"
)

type Role struct {
	Id               string       `json:"_id"`
	Rev              string       `json:"_rev"`
	Name             string       `json:"name"`
	Description      string       `json:"description"`
	Note             string       `json:"note"`
	Attributes       Attributes   `json:"attributes"`
	Weapon_main      RoleWeapon   `json:"weapon_main"`
	Weapon_secondary RoleWeapon   `json:"weapon_secondary"`
	Armor            Armor        `json:"armor"`
	Earring          JewelryPiece `json:"earring"`
	Amulet           JewelryPiece `json:"amulet"`
	Ring             JewelryPiece `json:"ring"`
}

type RoleWeapon struct {
	Name              string   `json:"name"`
	Perks             []string `json:"perks"`
	Gem               string   `json:"gem"`
	Skillpoints_image string   `json:"skillpoints_image"`
}

type Armor struct {
	Name   string   `json:"name"`
	Weight string   `json:"weight"`
	Gem    string   `json:"gem"`
	Perks  []string `json:"perks"`
}

type JewelryPiece struct {
	Name  string   `json:"name"`
	Gem   string   `json:"gem"`
	Perks []string `json:"perks"`
}

func init() {
	roleDB, err = couchdb.NewDatabase(DBURL + "/roles")
	if err != nil {
		panic(err)
	}
}

func AddRole(r Role) error {
	role := Role2Map(r)
	delete(role, "_id")
	delete(role, "_rev")
	_, _, err := roleDB.Save(role, nil)
	if err != nil {
		log.Println("Error in AddRole: ", err)
	} else {
		log.Println("added Role: " + r.Name)
	}
	return err
}

func UpdateRole(role Role) {
	r := Role2Map(role)
	roleDB.Set(role.Id, r)
	log.Println("updated Role: " + role.Name)
}

func DeleteRole(role Role) {
	err := roleDB.Delete(role.Id)
	if err != nil {
		log.Println("Error in DeleteRole", err)
	} else {
		log.Println("Role: " + role.Name + " has been deleted.")
	}
}

func GetRoleById(id string) (Role, error) {
	role, err := roleDB.Get(id, nil)
	if err != nil {
		return Role{}, err
	} else {
		return Map2Role(role), nil
	}
}
func GetRoleByName(name string) (role Role, err error) {
	roles, err := roleDB.QueryJSON(`{"selector": {"name": {"$eq": "` + name + `"}}}`)
	if err == nil && len(roles) > 0 {
		role := Map2Role(roles[0])
		return role, nil
	} else {
		return Role{}, err
	}
}

func GetAllRoles() ([]map[string]interface{}, error) {
	roles, err := roleDB.QueryJSON(`{"selector": {"_id": {"$ne": ""}}}`)
	if err != nil {
		return nil, err
	} else {
		return roles, nil
	}
}

func Role2Map(r Role) (role map[string]interface{}) {
	jJSON, _ := json.Marshal(r)
	json.Unmarshal(jJSON, &role)
	return role
}

func Map2Role(role map[string]interface{}) (r Role) {
	jJSON, _ := json.Marshal(role)
	json.Unmarshal(jJSON, &r)
	return r
}
