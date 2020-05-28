package main

import (
	"go-sample/pocket52/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/lpush", handlers.RedisHandler)
	http.HandleFunc("/rpush", handlers.RedisHandler)
	http.HandleFunc("/lpop", handlers.RedisHandler)
	http.HandleFunc("/rpop", handlers.RedisHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
