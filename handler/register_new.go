package handler

import (
	"net/http"
	// "../domain/model"
)

func RegisterNew(w http.ResponseWriter, r *http.Request) {
	log(authenticate(registerNew))(w, r)
}

func registerNew(w http.ResponseWriter, r *http.Request) {
	renderHTML(w, r, nil, NewTemplateOption(), "register/new")
}
