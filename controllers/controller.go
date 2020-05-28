package controllers

import (
	"go-sample/pocket52/redis"
	"go-sample/pocket52/utility"
)


var GlobalList redis.Redis

func LeftPush(data interface{}) (returnData utility.ReturnJson) {
	err:=GlobalList.Lpush(data)	
	return utility.SetReturnData(err,GlobalList)
}

func RightPush(data interface{}) (returnData utility.ReturnJson) {
	err:=GlobalList.Rpush(data)	
	return utility.SetReturnData(err,GlobalList)
}

func LeftPop() (returnData utility.ReturnJson) {
	err:=GlobalList.Lpop()	
	return utility.SetReturnData(err,GlobalList)
}

func RightPop() (returnData utility.ReturnJson) {
	err:=GlobalList.Rpop()
	return utility.SetReturnData(err,GlobalList)
}