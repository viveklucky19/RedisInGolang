package controllers

import (
	"go-sample/pocket52/redis"
	"go-sample/pocket52/utility"

	"github.com/spf13/cast"
)

var RedisString string

type StringResponse struct {
	RedisString       string `json:"string"`
	OperationResponse string `json:"response"`
}

func AppendString(data string) utility.ReturnJson {

	var response StringResponse
	response.OperationResponse = cast.ToString(redis.Append(&RedisString, data))
	response.RedisString = RedisString

	return utility.SetReturnData(nil, response)

}

func SetIfNotExist(key, data string) utility.ReturnJson {
	var response StringResponse

	response.OperationResponse = cast.ToString(redis.SetNx(&key, data))
	response.RedisString = key

	return utility.SetReturnData(nil, response)

}

func Incr(data string) utility.ReturnJson {
	var response StringResponse
	var err error

	response.RedisString = data
	response.OperationResponse, err = redis.INCR(&data)

	return utility.SetReturnData(err, response)

}

func GetRange(data string, start, end interface{}) utility.ReturnJson {
	var response StringResponse
	var err error

	response.RedisString = data
	response.OperationResponse, err = redis.GetRange(&data, start, end)

	return utility.SetReturnData(err, response)

}
