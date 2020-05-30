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

	//Endpoints for Redis List
	http.HandleFunc("/lpush", handlers.RedisListHandler)
	http.HandleFunc("/rpush", handlers.RedisListHandler)
	http.HandleFunc("/lpop", handlers.RedisListHandler)
	http.HandleFunc("/rpop", handlers.RedisListHandler)
	http.HandleFunc("/lrange", handlers.RedisListHandler)

	//Endpoints for Redis Strings
	http.HandleFunc("/append", handlers.RedisStringHandler)
	http.HandleFunc("/setnx", handlers.RedisStringHandler)
	http.HandleFunc("/incr", handlers.RedisStringHandler)
	http.HandleFunc("/getRange", handlers.RedisStringHandler)

	//Endpoints for Redis Set
	http.HandleFunc("/sadd", handlers.RedisSetHandler)
	http.HandleFunc("/scard", handlers.RedisSetHandler)

	//local server to run in port 8080
	log.Fatal(http.ListenAndServe(Port, nil))
}
