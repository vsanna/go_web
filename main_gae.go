package main

import (
	"net/http"

	"github.com/vsanna/go_web/handler"
)

func init() {
	http.HandleFunc("/ping", handler.Ping)
	http.HandleFunc("/json", handler.JSONSample)

	http.HandleFunc("/users/", handler.UserRequest)
	http.HandleFunc("/profile/", handler.ProfileRequest)

	http.HandleFunc("/register/new", handler.RegisterNew)
	http.HandleFunc("/register/", handler.Register)

	http.HandleFunc("/signin/new", handler.SessionNew)
	http.HandleFunc("/signin/", handler.Session)
	http.HandleFunc("/signout/", handler.SessionDelete)
	// NOTE 404処理はここでおこなう...? -> それっぽい動きはしている
	http.HandleFunc("/", handler.Root)
}
