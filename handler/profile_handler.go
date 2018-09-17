package handler

import (
	"fmt"
	"net/http"
)

func ProfileRequest(w http.ResponseWriter, r *http.Request) {
	profileRequest(w, r)
}

func profileRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ProfileShow(w, r)
	case "POST":
		r.ParseForm()
		// NOTE 適当に作った。ええのか？
		switch r.FormValue("_http_method") {
		case "PUT":
			ProfileUpdate(w, r)
		case "DELETE":
			// userDelete(w, r)
		default:
			// userCreate(w, r)
		}
	default:
		fmt.Fprintln(w, "not match http method")
	}
}
