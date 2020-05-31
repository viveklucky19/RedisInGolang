package test

import (
	"go-sample/pocket52/redis"
	"reflect"
	"testing"
)

func Test_Lpush_001(t *testing.T) {

	var r redis.RedisListInt
	r.RedisList = append(r.RedisList, 1, 2, 3, 4)
	result := []int{5, 1, 2, 3, 4}
	r.Lpush(5)
	if reflect.DeepEqual(r.RedisList, result) {
		t.Log("Success")
	} else {
		t.Error("Failed Lpush")
	}
}

func Test_Rpush_002(t *testing.T) {
	var r redis.RedisListInt
	r.RedisList = append(r.RedisList, 1, 2, 3, 4)
	result := []int{1, 2, 3, 4, 5}
	r.Rpush(5)
	if reflect.DeepEqual(r.RedisList, result) {
		t.Log("Success")
	} else {
		t.Error("Failed Rpush")
	}
}

func Test_Lpop_003(t *testing.T) {
	var r redis.RedisListInt
	r.RedisList = append(r.RedisList, 1, 2, 3, 4)
	result := []int{2, 3, 4}
	r.Lpop()
	if reflect.DeepEqual(r.RedisList, result) {
		t.Log("Success")
	} else {
		t.Error("Failed Lpop")
	}
}

func Test_Rpop_004(t *testing.T) {
	var r redis.RedisListInt
	r.RedisList = append(r.RedisList, 1, 2, 3, 4)
	result := []int{1, 2, 3}
	r.Rpop()
	if reflect.DeepEqual(r.RedisList, result) {
		t.Log("Success")
	} else {
		t.Error("Failed Rpop")
	}
}
