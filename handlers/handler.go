package handlers

import (
	"encoding/json"
	"go-sample/pocket52/controllers"
	"net/http"
	"strings"
)

func RedisHandler(w http.ResponseWriter, r *http.Request) {

	switch true {
	case strings.Contains(r.URL.Path, "/lpush") && r.Method == "GET":
		data := r.URL.Query().Get("data")
		ReturnResponse(w, controllers.LeftPush(data))
	case strings.Contains(r.URL.Path, "/rpush") && r.Method == "GET":
		data := r.URL.Query().Get("data")
		ReturnResponse(w, controllers.RightPush(data))
	case strings.Contains(r.URL.Path, "/lpop") && r.Method == "GET":
		ReturnResponse(w, controllers.LeftPop())
	case strings.Contains(r.URL.Path, "/rpop") && r.Method == "GET":
		ReturnResponse(w, controllers.RightPop())

	}

}

func ReturnResponse(w http.ResponseWriter, data interface{}) {
	returnJson, _ := json.Marshal(data)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(returnJson)
}
