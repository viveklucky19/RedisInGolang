package controllers

import (
	"go-sample/pocket52/redis"
	"go-sample/pocket52/utility"
)

//Global variable for RedisList
var GlobalList redis.Redis


//LeftPush... controller to call Redis Lpush
func LeftPush(data interface{}) (returnData utility.ReturnJson) {
	err:=GlobalList.Lpush(data)	
	return utility.SetReturnData(err,GlobalList)
}

//RightPush... controller to call Redis Rpush
func RightPush(data interface{}) (returnData utility.ReturnJson) {
	err:=GlobalList.Rpush(data)	
	return utility.SetReturnData(err,GlobalList)
}

//LeftPop... controller to call Redis Lpop
func LeftPop() (returnData utility.ReturnJson) {
	err:=GlobalList.Lpop()	
	return utility.SetReturnData(err,GlobalList)
}

//RightPop... controller to call Redis Rpop
func RightPop() (returnData utility.ReturnJson) {
	err:=GlobalList.Rpop()
	return utility.SetReturnData(err,GlobalList)
}