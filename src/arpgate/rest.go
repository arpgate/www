package arpgate

import (
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
	"log"
	"net/http"
	"time"
)

func SetupRestListener() {

	m := martini.Classic()

	// TODO
	// Validate email
	// Delete Account
	// Update User
	// Register IP
	// Query IP

	// Security - add Captcha

	// Emailer

	m.Post("/v1/check", binding.Bind(User{}), func(user User) string {
		if CheckUserPwd(user) {
			return "{status:'true'}"
		} else {
			return "{status:'false'}"
		}
	})

	m.Post("/v1/register", binding.Bind(User{}), func(user User) string {
		if CheckUserExists(user) {
			return "{status:'USEREXISTS'}"
		}

		t := time.Now()
		user.Id = NewUUID()
		user.Heartbeat = t
		user.Updated = t
		user.Created = t

		if InsertUser(user) {
			return "{status:'true'}"
		} else {
			return "{status:'false'}"
		}
	})

	log.Fatal(http.ListenAndServe(":8080", m))
}
