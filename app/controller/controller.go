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
	Session             Session
	Recruitment_Entries []map[string]interface{}
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

func GetSessionInformation(r *http.Request) (sessionInfo Session) {
	session, _ := store.Get(r, "session")
	username := ""
	loggedIn := false
	if session.Values["username"] != nil {
		username = session.Values["username"].(string)
	}
	if session.Values["authenticated"] != nil {
		loggedIn = session.Values["authenticated"].(bool)
	}
	user, user_err := model.GetUserByName(username)
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
		Date:   date_Now.Local(),
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
