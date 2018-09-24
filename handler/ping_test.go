package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	// テストするリクエスト(とそれにだけ対応するサーバー)
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", Ping)
	req, _ := http.NewRequest("GET", "/ping", nil)

	// サーバーが返すresponseを受け取るハコ
	writer := httptest.NewRecorder()

	// リクエスト実行
	mux.ServeHTTP(writer, req)

	// レスポンスを見てテストの結果を記述する
	if writer.Code != 200 {
		t.Errorf("Response Code is not 200, actual: %v", writer.Code)
	}

	responseBody := string(writer.Body.Bytes())
	if responseBody != "pong" {
		t.Errorf("Response is not 'pong', actual: %v", responseBody)
	}
}
