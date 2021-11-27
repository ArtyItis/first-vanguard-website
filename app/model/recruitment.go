package model

import (
	"encoding/json"
	"log"

	couchdb "github.com/leesper/couchdb-golang"
)

type RecruitmentEntry struct {
	Id        string    `json:"_id"`
	Rev       string    `json:"_rev"`
	Character Character `json:"character"`
	Status    string    `json:"status"`
	Date      Date      `json:"date"`
}

func init() {
	recruitmentDB, err = couchdb.NewDatabase(DBURL + "/recruitments")
	if err != nil {
		panic(err)
	}
}

func AddRecruitmentEntry(r RecruitmentEntry) error {
	recruitmentEntry := RecruitmentEntry2Map(r)
	delete(recruitmentEntry, "_id")
	delete(recruitmentEntry, "_rev")
	_, _, err := recruitmentDB.Save(recruitmentEntry, nil)
	if err != nil {
		log.Println("Error in AddRecruitment_Entry: ", err)
	} else {
		log.Println("added Recruitment_Entry")
	}
	return err
}

func UpdateRecruitmentEntry(recruitmentEntry RecruitmentEntry) {
	r := RecruitmentEntry2Map(recruitmentEntry)
	recruitmentDB.Set(recruitmentEntry.Id, r)
	log.Println("updated " + recruitmentEntry.Character.Name + "'s recruitment entry")
}

func GetRecruitmentEntryById(id string) (RecruitmentEntry, error) {
	recruitmentEntry, err := recruitmentDB.Get(id, nil)
	if err != nil {
		return RecruitmentEntry{}, err
	} else {
		return Map2Recruitment_Entry(recruitmentEntry), nil
	}
}

func GetAllRecruitmentEntries() ([]map[string]interface{}, error) {
	recruitmentEntries, err := recruitmentDB.QueryJSON(`{"selector": {"_id": {"$ne": ""}}}`)
	if err != nil {
		return nil, err
	} else {
		return recruitmentEntries, nil
	}
}

func RecruitmentEntry2Map(r RecruitmentEntry) (recruitmentEntry map[string]interface{}) {
	jJSON, _ := json.Marshal(r)
	json.Unmarshal(jJSON, &recruitmentEntry)
	return recruitmentEntry
}

func Map2Recruitment_Entry(recruitmentEntry map[string]interface{}) (r RecruitmentEntry) {
	jJSON, _ := json.Marshal(recruitmentEntry)
	json.Unmarshal(jJSON, &r)
	return r
}
