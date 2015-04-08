package main

/*
install packages
go get github.com/go-martini/martini
go get  gopkg.in/mgo.v2
go get gopkg.in/mgo.v2/bson
go get github.com/codegangsta/martini-contrib/binding
*/
import (
	"arpgate"
)

func main() {

	// start API Listener
	arpgate.SetupRestListener()
}
