package handler

import (
	"fmt"
	"net/http"
)

func UserRequest(w http.ResponseWriter, r *http.Request) {
	userRequest(w, r)
}

func userRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		UserIndex(w, r)
	case "POST":
		// NOTE こんな感じでいいのかな？
		r.ParseForm()
		switch r.FormValue("_http_method") {
		case "PUT":
			// userUpdate(w, r)
		case "DELETE":
			// userDelete(w, r)
		default:
			// userCreate(w, r)
		}
	default:
		fmt.Fprintln(w, "not match http method")
	}
}
