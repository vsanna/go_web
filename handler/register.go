package handler

import (
	"net/http"
	"time"

	"github.com/vsanna/go_web/domain/model"
)

func Register(w http.ResponseWriter, r *http.Request) {
	log(authenticate(register))(w, r)
}

func register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" || password == "" || name == "" {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	user, err := model.NewUser(name, email, password)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// NOTE 手を抜いてerrorチェック省略
	urepo := repo.NewUserRepo()
	err = urepo.Create(r.Context(), user)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	c := &http.Cookie{
		Name:     "_sid",
		Value:    user.AccessToken,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().AddDate(0, 0, 14),
	}
	http.SetCookie(w, c)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
