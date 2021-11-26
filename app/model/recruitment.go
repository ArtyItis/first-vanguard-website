package model

import (
	"encoding/json"
	"log"

	couchdb "github.com/leesper/couchdb-golang"
)

type Recruitment_Entry struct {
	Id        string    `json:"_id"`
	Rev       string    `json:"_rev"`
	Character Character `json:"character"`
	Status    string    `json:"status"`
	Date      Date      `json:"date"`
}

func init() {
	recruitmentDB, err = couchdb.NewDatabase(DBURL + "/recruitment")
	if err != nil {
		panic(err)
	}
}

func AddRecruitment_Entry(r Recruitment_Entry) error {
	recruitment_entry := Recruitment_entry2Map(r)
	delete(recruitment_entry, "_id")
	delete(recruitment_entry, "_rev")
	_, _, err := recruitmentDB.Save(recruitment_entry, nil)
	if err != nil {
		log.Println("Error in AddRecruitment_Entry: ", err)
	} else {
		log.Println("added Recruitment_Entry")
	}
	return err
}

func GetAllRecruitment_Entries() ([]map[string]interface{}, error) {
	recruitment_entries, err := recruitmentDB.QueryJSON(`{"selector": {"_id": {"$ne": ""}}}`)
	if err != nil {
		return nil, err
	} else {
		return recruitment_entries, nil
	}
}

func Recruitment_entry2Map(r Recruitment_Entry) (recruitment_entry map[string]interface{}) {
	jJSON, _ := json.Marshal(r)
	json.Unmarshal(jJSON, &recruitment_entry)
	return recruitment_entry
}

func Map2Recruitment_Entry(recruitment_entry map[string]interface{}) (r Recruitment_Entry) {
	jJSON, _ := json.Marshal(recruitment_entry)
	json.Unmarshal(jJSON, &r)
	return r
}
