package controllers

import (
	"encoding/json"
	"errors"
	"go-sample/pocket52/redis"
	"go-sample/pocket52/utility"
	"io/ioutil"
	"net/http"
)

type SetResponse struct {
	SetData              []string    `json:"set"`
	SetOperationResponse interface{} `json:"response"`
}

//Global variable for RedisSet
var GlobalSet redis.RedisStringSet

func SADD(s string) utility.ReturnJson {
	GlobalSet.Sadd(s)
	return utility.SetReturnData(nil, GlobalSet)
}

func SCARD() utility.ReturnJson {

	var res SetResponse

	res.SetOperationResponse = GlobalSet.Scard()
	res.SetData = GlobalSet.RedisSet
	return utility.SetReturnData(nil, res)
}

func SREM(s string) utility.ReturnJson {

	var res SetResponse

	res.SetOperationResponse = GlobalSet.Srem(s)
	res.SetData = GlobalSet.RedisSet
	return utility.SetReturnData(nil, res)
}

func SUNION(r *http.Request) utility.ReturnJson {

	var res SetResponse

	body, _ := ioutil.ReadAll(r.Body)

	data := make(map[string]interface{})

	err := json.Unmarshal(body, &data)
	if err != nil {
		return utility.SetReturnData(errors.New("Input not map of string and string slice Unmarshall"), nil)
	}

	err = GlobalSet.Sunion(data)
	res.SetData = GlobalSet.RedisSet
	return utility.SetReturnData(err, res)
}
