package main

import (
	"fetch-saldo/src/handler"
	"fetch-saldo/src/helper"
	"fmt"
	"log"
	"net/http"
	"time"
)

func init() {
	helper.LoadEnv()
	helper.ConnectDB()
}

func main() {
	rateLimiter := helper.NewRateLimiter(10, 10*time.Second)

	mux := http.NewServeMux()

	type route struct {
		Method  string
		Path    string
		Handler http.HandlerFunc
	}
	routes := []route{
		{http.MethodPost, "/api/add-api-key", handler.AddApiKey},
		{http.MethodPost, "/api/get-balances", handler.GetBalance},
	}

	for _, route := range routes {
		mux.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != route.Method {
				http.Error(w, `{"error": "method not allowed"}`, http.StatusMethodNotAllowed)
				return
			}
			helper.WithRateLimit(rateLimiter, route.Handler)(w, r)
		})
	}

	fmt.Println("Server started at http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", mux))
}
