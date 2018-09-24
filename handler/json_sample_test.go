package handler_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"

	. "github.com/vsanna/go_web/handler"
)

var _ = Describe("JsonSample", func() {
	Context("if not signed in", func() {
		It("cannot get json", func() {
			mux := http.NewServeMux()
			mux.HandleFunc("/json", JSONSample)

			writer := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", "/json", nil)

			mux.ServeHTTP(writer, request)

			if writer.Code != http.StatusUnauthorized {
				GinkgoT().Errorf("Response Code is not 401, actual: %v", writer.Code)
			}

			responseBody := string(writer.Body.Bytes())
			if responseBody != "not authorized" {
				GinkgoT().Errorf("Response Body is not 'nont authorized', actual: %v", responseBody)
			}
		})
	})
})
