package session

import (
	"sync"
	"time"
	"video_server/dbops"
	"video_server/defs"
	"video_server/utils"
)

type SimpleSessions struct {
	Username string
	TTL int64
}

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func LoadSessionFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}
	r.Range(func(k, v interface{}) bool {
		session := v.(*defs.SimpleSession)
		sessionMap.Store(k, session)
		return true
	})
}

func GenerateNewSessionId(un string) string {
	id, _ := utils.NewUUID()
	createTime := nowInMilli()/1000000
	ttl := createTime + 30 * 60 * 1000 // 30分钟

	session := &defs.SimpleSession{Username:un, TTL: ttl}
	sessionMap.Store(id, session)
	dbops.InserSession(id, ttl, un)
	return id
}

func IsSessionExpired(sid string) (string, bool) {
	session, ok := sessionMap.Load(sid)
	if ok {
		createTime := nowInMilli()
		if session.(*defs.SimpleSession).TTL < createTime {
			deleteExpiredSession(sid)
			return "", true
		}
		return session.(*defs.SimpleSession).Username, false
	}
	return "", true
}

func nowInMilli() int64 {
	return time.Now().UnixNano()
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}
