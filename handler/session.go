package handler

import (
	"net/http"

	"github.com/vsanna/go_web/lib/flash"
	"github.com/vsanna/go_web/usecase"
)

func Session(w http.ResponseWriter, r *http.Request) {
	session(w, r)
}

func session(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")

	err := usecase.NewSigninUsecase(repo.NewUserRepo()).Signin(r.Context(), email, password)

	if err != nil {
		flash.SetAlert(w, err.Error())
		http.Redirect(w, r, "/signin/new", http.StatusSeeOther)
		return
	}

	flash.SetNotice(w, "success!")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
