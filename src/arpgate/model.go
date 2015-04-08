package arpgate

import (
	"time"
)

const DB_COLL_USER = "arpgateUser"

var DB_NAME string

type User struct {
	Id        string    `json:"id"`
	Username  string    `json:"username"`
	Pwd       string    `json:"pwd"`
	Email     string    `json:"email"`
	IP        string    `json:"ip"`
	Confirmed bool      `json:"confirmed"`
	Heartbeat time.Time `json:"heartbeat"`
	Updated   time.Time `json:"updated"`
	Created   time.Time `json:"created"`
}

type DbConfig struct {
	MongodbHost string `json:"mongodb_host"`
	MongodbPwd  string `json:"mongodb_pwd"`
	MongodbUser string `json:"mongodb_user"`
	MongodbDB   string `json:"mongodb_db"`
}
