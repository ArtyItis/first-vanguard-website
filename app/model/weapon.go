package model

type Weapon struct {
	Id   string `json:"_id"`
	Rev  string `json:"_rev"`
	Name string `json:"name"`
	Type string `json:"type"`
}
