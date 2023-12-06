package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	A int `json:"a"`
	B int `json:"b"`
}

func main() {
	http.HandleFunc("/api/v1/data", func(w http.ResponseWriter, r *http.Request) {
		// CORS対応
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// レスポンスデータの作成
		res := Response{A: 100, B: 200}

		// JSONにエンコード
		json.NewEncoder(w).Encode(res)
	})

	http.ListenAndServe(":5173", nil)
}