package main

import (
	"forgottennw/app/model"
	"forgottennw/app/route"
	"log"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	//schedule task to update taxes of every user once a week
	s := gocron.NewScheduler(time.Local)
	weeklyJob, _ := s.Every(1).Week().Monday().At("04:00").Do(func() {
		users, _ := model.GetAllUsers()
		for _, u := range users {
			user := model.Map2User(u)
			model.ShiftTaxes(user)
			model.UpdateUser(user)
		}
	})
	weekday, _ := weeklyJob.Weekday()
	log.Println("updating taxes every " + weekday.String() + ", 04:00am")
	s.StartAsync()

	router := route.NewRouter()
	log.Println("Listening...")
	http.ListenAndServe(":8080", router)
}
