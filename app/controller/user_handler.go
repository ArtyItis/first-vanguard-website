package controller

import (
	"encoding/base64"
	"forgottennw/app/model"
	"html/template"
	"log"
	"net/http"

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
	username := r.FormValue("username")
	password := r.FormValue("password")
	//check if username is taken
	user_1, _ := model.GetUserByName(username)
	if user_1 != (model.User{}) {
		//TODO username is taken
		http.Redirect(w, r, "/username_is_taken", http.StatusFound)
		return
	}
	//hash password
	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	b64HashedPwd := base64.StdEncoding.EncodeToString(hashedPwd)

	user := model.User{
		Username:         username,
		Password:         b64HashedPwd,
		Permission_Level: 0,
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
	if user_err != nil || user == (model.User{}) {
		//TODO user not found
		log.Println("invalid username ")
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
		return
	}
	//validate password
	passwordDB, _ := base64.StdEncoding.DecodeString(user.Password)
	password_err := bcrypt.CompareHashAndPassword(passwordDB, []byte(password))
	if password_err != nil {
		//TODO wrong password
		log.Println("invalid password ", password_err)
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
		return
	}
	//save session
	session, _ := store.Get(r, "session")
	session.Values["authenticated"] = true
	session.Values["username"] = username
	session.Save(r, w)
	log.Println(username + " logged in")
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	username := session.Values["username"].(string)
	session.Values["authenticated"] = false
	session.Values["username"] = ""
	session.Save(r, w)
	log.Println(username + " logged out")
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}
