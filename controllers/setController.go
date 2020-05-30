package controllers

import (
	"go-sample/pocket52/redis"
	"go-sample/pocket52/utility"
)

type SetResponse struct {
	SetData []string `json:"set"`
	SetSize int      `json:"size"`
}

//Global variable for RedisSet
var GlobalSet redis.RedisStringSet

func SADD(s string) utility.ReturnJson {
	GlobalSet.Sadd(s)
	return utility.SetReturnData(nil, GlobalSet)
}

func SCARD(s string) utility.ReturnJson {

	var res SetResponse

	res.SetSize = GlobalSet.Scard(s)
	res.SetData = GlobalSet.RedisSet
	return utility.SetReturnData(nil, res)
}
