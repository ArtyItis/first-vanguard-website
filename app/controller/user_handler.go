package controller

import (
	"encoding/base64"
	"first-vanguard/app/model"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func RegisterGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/users/register.html", head, navigation, footer))
	data := Data{
		Session: GetSessionInformation(r),
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RegisterPOST(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("username")
	password := r.FormValue("password")
	//check if username is taken
	user_Exists, _ := model.GetUserByName(name)
	if user_Exists.Date != (model.Date{}) {
		//TODO username is taken
		http.Redirect(w, r, "/username_is_taken", http.StatusFound)
		return
	}
	//hash password
	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	b64HashedPwd := base64.StdEncoding.EncodeToString(hashedPwd)

	user := model.User{
		Name:             name,
		Password:         b64HashedPwd,
		Permission_level: 0,
		Date:             GetCurrentDate(),
	}
	err := model.AddUser(user)
	if err != nil {
		http.Redirect(w, r, "/bad_request", http.StatusFound)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	//validate user
	user, user_err := model.GetUserByName(username)
	//check if user exists
	if user_err != nil || (user.Password == "") {
		log.Println("invalid username")
		http.Redirect(w, r, GetPreviousRoute(r)+"?error=login&type=username", http.StatusFound)
		return
	}
	//validate password
	passwordDB, _ := base64.StdEncoding.DecodeString(user.Password)
	password_err := bcrypt.CompareHashAndPassword(passwordDB, []byte(password))
	if password_err != nil {
		log.Println("invalid password ", password_err)
		http.Redirect(w, r, GetPreviousRoute(r)+"?error=login&type=password", http.StatusFound)
		return
	}
	//save session
	session, _ := store.Get(r, "session")
	session.Values["authenticated"] = true
	session.Values["userId"] = user.Id
	session.Values["userCompany"] = user.Company
	session.Save(r, w)
	log.Println(username + " logged in")
	http.Redirect(w, r, GetPreviousRoute(r), http.StatusFound)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	userId := session.Values["userId"].(string)
	user, _ := model.GetUserById(userId)
	session.Values["authenticated"] = false
	session.Values["userId"] = ""
	session.Save(r, w)
	log.Println(user.Name + " logged out")
	http.Redirect(w, r, "/", http.StatusFound)
}

func UsersGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("users.html").Funcs(template.FuncMap{"add": Add}).
		Funcs(template.FuncMap{"convertPermissionLevel": ConvertPermissionLevelMap}).
		Funcs(template.FuncMap{"getRoleName": GetRoleName}).
		ParseFiles("template/users/users.html", head, navigation, footer))
	company := mux.Vars(r)["company"]
	users, _ := model.GetUsersByCompany(company)
	data := Data{
		Session: GetSessionInformation(r),
		Users:   users,
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RefiningGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("refining.html").Funcs(template.FuncMap{"add": Add}).
		Funcs(template.FuncMap{"convertPermissionLevel": ConvertPermissionLevelMap}).
		Funcs(template.FuncMap{"getRoleName": GetRoleName}).
		ParseFiles("template/refining.html", head, navigation, footer))
	users, _ := model.GetAllUsers()
	data := Data{
		Session: GetSessionInformation(r),
		Users:   users,
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CraftingGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("crafting.html").Funcs(template.FuncMap{"add": Add}).
		Funcs(template.FuncMap{"convertPermissionLevel": ConvertPermissionLevelMap}).
		Funcs(template.FuncMap{"getRoleName": GetRoleName}).
		ParseFiles("template/crafting.html", head, navigation, footer))
	users, _ := model.GetAllUsers()
	data := Data{
		Session: GetSessionInformation(r),
		Users:   users,
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UserGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("user.html").
		Funcs(template.FuncMap{"getWeaponName": GetWeaponName}).
		Funcs(template.FuncMap{"getWeaponByType": GetWeaponByType}).
		Funcs(template.FuncMap{"getRoleName": GetRoleName}).
		ParseFiles("template/users/user.html", attributes, jobs, roles, weapons, head, navigation, footer))
	weapons, _ := model.GetAllWeapons()
	roles, _ := model.GetAllRoles()
	user, _ := model.GetUserById(mux.Vars(r)["id"])
	data := Data{
		Session: GetSessionInformation(r),
		Weapons: weapons,
		Roles:   roles,
		User:    user,
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ChangePasswordGET(w http.ResponseWriter, r *http.Request) {
	ExecuteTemplate(w, r, "users/changePassword")
}

func ChangePasswordPOST(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	userId := session.Values["userId"].(string)
	userCompany := session.Values["userCompany"].(string)
	user, _ := model.GetUserById(userId)

	oldPassword := r.FormValue("oldPassword")
	newPassword := r.FormValue("newPassword")
	newPasswordrep := r.FormValue("newRepeatPassword")
	//check if old password is really the old password
	passwordDB, _ := base64.StdEncoding.DecodeString(user.Password)
	oldPasswordErr := bcrypt.CompareHashAndPassword(passwordDB, []byte(oldPassword))
	if oldPasswordErr != nil {
		log.Println(oldPasswordErr)
		http.Redirect(w, r, "/companies/"+userCompany+"/members/"+userId+"/changePassword?error=password&type=old", http.StatusFound)
		return
	}
	//compare newPassword and newPasswordRep
	if newPassword != newPasswordrep {
		log.Println("newPassword and newPasswordRep are not the same")
		http.Redirect(w, r, "/companies/"+userCompany+"/members/"+userId+"/changePassword?error=password&type=new", http.StatusFound)
		return
	}
	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(newPassword), 14)
	b64HashedPwd := base64.StdEncoding.EncodeToString(hashedPwd)
	user.Password_tmp = ""
	user.Password = b64HashedPwd
	model.UpdateUser(user)
	http.Redirect(w, r, "/companies/"+userCompany+"/members/"+userId+"/changePassword?success=passwordChange", http.StatusFound)
}

func TaxesGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("taxes.html").
		Funcs(template.FuncMap{"add": Add}).
		ParseFiles("template/taxes.html", head, navigation, footer))
	company := mux.Vars(r)["company"]
	users, _ := model.GetUsersByCompany(company)
	data := Data{
		Session: GetSessionInformation(r),
		Users:   users,
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func TaxesPOST(w http.ResponseWriter, r *http.Request) {
	company := mux.Vars(r)["company"]
	users, _ := model.GetAllUsers()
	for _, u := range users {
		user := model.Map2User(u)
		//values
		if pa := ParseInt(r.FormValue(user.Id + "-PA")); pa > 0 {
			user.Taxes.Previous_week.Amount = pa
		} else if pa < 0 {
			user.Taxes.Previous_week.Amount = 0
		}
		if ca := ParseInt(r.FormValue(user.Id + "-CA")); ca > 0 {
			user.Taxes.Current_week.Amount = ca
		} else if ca < 0 {
			user.Taxes.Current_week.Amount = 0
		}
		if na := ParseInt(r.FormValue(user.Id + "-NA")); na > 0 {
			user.Taxes.Next_week.Amount = na
		} else if na < 0 {
			user.Taxes.Next_week.Amount = 0
		}
		if sna := ParseInt(r.FormValue(user.Id + "-SNA")); sna > 0 {
			user.Taxes.Second_next_week.Amount = sna
		} else if sna < 0 {
			user.Taxes.Second_next_week.Amount = 0
		}
		//booleans
		if pp := r.FormValue(user.Id + "-PP"); pp == "on" {
			user.Taxes.Previous_week.Payed = true
		} else {
			user.Taxes.Previous_week.Payed = false
		}
		if cp := r.FormValue(user.Id + "-CP"); cp == "on" {
			user.Taxes.Current_week.Payed = true
		} else {
			user.Taxes.Current_week.Payed = false
		}
		if np := r.FormValue(user.Id + "-NP"); np == "on" {
			user.Taxes.Next_week.Payed = true
		} else {
			user.Taxes.Next_week.Payed = false
		}
		if snp := r.FormValue(user.Id + "-SNP"); snp == "on" {
			user.Taxes.Second_next_week.Payed = true
		} else {
			user.Taxes.Second_next_week.Payed = false
		}
		model.UpdateUser(user)
	}
	http.Redirect(w, r, "/companies/"+company+"/taxes", http.StatusFound)
}

func UserDeleteGET(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	company := mux.Vars(r)["company"]
	user, err := model.GetUserById(id)
	if err == nil {
		model.DeleteUser(user)
		http.Redirect(w, r, "/members", http.StatusFound)
	} else {
		http.Redirect(w, r, "/companies/"+company+"/members/"+id+"/?error=user&type=delete", http.StatusFound)
	}
}

func UserEditGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("userEdit.html").
		ParseFiles("template/users/userEdit.html", head, navigation, footer))
	weapons, _ := model.GetAllWeapons()
	roles, _ := model.GetAllRoles()
	user, _ := model.GetUserById(mux.Vars(r)["id"])
	data := Data{
		Session: GetSessionInformation(r),
		Weapons: weapons,
		Roles:   roles,
		User:    user,
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UserEditPost(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	company := mux.Vars(r)["company"]
	user, _ := model.GetUserById(id)

	if equipment_weight := r.FormValue("armor-weight"); equipment_weight != "" {
		user.Character.Equipment_weight = equipment_weight
	}
	if gearscore := r.FormValue("gearscore"); gearscore != "" {
		user.Character.Gearscore = ParseInt(gearscore)
	}
	if expertise := r.FormValue("expertise"); expertise != "" {
		user.Character.Expertise = ParseInt(expertise)
	}
	if discord_tag := r.FormValue("discord-tag"); discord_tag != "" {
		user.Discord_tag = discord_tag
	}
	// if company := r.FormValue("company"); company != "" {
	// 	user.Company = company
	// }
	if permission_level := r.FormValue("permission-level"); permission_level != "" {
		user.Permission_level = ParseInt(permission_level)
	}

	user.Character.Roles = r.Form["roles"]

	user.Character.Weapons = r.Form["weapons"]

	if strength := r.FormValue("strength"); strength != "" {
		user.Character.Attributes.Strength = ParseInt(strength)
	}
	if dexterity := r.FormValue("dexterity"); dexterity != "" {
		user.Character.Attributes.Dexterity = ParseInt(dexterity)
	}
	if intelligence := r.FormValue("intelligence"); intelligence != "" {
		user.Character.Attributes.Intelligence = ParseInt(intelligence)
	}
	if focus := r.FormValue("focus"); focus != "" {
		user.Character.Attributes.Focus = ParseInt(focus)
	}
	if constitution := r.FormValue("constitution"); constitution != "" {
		user.Character.Attributes.Constitution = ParseInt(constitution)
	}

	craftingJobs_Form := r.Form["crafting-jobs"]
	craftingJobs := model.CraftingJobs{}
	for _, craftingJob := range craftingJobs_Form {
		switch craftingJob {
		case "armoring":
			craftingJobs.Armoring = true
		case "weaponsmithing":
			craftingJobs.Weaponsmithing = true
		case "engineering":
			craftingJobs.Engineering = true
		case "jewelcrafting":
			craftingJobs.Jewelcrafting = true
		case "arcana":
			craftingJobs.Arcana = true
		case "cooking":
			craftingJobs.Cooking = true
		case "furnishing":
			craftingJobs.Furnishing = true
		}
	}

	refiningJobs_Form := r.Form["refining-jobs"]
	refiningJobs := model.RefiningJobs{}
	for _, refiningJob := range refiningJobs_Form {
		switch refiningJob {
		case "smelting":
			refiningJobs.Smelting = true
		case "stonecutting":
			refiningJobs.Stonecutting = true
		case "tanning":
			refiningJobs.Tanning = true
		case "weaving":
			refiningJobs.Weaving = true
		case "woodworking":
			refiningJobs.Woodworking = true
		}
	}

	user.Character.Crafting_jobs = craftingJobs
	user.Character.Refining_jobs = refiningJobs

	model.UpdateUser(user)
	http.Redirect(w, r, "/companies/"+company+"/members/"+id, http.StatusFound)
}
