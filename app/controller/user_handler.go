package controller

import (
	"encoding/base64"
	"forgottennw/app/model"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func RegisterGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/register.html", head, navigation, footer))
	data := Data{
		Session: GetSessionInformation(r),
	}
	tmpl.Execute(w, data)
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
		http.Redirect(w, r, GetPreviousRoute(r)+"?loginError=username", http.StatusFound)
		return
	}
	//validate password
	passwordDB, _ := base64.StdEncoding.DecodeString(user.Password)
	password_err := bcrypt.CompareHashAndPassword(passwordDB, []byte(password))
	if password_err != nil {
		log.Println("invalid password ", password_err)
		http.Redirect(w, r, GetPreviousRoute(r)+"?loginError=password", http.StatusFound)
		return
	}
	//save session
	session, _ := store.Get(r, "session")
	session.Values["authenticated"] = true
	session.Values["userId"] = user.Id
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
	tmpl := template.Must(template.New("users.html").Funcs(template.FuncMap{"add": Add}).Funcs(template.FuncMap{"convertPermissionLevel": ConvertPermissionLevelMap}).ParseFiles("template/users.html", head, navigation, footer))
	users, _ := model.GetAllUsers()
	data := Data{
		Session: GetSessionInformation(r),
		Users:   users,
	}
	tmpl.Execute(w, data)
}

func UserGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("user.html").
		Funcs(template.FuncMap{"convertPermissionLevel": ConvertPermissionLevelUser}).
		Funcs(template.FuncMap{"containsWeapon": ContainsWeapon}).
		Funcs(template.FuncMap{"containsRole": ContainsRole}).
		ParseFiles("template/user.html", head, navigation, footer))
	weapons, _ := model.GetAllWeapons()
	roles, _ := model.GetAllRoles()
	user, _ := model.GetUserById(mux.Vars(r)["id"])
	data := Data{
		Session: GetSessionInformation(r),
		Weapons: weapons,
		Roles:   roles,
		User:    user,
	}
	tmpl.Execute(w, data)
}

func ChangePasswordGET(w http.ResponseWriter, r *http.Request) {
	ExecuteTemplate(w, r, "changePassword")
}

func ChangePasswordPOST(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	userId := session.Values["userId"].(string)
	user, _ := model.GetUserById(userId)

	oldPassword := r.FormValue("oldPassword")
	newPassword := r.FormValue("newPassword")
	newPasswordrep := r.FormValue("newRepeatPassword")
	//check if old password is really the old password
	passwordDB, _ := base64.StdEncoding.DecodeString(user.Password)
	oldPasswordErr := bcrypt.CompareHashAndPassword(passwordDB, []byte(oldPassword))
	if oldPasswordErr != nil {
		log.Println(oldPasswordErr)
		http.Redirect(w, r, "/members/"+userId+"/changePassword?changePasswordError=oldPassword", http.StatusFound)
		return
	}
	//compare newPassword and newPasswordRep
	if newPassword != newPasswordrep {
		log.Println("newPassword and newPasswordRep are not the same")
		http.Redirect(w, r, "/members/"+userId+"/changePassword?changePasswordError=newPassword", http.StatusFound)
		return
	}
	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(newPassword), 14)
	b64HashedPwd := base64.StdEncoding.EncodeToString(hashedPwd)
	user.Password_tmp = ""
	user.Password = b64HashedPwd
	model.UpdateUser(user)
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}

func TaxesGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("taxes.html").
		Funcs(template.FuncMap{"add": Add}).
		Funcs(template.FuncMap{"minus": Minus}).
		ParseFiles("template/taxes.html", head, navigation, footer))
	users, _ := model.GetAllUsers()
	data := Data{
		Session: GetSessionInformation(r),
		Users:   users,
	}
	tmpl.Execute(w, data)
}

func TaxesPOST(w http.ResponseWriter, r *http.Request) {
	users, _ := model.GetAllUsers()
	for _, u := range users {
		user := model.Map2User(u)
		if pa := ParseInt(r.FormValue(user.Id + "-PA")); pa >= 0 {
			user.Taxes.Previous_week.Amount = pa
		}
		if ca := ParseInt(r.FormValue(user.Id + "-CA")); ca >= 0 {
			user.Taxes.Current_week.Amount = ca
		}
		if na := ParseInt(r.FormValue(user.Id + "-NA")); na >= 0 {
			user.Taxes.Next_week.Amount = na
		}
		if sna := ParseInt(r.FormValue(user.Id + "-SNA")); sna >= 0 {
			user.Taxes.Second_next_week.Amount = sna
		}
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
	http.Redirect(w, r, "/members/taxes", http.StatusFound)
}

func UserDeleteGET(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	user, err := model.GetUserById(id)
	if err == nil {
		http.Redirect(w, r, "/members/id?userError=delete", http.StatusFound)
	}
	model.DeleteUser(user)
	http.Redirect(w, r, "/members", http.StatusFound)
}
