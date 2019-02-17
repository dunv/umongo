package umongo

import (
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

// SetDebug <-
func (s *DbSession) SetDebug(enable bool) {
	mgo.SetDebug(true)
	var aLogger *log.Logger
	aLogger = log.New(os.Stderr, "", log.LstdFlags)
	mgo.SetLogger(aLogger)
}
