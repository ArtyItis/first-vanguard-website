package model

type Character struct {
	Name             string       `json:"name"`
	Gearscore        int          `json:"gearscore"`
	Equipment_Weight int          `json:"equipment_Weight"`
	Attributes       Attributes   `json:"attributes"`
	Roles            []Role       `json:"roles"`
	Weapons          []Weapon     `json:"weapons"`
	Old_Company      string       `json:"old_company"`
	Discord_Tag      string       `json:"discord_tag"`
	CraftingJobs     CraftingJobs `json:"craftingjobs"`
	RefiningJobs     RefiningJobs `json:"refiningjobs"`
}

type Attributes struct {
	Strength     int `json:"strength"`
	Dexterity    int `json:"dexterity"`
	Intelligence int `json:"intelligence"`
	Focus        int `json:"focus"`
	Constitution int `json:"constitution"`
}

type CraftingJobs struct {
	Armoring       bool `json:"armoring"`
	Weaponsmithing bool `json:"weaponsmithing"`
	Engineering    bool `json:"engineering"`
	Jewelcrafting  bool `json:"jewelcrafting"`
	Arcana         bool `json:"arcana"`
	Cooking        bool `json:"cooking"`
	Furnishing     bool `json:"furnishing"`
}
type RefiningJobs struct {
	Smelting     bool `json:"semlting"`
	Stonecutting bool `json:"stonecutting"`
	Tanning      bool `json:"tanning"`
	Weaving      bool `json:"weaving"`
	Woodworking  bool `json:"woodworking"`
}
