package session

import "sync"

type SimpleSessions struct {
	Username string
	TTL int64
}

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func LoadSessionFromDB() {

}

func GenerateNewSessionId(un string) string {

}

func IsSessionExpired(sid string) (string, bool) {

}
