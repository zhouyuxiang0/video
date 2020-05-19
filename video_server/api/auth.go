package main

import (
	"net/http"
	"video_server/defs"
	"video_server/session"
)

var headerFieldSession = "X-Session-Id"
var headerFieldUname = "X-User-Name"

func validateUserSession(r *http.Request) bool {
	sessionID := r.Header.Get(headerFieldSession)
	if len(sessionID) == 0 {
		return false
	}
	uname, ok := session.IsSessionExpired(sessionID)
	if ok {
		return false
	}
	r.Header.Add(headerFieldUname, uname)
	return true
}

func validateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(headerFieldUname)
	if len(uname) == 0 {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}
	return true
}
