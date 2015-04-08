package arpgate

import (
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
)

var conf DbConfig

func CheckUserExists(user User) bool {
	existingUser := User{}
	dbUrl := LoadDBConfig()
	session, err := mgo.Dial(dbUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB(DB_NAME).C(DB_COLL_USER)
	c.Find(bson.M{"username": user.Username}).One(&existingUser)
	if len(existingUser.Id) > 0 {
		return true
	}
	return false
}

func CheckUserPwd(user User) bool {
	result := User{}
	dbUrl := LoadDBConfig()
	session, err := mgo.Dial(dbUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB(DB_NAME).C(DB_COLL_USER)
	c.Find(bson.M{"username": user.Username}).One(&result)
	if user.Pwd == result.Pwd {
		return true
	}
	return false
}

func InsertUser(user User) bool {
	dbUrl := LoadDBConfig()
	session, err := mgo.Dial(dbUrl)
	if err != nil {
		return false
	}
	defer session.Close()
	c := session.DB(DB_NAME).C(DB_COLL_USER)
	c.Insert(user)
	return true
}

func LoadDBConfig() string {
	content, err := ioutil.ReadFile("/etc/arpgate.conf")
	if err != nil {
		log.Fatal("Error:", err)
		return ""
	}
	json.Unmarshal(content, &conf)
	DB_NAME = conf.MongodbDB
	return "mongodb://" + conf.MongodbUser + ":" + conf.MongodbPwd + "@" + conf.MongodbHost + "/" + conf.MongodbDB
}
