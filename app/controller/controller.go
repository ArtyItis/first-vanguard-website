package controller

import (
	"crypto/rand"
	"forgottennw/app/model"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/sessions"
)

var (
	store      *sessions.CookieStore
	head       = "template/assets/head.html"
	navigation = "template/assets/navigation.html"
	footer     = "template/assets/footer.html"
)

type Data struct {
	Session      Session
	Applications []map[string]interface{}
	Roles        []map[string]interface{}
	Weapons      []map[string]interface{}
	Users        []map[string]interface{}
	User         model.User
	Application  model.Application
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
			http.Redirect(w, r, "/?loginError=notLoggedIn", http.StatusFound)
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
	_, week := date_Now.ISOWeek()

	date = model.Date{
		Second: date_Now.Second(),
		Minute: date_Now.Minute(),
		Hour:   date_Now.Hour(),
		Day:    date_Now.Day(),
		Month:  int(date_Now.Month()),
		Year:   date_Now.Year(),
		Week:   week,
	}
	return date
}

func ParseInt(str string) (result int) {
	result, _ = strconv.Atoi(str)
	return result
}

func GetPreviousRoute(r *http.Request) string {
	previous_route := r.Header.Get("Referer")
	if strings.Contains(previous_route, "?") {
		previous_route = strings.Split(previous_route, "?")[0]
	}
	return previous_route
}
