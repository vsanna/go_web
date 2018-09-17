package handler

import (
	"fmt"
	"net/http"

	"github.com/vsanna/go_web/domain/model"
)

func UserIndex(w http.ResponseWriter, r *http.Request) {
	log(authenticate(authorize(userIndex)))(w, r)
}

func userIndex(w http.ResponseWriter, r *http.Request) {

	userrepo := repo.NewUserRepo()
	users, err := userrepo.All(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "not found")
		return
	}

	vals := struct {
		Title string
		Users []*model.User
	}{
		Title: "user#index",
		Users: users,
	}
	renderHTML(w, r, vals, NewTemplateOption(), "users/index")
}
