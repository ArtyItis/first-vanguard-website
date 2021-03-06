package controller

import (
	"crypto/rand"
	"first-vanguard/app/model"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	store      *sessions.CookieStore
	head       = "template/assets/head.html"
	navigation = "template/assets/navigation.html"
	footer     = "template/assets/footer.html"

	attributes = "template/assets/character/attributes.html"
	jobs       = "template/assets/character/jobs.html"
	roles      = "template/assets/character/roles.html"
	weapons    = "template/assets/character/weapons.html"
)

type Data struct {
	Session      Session
	Applications []map[string]interface{}
	Application  model.Application
	Roles        []map[string]interface{}
	Role         model.Role
	Weapons      []map[string]interface{}
	Users        []map[string]interface{}
	User         model.User
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
		if auth, ok := session.Values["authenticated"].(bool); !auth || !ok {
			http.Redirect(w, r, "/?error=authentication&type=login", http.StatusFound)
			return
		}
		if companyRoute := mux.Vars(r)["company"]; companyRoute != "" {
			if companySession, ok := session.Values["userCompany"]; companySession.(string) != companyRoute || !ok {
				http.Redirect(w, r, "/?error=authentication&type=company", http.StatusFound)
				return
			}
		}
		h(w, r)
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

//inspired by: https://tutorialedge.net/golang/go-file-upload-tutorial/
//returns "empty" if no file was found
func UploadFile(r *http.Request, formInputKey string, directory string, maxFileSize int64) string {
	// Parse our multipart form, example: 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(maxFileSize << 20)
	// FormFile returns the first file for the given key `formInputKey`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile(formInputKey)
	if err != nil {
		log.Println(err)
		return "empty"
	}
	defer file.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}
	dirPath := "static/storage/" + directory
	//ensure that all directories and subdirectories exist
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		log.Println(err)
	}
	// write this byte array to our file
	filePath := dirPath + "/" + handler.Filename
	//create file
	err = ioutil.WriteFile(filePath, fileBytes, 0644)

	if err != nil {
		panic(err)
	}
	return "/" + filePath
}
