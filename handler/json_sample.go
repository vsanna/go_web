package handler

import (
	"net/http"
)

func JSONSample(w http.ResponseWriter, r *http.Request) {
	log(authenticate(authorize(json_sample)))(w, r)
}

func json_sample(w http.ResponseWriter, r *http.Request) {
	urepo := repo.NewUserRepo()
	users, _ := urepo.All(r.Context())

	// pluck処理. gormにはPluckがあるみたい
	data := []userSample{}
	for _, u := range users {
		data = append(data, userSample{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
		})
	}
	renderJSON(w, r, data)
}

type userSample struct {
	ID    int
	Name  string
	Email string
}
