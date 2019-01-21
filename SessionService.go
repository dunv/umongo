package mongo

import (
	"crypto/tls"
	"log"
	"net"
	"strings"

	mgo "gopkg.in/mgo.v2"
)

// DbSession datastructure for a MongoDB session
type DbSession struct {
	session *mgo.Session
	dbName  string
}

// NewDbSession create new session
func NewDbSession(connectionString string, dbName string) (*DbSession, error) {
	log.Printf("Connecting to MongoDB...")

	// Manual TLS
	tlsConfig := &tls.Config{}
	tlsConfig.InsecureSkipVerify = true

	// Parse string
	dialInfo, err1 := mgo.ParseURL(connectionString)
	dialInfo.Direct = true
	if err1 != nil {
		return nil, err1
	}

	// TODO: this is still quite dirty -> if mongodb is in the
	// ConnectionString: assume we want to connect to Atlas via SSL
	if strings.Contains(connectionString, "mongodb") {
		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
			return conn, err
		}

	}

	dbSession, err2 := mgo.DialWithInfo(dialInfo)
	if err2 != nil {
		log.Printf("Error establishing connection %s \n", err2)
		return nil, err2
	}

	log.Printf("MongoDB connected.")
	return &DbSession{
		session: dbSession,
		dbName:  dbName,
	}, nil
}

// Copy a session
func (s *DbSession) Copy() *DbSession {
	if s.session == nil {
		log.Fatal("Session is nil?!")
		log.Println(s.session)
		log.Println(s)
	}
	return &DbSession{
		session: s.session.Copy(),
		dbName:  s.dbName,
	}
}

// GetCollection from a session
func (s *DbSession) GetCollection(col string) *mgo.Collection {
	return s.session.DB(s.dbName).C(col)
}

// Close mongoDB session
func (s *DbSession) Close() {
	if s.session != nil {
		s.session.Close()
	}
}
