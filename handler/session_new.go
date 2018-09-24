package handler

import (
	"net/http"
)

func SessionNew(w http.ResponseWriter, r *http.Request) {
	NonAuthorizeWrapper(sessionNew)(w, r)
}

func sessionNew(w http.ResponseWriter, r *http.Request) {
	renderHTML(w, r, nil, NewTemplateOption(), "session/new")
}
