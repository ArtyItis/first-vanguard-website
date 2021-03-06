package model

type Character struct {
	Name             string       `json:"name"`
	Gearscore        int          `json:"gearscore"`
	Expertise        int          `json:"expertise"`
	Equipment_weight string       `json:"equipment_weight"`
	Attributes       Attributes   `json:"attributes"`
	Roles            []string     `json:"roles"`
	Weapons          []string     `json:"weapons"`
	Old_company      string       `json:"old_company"`
	Crafting_jobs    CraftingJobs `json:"crafting_jobs"`
	Refining_jobs    RefiningJobs `json:"refining_jobs"`
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
	Smelting     bool `json:"smelting"`
	Stonecutting bool `json:"stonecutting"`
	Tanning      bool `json:"tanning"`
	Weaving      bool `json:"weaving"`
	Woodworking  bool `json:"woodworking"`
}

func (character Character) ContainsWeaponOfType(weaponType string) bool {
	for _, weaponID := range character.Weapons {
		weapon, _ := GetWeaponById(weaponID)
		if weapon.Type == weaponType {
			return true
		}
	}
	return false
}
