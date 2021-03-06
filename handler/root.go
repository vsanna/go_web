package handler

import (
	"fmt"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	NonAuthorizeWrapper(root)(w, r)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Println("==========")
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "no page")
		return
	}

	renderHTML(w, r, nil, NewTemplateOption(), "pages/root")
}
