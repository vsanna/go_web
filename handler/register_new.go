package handler

import (
	"net/http"
)

func RegisterNew(w http.ResponseWriter, r *http.Request) {
	NonAuthorizeWrapper(registerNew)(w, r)
}

func registerNew(w http.ResponseWriter, r *http.Request) {
	renderHTML(w, r, nil, NewTemplateOption(), "register/new")
}
