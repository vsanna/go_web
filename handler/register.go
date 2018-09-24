package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/vsanna/go_web/lib/flash"
	"github.com/vsanna/go_web/usecase"
	"github.com/vsanna/go_web/usecase/input"
)

func Register(w http.ResponseWriter, r *http.Request) {
	register(w, r)
}

func register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	input := &input.Register{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	user, err := usecase.NewRegisterUsecase(repo.NewUserRepo()).Register(r.Context(), input)

	fmt.Println(err)

	if err != nil {
		flash.SetAlert(w, err.Error())
		http.Redirect(w, r, "/register/new", http.StatusSeeOther)
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
	flash.SetNotice(w, "success!!")

	http.Redirect(w, r, "/", http.StatusFound)
}
