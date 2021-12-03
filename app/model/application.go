package model

import (
	"encoding/json"
	"log"

	couchdb "github.com/leesper/couchdb-golang"
)

type Application struct {
	Id          string    `json:"_id"`
	Rev         string    `json:"_rev"`
	Status      string    `json:"status"`
	Discord_tag string    `json:"discord_tag"`
	Character   Character `json:"character"`
	Date        Date      `json:"date"`
}

func init() {
	applicationDB, err = couchdb.NewDatabase(DBURL + "/applications")
	if err != nil {
		panic(err)
	}
}

func AddApplication(r Application) error {
	application := Application2Map(r)
	delete(application, "_id")
	delete(application, "_rev")
	_, _, err := applicationDB.Save(application, nil)
	if err != nil {
		log.Println("Error in AddApplication_Entry: ", err)
	} else {
		log.Println("added Application_Entry")
	}
	return err
}

func UpdateApplication(application Application) {
	r := Application2Map(application)
	applicationDB.Set(application.Id, r)
	log.Println("updated " + application.Character.Name + "'s application entry")
}

func GetApplicationById(id string) (Application, error) {
	application, err := applicationDB.Get(id, nil)
	if err != nil {
		return Application{}, err
	} else {
		return Map2Application(application), nil
	}
}

func GetAllApplications() ([]map[string]interface{}, error) {
	applications, err := applicationDB.QueryJSON(`{"selector": {"_id": {"$ne": ""}}}`)
	if err != nil {
		return nil, err
	} else {
		return applications, nil
	}
}

func GetAllApplicationsOpen() ([]map[string]interface{}, error) {
	applications, err := applicationDB.QueryJSON(`{"selector": {"status": {"$ne": "angenommen"}}}`)
	if err != nil {
		return nil, err
	} else {
		return applications, nil
	}
}

func Application2Map(r Application) (application map[string]interface{}) {
	jJSON, _ := json.Marshal(r)
	json.Unmarshal(jJSON, &application)
	return application
}

func Map2Application(application map[string]interface{}) (r Application) {
	jJSON, _ := json.Marshal(application)
	json.Unmarshal(jJSON, &r)
	return r
}
