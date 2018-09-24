package handler

import (
	"net/http"

	"github.com/vsanna/go_web/usecase"
)

func JSONSample(w http.ResponseWriter, r *http.Request) {
	AuthorizeWrapper(json_sample)(w, r)
}

func json_sample(w http.ResponseWriter, r *http.Request) {
	users, err := usecase.NewJsonExample(repo.NewUserRepo()).Fetch(r.Context())
	if err != nil {
		http.Error(w, "cannot fetch users", http.StatusInternalServerError)
		return
	}

	renderJSON(w, r, users)
}

type userSample struct {
	ID    int
	Name  string
	Email string
}
