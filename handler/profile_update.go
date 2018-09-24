package handler

import (
	"net/http"

	"github.com/vsanna/go_web/usecase"
	"github.com/vsanna/go_web/usecase/input"
)

func ProfileUpdate(w http.ResponseWriter, r *http.Request) {
	AuthorizeWrapper(profileUpdate)(w, r)
}

func profileUpdate(w http.ResponseWriter, r *http.Request) {
	user := CurrentUser(r.Context())
	if user == nil {
		panic("broken")
	}

	r.ParseForm()

	updateUserInput := &input.UpdateUser{
		User:     user,
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	err := usecase.NewUpdateUserUsecase(repo.NewUserRepo()).Update(r.Context(), updateUserInput)
	if err != nil {
		http.Error(w, "failed to update", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/profile/", http.StatusFound)
}
