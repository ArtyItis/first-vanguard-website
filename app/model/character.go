package model

type Character struct {
	Name          string     `json:"name"`
	Gearscore     int        `json:"gearscore"`
	Attributes    Attributes `json:"attributes"`
	Roles         Roles      `json:"roles"`
	Weapons       Weapons    `json:"weapons"`
	Old_Companies string     `json:"old_companies"`
}

type Attributes struct {
	Strength     int `json:"strength"`
	Dexterity    int `json:"dexterity"`
	Intelligence int `json:"intelligence"`
	Focus        int `json:"focus"`
	Constitution int `json:"constitution"`
}

type Roles struct {
	DD    bool `json:"dd"`
	Tank  bool `json:"tank"`
	Heal  bool `json:"heal"`
	Siege bool `json:"siege"`
}

type Weapons struct {
	Sword         bool `json:"sword"`
	Shield        bool `json:"shield"`
	Rapier        bool `json:"rapier"`
	Hatchet       bool `json:"hatchet"`
	Spear         bool `json:"spear"`
	Great_Axe     bool `json:"great_axe"`
	Warhammer     bool `json:"warhammer"`
	Bow           bool `json:"bow"`
	Musket        bool `json:"musket"`
	Firestaff     bool `json:"firestaff"`
	Lifestaff     bool `json:"lifestaff"`
	Ice_Gauntlet  bool `json:"ice_gauntlet"`
	Void_Gauntlet bool `json:"void_gauntlet"`
}
