package test

import (
	"go-sample/pocket52/redis"
	"reflect"
	"testing"
)

func Test_Sadd_001(t *testing.T) {

	var r redis.RedisStringSet

	r.RedisSet = append(r.RedisSet, "a", "b")
	r.Sadd("c")
	r.Sadd("b")
	result := []string{"a", "b", "c"}

	if reflect.DeepEqual(r.RedisSet, result) {
		t.Log("Success")
	} else {
		t.Error("Failed Sadd ")
	}
}

func Test_Scard_002(t *testing.T) {

	var r redis.RedisStringSet

	r.RedisSet = append(r.RedisSet, "a", "b")

	if r.Scard() == 2 {
		t.Log("Success")
	} else {
		t.Error("Failed Scard ")
	}
}

func Test_Srem_003(t *testing.T) {

	var r redis.RedisStringSet

	r.RedisSet = append(r.RedisSet, "a", "b", "c")

	if r.Srem("c") == 1 && r.Srem("d") == 0 {
		t.Log("Success")
	} else {
		t.Error("Failed Srem ")
	}
}

func Test_Sunion_004(t *testing.T) {

	var r redis.RedisStringSet

	m := make(map[string][]string)
	m["1"] = []string{"a", "b", "c"}
	m["2"] = []string{"d", "e", "f"}
	m["3"] = []string{"a", "e"}

	result := []string{"a", "b", "c", "d", "e", "f"}

	r.Sunion(m)

	if reflect.DeepEqual(r.RedisSet, result) {
		t.Log("Success")
	} else {
		t.Error("Failed Sunion ")
	}
}
