package controllers

import (
	"go-sample/pocket52/redis"
	"go-sample/pocket52/utility"
)

//Global variable for RedisList
var GlobalList redis.RedisListInt

//LeftPush... controller to call Redis List Lpush
func LeftPush(data interface{}) utility.ReturnJson {
	err := GlobalList.Lpush(data)
	return utility.SetReturnData(err, GlobalList)
}

//RightPush... controller to call Redis List Rpush
func RightPush(data interface{}) utility.ReturnJson {
	err := GlobalList.Rpush(data)
	return utility.SetReturnData(err, GlobalList)
}

//LeftPop... controller to call Redis List Lpop
func LeftPop() utility.ReturnJson {
	err := GlobalList.Lpop()
	return utility.SetReturnData(err, GlobalList)
}

//RightPop... controller to call Redis List Rpop
func RightPop() utility.ReturnJson {
	err := GlobalList.Rpop()
	return utility.SetReturnData(err, GlobalList)
}

//LRange... controller to call Redis List Lrange
func LRange(start, stop interface{}) (returnData utility.ReturnJson) {

	list, err := GlobalList.Lrange(start, stop)
	return utility.SetReturnData(err, list)
}
