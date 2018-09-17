package handler

import (
	"net/http"
)

func ProfileUpdate(w http.ResponseWriter, r *http.Request) {
	log(authenticate(authorize(profileUpdate)))(w, r)
}

func profileUpdate(w http.ResponseWriter, r *http.Request) {
	user := CurrentUser(r.Context())

	r.ParseForm()
	name := user.Name
	if n := r.FormValue("name"); n != "" {
		name = n
	}

	email := user.Email
	if e := r.FormValue("email"); e != "" {
		email = e
	}

	password := r.FormValue("password")
	if password != "" {
		user.SetEncryptedPassword(password)
	}

	user.Name = name
	user.Email = email

	urepo := repo.NewUserRepo()
	_ = urepo.Update(r.Context(), user)

	http.Redirect(w, r, "/profile/", http.StatusFound)
}
