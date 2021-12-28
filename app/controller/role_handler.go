package controller

import (
	"first-vanguard/app/model"
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func RolesGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("roles.html").
		Funcs(template.FuncMap{"modulo": Modulo}).
		ParseFiles("template/roles/roles.html", head, navigation, footer))
	roles, _ := model.GetAllRoles()
	data := Data{
		Session: GetSessionInformation(r),
		Roles:   roles,
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RoleGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/roles/role.html", "template/roles/assets/gemPerk.html", head, navigation, footer))
	role, _ := model.GetRoleById(mux.Vars(r)["id"])
	data := Data{
		Session: GetSessionInformation(r),
		Role:    role,
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RoleEditGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/roles/roleEdit.html", "template/roles/assets/gemPerkEdit.html", head, navigation, footer))
	role, _ := model.GetRoleById(mux.Vars(r)["id"])
	data := Data{
		Session: GetSessionInformation(r),
		Role:    role,
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RoleEditPOST(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	role, _ := model.GetRoleById(id)

	if name := r.FormValue("name"); name != "" {
		role.Name = name
	}
	if description := r.FormValue("description"); description != "" {
		role.Description = description
	}
	if note := r.FormValue("note"); note != "" {
		role.Note = note
	}

	// Attributes
	if strength := r.FormValue("strength"); strength != "" {
		role.Attributes.Strength = ParseInt(strength)
	}
	if dexterity := r.FormValue("dexterity"); dexterity != "" {
		role.Attributes.Dexterity = ParseInt(dexterity)
	}
	if intelligence := r.FormValue("intelligence"); intelligence != "" {
		role.Attributes.Intelligence = ParseInt(intelligence)
	}
	if focus := r.FormValue("focus"); focus != "" {
		role.Attributes.Focus = ParseInt(focus)
	}
	if constitution := r.FormValue("constitution"); constitution != "" {
		role.Attributes.Constitution = ParseInt(constitution)
	}
	// Weapon 1
	if weapon1_name := r.FormValue("weapon-1-name"); weapon1_name != "" {
		role.Weapon_main.Name = weapon1_name
	}
	if weapon1_perks := r.FormValue("weapon-1-perks"); weapon1_perks != "" {
		role.Weapon_main.Perks = strings.Split(weapon1_perks, ", ")
	}
	if weapon1_gem := r.FormValue("weapon-1-gem"); weapon1_gem != "" {
		role.Weapon_main.Gem = weapon1_gem
	}
	weapon1_file := UploadFile(r, "weapon-1-file", "roles", 10)
	if weapon1_file != "empty" {
		role.Weapon_main.Skillpoints_image = weapon1_file
	}
	// Weapon 2
	if weapon2_name := r.FormValue("weapon-1-name"); weapon2_name != "" {
		role.Weapon_secondary.Name = weapon2_name
	}
	if weapon2_perks := r.FormValue("weapon-2-perks"); weapon2_perks != "" {
		role.Weapon_secondary.Perks = strings.Split(weapon2_perks, ", ")
	}
	if weapon2_gem := r.FormValue("weapon-2-gem"); weapon2_gem != "" {
		role.Weapon_secondary.Gem = weapon2_gem
	}
	weapon2_file := UploadFile(r, "weapon-2-file", "roles", 10)
	if weapon2_file != "empty" {
		role.Weapon_secondary.Skillpoints_image = weapon2_file
	}
	// Armor & Jewelry
	if armor_weight := r.FormValue("armor-weight"); armor_weight != "" {
		role.Armor.Weight = armor_weight
	}
	if armor_perks := r.FormValue("armor-perks"); armor_perks != "" {
		role.Armor.Perks = strings.Split(armor_perks, ", ")
	}
	if armor_gem := r.FormValue("armor-gem"); armor_gem != "" {
		role.Armor.Gem = armor_gem
	}

	if amulet_perks := r.FormValue("amulet-perks"); amulet_perks != "" {
		role.Amulet.Perks = strings.Split(amulet_perks, ", ")
	}
	if amulet_gem := r.FormValue("amulet-gem"); amulet_gem != "" {
		role.Amulet.Gem = amulet_gem
	}

	if ring_perks := r.FormValue("ring-perks"); ring_perks != "" {
		role.Ring.Perks = strings.Split(ring_perks, ", ")
	}
	if ring_gem := r.FormValue("ring-gem"); ring_gem != "" {
		role.Ring.Gem = ring_gem
	}

	if earring_perks := r.FormValue("earring-perks"); earring_perks != "" {
		role.Earring.Perks = strings.Split(earring_perks, ", ")
	}
	if earring_gem := r.FormValue("earring-gem"); earring_gem != "" {
		role.Earring.Gem = earring_gem
	}
	model.UpdateRole(role, weapon1_file, weapon2_file)
	http.Redirect(w, r, "/roles/"+id, http.StatusFound)
}

func RoleFormGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/roles/roleForm.html", head, navigation, footer))
	data := Data{
		Session: GetSessionInformation(r),
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RoleFormPOST(w http.ResponseWriter, r *http.Request) {

	attributes := model.Attributes{
		Strength:     ParseInt(r.FormValue("strength")),
		Dexterity:    ParseInt(r.FormValue("dexterity")),
		Intelligence: ParseInt(r.FormValue("intelligence")),
		Focus:        ParseInt(r.FormValue("focus")),
		Constitution: ParseInt(r.FormValue("constitution")),
	}

	weapon1 := model.RoleWeapon{
		Name:              r.FormValue("weapon-1-name"),
		Perks:             strings.Split(r.FormValue("weapon-1-perks"), ", "),
		Gem:               r.FormValue("weapon-1-gem"),
		Skillpoints_image: UploadFile(r, "weapon-1-file", "roles", 10),
	}

	weapon2 := model.RoleWeapon{
		Name:              r.FormValue("weapon-2-name"),
		Perks:             strings.Split(r.FormValue("weapon-2-perks"), ", "),
		Gem:               r.FormValue("weapon-2-gem"),
		Skillpoints_image: UploadFile(r, "weapon-2-file", "roles", 10),
	}

	armor := model.Armor{
		Name:   "armor",
		Weight: r.FormValue("armor-weight"),
		Perks:  strings.Split(r.FormValue("armor-perks"), ", "),
		Gem:    r.FormValue("armor-gem"),
	}

	earring := model.JewelryPiece{
		Name:  "earring",
		Perks: strings.Split(r.FormValue("earring-perks"), ", "),
		Gem:   r.FormValue("earring-gem"),
	}

	amulet := model.JewelryPiece{
		Name:  "amulet",
		Perks: strings.Split(r.FormValue("amulet-perks"), ", "),
		Gem:   r.FormValue("amulet-gem"),
	}

	ring := model.JewelryPiece{
		Name:  "ring",
		Perks: strings.Split(r.FormValue("ring-perks"), ", "),
		Gem:   r.FormValue("ring-gem"),
	}

	role := model.Role{
		Name:             r.FormValue("name"),
		Description:      r.FormValue("description"),
		Note:             r.FormValue("note"),
		Attributes:       attributes,
		Weapon_main:      weapon1,
		Weapon_secondary: weapon2,
		Armor:            armor,
		Earring:          earring,
		Amulet:           amulet,
		Ring:             ring,
	}

	if err := model.AddRole(role); err == nil {
		http.Redirect(w, r, "/roles?success=role", http.StatusFound)
	}
}
