package main

import (
	"log"
	"net/http"
	"server/handler"
	"server/middleware"
	"time"
)

func main() {
	server := http.Server{
		Addr:         ":8000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	http.HandleFunc("/", handler.Root)
	http.HandleFunc("/hello",
		middleware.Set(
			handler.HelloWorld,
			middleware.LogMiddleware,
		))

	log.Println("server is listening on port: 8000")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("failed to start the server. error: ", err)
	}
}
