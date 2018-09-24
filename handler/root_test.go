package handler

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// 各テストでの共通処理(前処理/後処理)を書く
func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	// tearDown()
	os.Exit(code)
}

var (
	writer *httptest.ResponseRecorder
	mux    *http.ServeMux
)

func setUp() {
	mux = http.NewServeMux()
	mux.HandleFunc("/", Root)
	os.Setenv("GOAPP_ENV", "test")
}

func TestRoot(t *testing.T) {
	writer = httptest.NewRecorder()
	// setUpと合わせて、テストするリクエスト(とそれにだけ対応するサーバー)を用意
	req, _ := http.NewRequest("GET", "/", nil)

	// リクエスト実行
	mux.ServeHTTP(writer, req)

	// レスポンスを見てテストの結果を記述する
	if writer.Code != 200 {
		t.Errorf("Response Code is not 200, actual: %v", writer.Code)
	}
}

func TestNotFound(t *testing.T) {
	writer = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/no_exist", nil)
	mux.ServeHTTP(writer, req)

	if writer.Code != 404 {
		t.Errorf("Response Code is not 404, actual: %v", writer.Code)
	}

	responseBody := string(writer.Body.Bytes())
	if responseBody != "no page" {
		t.Errorf("Response is not 'no page', actual: %v", responseBody)
	}
}
