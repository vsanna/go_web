package handler

import (
	"net/http"

	"github.com/vsanna/go_web/usecase"
)

func ProfileShow(w http.ResponseWriter, r *http.Request) {
	AuthorizeWrapper(profileShow)(w, r)
}

func profileShow(w http.ResponseWriter, r *http.Request) {
	currentUser := CurrentUser(r.Context())
	if currentUser == nil {
		panic("check authorization")
	}

	profile, err := usecase.NewProfileUsecase(repo.NewUserRepo()).GetProfile(r.Context(), currentUser.ID)
	if err != nil {
		http.Error(w, "cannot fetch profile", http.StatusInternalServerError)
		return
	}

	renderHTML(w, r, profile, NewTemplateOption(), "profile/show")
}
