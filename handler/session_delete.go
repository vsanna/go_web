package handler

import (
	"net/http"
)

func SessionDelete(w http.ResponseWriter, r *http.Request) {
	NonAuthorizeWrapper(sessionDelete)(w, r)
}

func sessionDelete(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{
		Name:     "_sid",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	}
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
