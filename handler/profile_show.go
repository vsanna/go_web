package handler

import (
	"net/http"
)

func ProfileShow(w http.ResponseWriter, r *http.Request) {
	log(authenticate(authorize(profileShow)))(w, r)
}

func profileShow(w http.ResponseWriter, r *http.Request) {
	vals := struct {
		Name  string
		Email string
	}{
		Name:  CurrentUser(r.Context()).Name,
		Email: CurrentUser(r.Context()).Email,
	}

	renderHTML(w, r, vals, NewTemplateOption(), "profile/show")
}
