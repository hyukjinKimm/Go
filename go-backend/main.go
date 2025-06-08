package main

import (
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"message": "pong"}`)
}

func main() {
	http.HandleFunc("/ping", pingHandler) // /ping 경로에 핸들러 등록
	http.ListenAndServe(":8080", nil)     // 8080 포트에서 서버 실행
}
