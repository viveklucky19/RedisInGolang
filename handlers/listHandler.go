package handlers

import (
	"go-sample/pocket52/controllers"
	"go-sample/pocket52/utility"
	"net/http"
	"strings"
)

//RedisHandler... Handler Function for Redis
func RedisListHandler(w http.ResponseWriter, r *http.Request) {

	switch true {
	case strings.Contains(r.URL.Path, "/lpush") && r.Method == utility.ConstMethodGET:
		data := r.URL.Query().Get("data")
		utility.ReturnResponse(w, controllers.LeftPush(data))
	case strings.Contains(r.URL.Path, "/rpush") && r.Method == utility.ConstMethodGET:
		data := r.URL.Query().Get("data")
		utility.ReturnResponse(w, controllers.RightPush(data))
	case strings.Contains(r.URL.Path, "/lpop") && r.Method == utility.ConstMethodGET:
		utility.ReturnResponse(w, controllers.LeftPop())
	case strings.Contains(r.URL.Path, "/rpop") && r.Method == utility.ConstMethodGET:
		utility.ReturnResponse(w, controllers.RightPop())
	case strings.Contains(r.URL.Path, "/lrange") && r.Method == utility.ConstMethodGET:
		start := r.URL.Query().Get("start")
		stop := r.URL.Query().Get("stop")
		utility.ReturnResponse(w, controllers.LRange(start, stop))
	default:
		utility.ReturnResponse(w, "Invalid Endpoint")

	}

}
