package main

import (
	"net/http"
	"video_server/session"
)

var HEADER_FIELD_SESSION  = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"

func validateUserSession(r *http.Request) bool {
	sessionId := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sessionId) == 0 {
		return false
	}
	uname, ok := session.IsSessionExpired(sessionId)
	if ok {
		return false
	}
	r.Header.Add(HEADER_FIELD_UNAME, uname)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FIELD_UNAME)
	if len(uname) == 0 {
		sendErrorResponse()
		return false
	}
	return true
}
