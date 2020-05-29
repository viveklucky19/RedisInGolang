package main

import (
	"go-sample/pocket52/handlers"
	"log"
	"net/http"
)

const (
	Port = ":8080"
)

func main() {

	//Endpoints
	http.HandleFunc("/lpush", handlers.RedisHandler)
	http.HandleFunc("/rpush", handlers.RedisHandler)
	http.HandleFunc("/lpop", handlers.RedisHandler)
	http.HandleFunc("/rpop", handlers.RedisHandler)

	//local server to run in port 8080
	log.Fatal(http.ListenAndServe(Port, nil))
}
