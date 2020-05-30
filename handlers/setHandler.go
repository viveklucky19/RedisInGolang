package handlers

import (
	"go-sample/pocket52/controllers"
	"go-sample/pocket52/utility"
	"net/http"
	"strings"
)

//RedisHandler... Handler Function for Redis
func RedisSetHandler(w http.ResponseWriter, r *http.Request) {

	switch true {
	case strings.Contains(r.URL.Path, "/sadd") && r.Method == "GET":
		data := r.URL.Query().Get("data")
		utility.ReturnResponse(w, controllers.SADD(data))
	case strings.Contains(r.URL.Path, "/scard") && r.Method == "GET":
		data := r.URL.Query().Get("data")
		utility.ReturnResponse(w, controllers.SCARD(data))

	default:
		utility.ReturnResponse(w, "Invalid Endpoint")

	}

}
