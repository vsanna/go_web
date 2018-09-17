package handler

import (
	"net/http"
	"time"

	"github.com/vsanna/go_web/lib"
)

func Session(w http.ResponseWriter, r *http.Request) {
	log(authenticate(session))(w, r)
}

func session(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")
	if email == "" || password == "" {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	user, err := repo.NewUserRepo().FindByEmail(r.Context(), email)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	if lib.Compare(user.EncryptedPassword, password) {
		c := &http.Cookie{
			Name:     "_sid",
			Value:    user.AccessToken,
			Path:     "/",
			Domain:   host,
			HttpOnly: true,
			Expires:  time.Now().AddDate(0, 0, 14),
		}
		http.SetCookie(w, c)
	}

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
