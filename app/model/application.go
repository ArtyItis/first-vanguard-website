package model

import (
	"encoding/json"
	"log"

	couchdb "github.com/leesper/couchdb-golang"
)

type ApplicationEntry struct {
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

func AddApplicationEntry(r ApplicationEntry) error {
	applicationEntry := ApplicationEntry2Map(r)
	delete(applicationEntry, "_id")
	delete(applicationEntry, "_rev")
	_, _, err := applicationDB.Save(applicationEntry, nil)
	if err != nil {
		log.Println("Error in AddApplication_Entry: ", err)
	} else {
		log.Println("added Application_Entry")
	}
	return err
}

func UpdateApplicationEntry(applicationEntry ApplicationEntry) {
	r := ApplicationEntry2Map(applicationEntry)
	applicationDB.Set(applicationEntry.Id, r)
	log.Println("updated " + applicationEntry.Character.Name + "'s application entry")
}

func GetApplicationEntryById(id string) (ApplicationEntry, error) {
	applicationEntry, err := applicationDB.Get(id, nil)
	if err != nil {
		return ApplicationEntry{}, err
	} else {
		return Map2ApplicationEntry(applicationEntry), nil
	}
}

func GetAllApplicationEntries() ([]map[string]interface{}, error) {
	applicationEntries, err := applicationDB.QueryJSON(`{"selector": {"_id": {"$ne": ""}}}`)
	if err != nil {
		return nil, err
	} else {
		return applicationEntries, nil
	}
}

func GetAllApplicationEntriesOpen() ([]map[string]interface{}, error) {
	applicationEntries, err := applicationDB.QueryJSON(`{"selector": {"status": {"$ne": "angenommen"}}}`)
	if err != nil {
		return nil, err
	} else {
		return applicationEntries, nil
	}
}

func ApplicationEntry2Map(r ApplicationEntry) (applicationEntry map[string]interface{}) {
	jJSON, _ := json.Marshal(r)
	json.Unmarshal(jJSON, &applicationEntry)
	return applicationEntry
}

func Map2ApplicationEntry(applicationEntry map[string]interface{}) (r ApplicationEntry) {
	jJSON, _ := json.Marshal(applicationEntry)
	json.Unmarshal(jJSON, &r)
	return r
}
