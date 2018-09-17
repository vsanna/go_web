package main

import (
	"net/http"

	"golang.org/x/net/http2"

	"./handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handler.Ping)
	mux.HandleFunc("/json", handler.JSONSample)

	mux.HandleFunc("/users/", handler.UserRequest)
	mux.HandleFunc("/profile/", handler.ProfileRequest)
	// mux.HandleFunc("/users/", handler.UserIndex)

	mux.HandleFunc("/register/new", handler.RegisterNew)
	mux.HandleFunc("/register/", handler.Register)
	mux.HandleFunc("/signin/new", handler.SessionNew)
	mux.HandleFunc("/signin/", handler.Session)
	mux.HandleFunc("/signout/", handler.SessionDelete)

	// NOTE 404処理はここでおこなう...? -> それっぽい動きはしている
	mux.HandleFunc("/", handler.Root)

	server := &http.Server{
		Addr:    "127.0.0.1:5000",
		Handler: mux,
	}

	http2.ConfigureServer(server, &http2.Server{})
	server.ListenAndServe()
}
