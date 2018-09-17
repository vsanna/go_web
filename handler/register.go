package handler

import (
	"fmt"
	"net/http"
	"time"

	"../domain/model"
	// "../domain/model"
)

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("debug: pass here")
	log(authenticate(register))(w, r)
}

func register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	fmt.Println("debug: ", name, email, password)

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
	fmt.Println("debug: ", err)
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
