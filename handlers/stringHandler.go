package handlers

import (
	"go-sample/pocket52/controllers"
	"go-sample/pocket52/utility"
	"net/http"
	"strings"
)

//RedisHandler... Handler Function for Redis
func RedisStringHandler(w http.ResponseWriter, r *http.Request) {

	switch true {
	case strings.Contains(r.URL.Path, "/append") && r.Method == "GET":
		data := r.URL.Query().Get("data")
		utility.ReturnResponse(w, controllers.AppendString(data))
	case strings.Contains(r.URL.Path, "/setnx") && r.Method == "GET":
		keyValue := r.URL.Query().Get("keyValue")
		data := r.URL.Query().Get("data")
		utility.ReturnResponse(w, controllers.SetIfNotExist(keyValue, data))
	case strings.Contains(r.URL.Path, "/incr") && r.Method == "GET":
		data := r.URL.Query().Get("data")
		utility.ReturnResponse(w, controllers.Incr(data))
	case strings.Contains(r.URL.Path, "/getRange") && r.Method == "GET":
		data := r.URL.Query().Get("data")
		start := r.URL.Query().Get("start")
		stop := r.URL.Query().Get("stop")
		utility.ReturnResponse(w, controllers.GetRange(data, start, stop))
	default:
		utility.ReturnResponse(w, "Invalid Endpoint")

	}

}
