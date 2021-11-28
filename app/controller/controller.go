package controller

import (
	"crypto/rand"
	"forgottennw/app/model"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
)

var (
	store      *sessions.CookieStore
	head       = "template/head.html"
	navigation = "template/navigation.html"
	footer     = "template/footer.html"
)

type Data struct {
	Session            Session
	RecruitmentEntries []map[string]interface{}
	RecruitmentEntry   model.RecruitmentEntry
}

type Session struct {
	LoggedIn bool
	User     model.User
}

func init() {
	key := make([]byte, 32)
	rand.Read(key)
	store = sessions.NewCookieStore(key)
}

func Authenticate(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")
		//Check if user is authentificated
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
		} else {
			h(w, r)
		}
	}
}

func GetSessionInformation(r *http.Request) (sessionInfo Session) {
	session, _ := store.Get(r, "session")
	userId := ""
	loggedIn := false
	if session.Values["userId"] != nil {
		userId = session.Values["userId"].(string)
	}
	if session.Values["authenticated"] != nil {
		loggedIn = session.Values["authenticated"].(bool)
	}
	user, user_err := model.GetUserById(userId)
	if user_err != nil {
		log.Println("couldn't find user in getSessionInformation()")
	}
	sessionInfo = Session{
		LoggedIn: loggedIn,
		User:     user,
	}
	return sessionInfo
}

func GetCurrentDate() (date model.Date) {
	date_Now := time.Now()
	date = model.Date{
		Second: date_Now.Second(),
		Minute: date_Now.Minute(),
		Hour:   date_Now.Hour(),
		Day:    date_Now.Day(),
		Month:  int(date_Now.Month()),
		Year:   date_Now.Year(),
	}
	return date
}

func ParseInt(str string) (result int) {
	result, _ = strconv.Atoi(str)
	return result
}
