package test

import (
	"go-sample/pocket52/redis"
	"testing"
)

func Test_Append_001(t *testing.T) {

	var r string = "abc"

	redis.Append(&r, "efg")

	if r == "abcefg" {
		t.Log("Success")
	} else {
		t.Error("Failed Append")
	}
}

func Test_SetNx_002(t *testing.T) {

	var r string = ""
	s := "abc"

	if redis.SetNx(&r, "efg") == 1 && r == "efg" && redis.SetNx(&s, "efg") == 0 {
		t.Log("Success")
	} else {
		t.Error("Failed SetNx")
	}
}

func Test_Incr_003(t *testing.T) {

	var r string = "123"

	redis.INCR(&r)

	if r == "124" {
		t.Log("Success")
	} else {
		t.Error("Failed INCR")
	}
}

func Test_GetRange_004(t *testing.T) {

	var r string = "abcdefgh"

	s1, _ := redis.GetRange(&r, 0, 3)
	s2, _ := redis.GetRange(&r, 17, 4)
	s3, _ := redis.GetRange(&r, 0, -1)
	s4, _ := redis.GetRange(&r, -5, -2)

	if s1 == "abcd" && s2 == "" && s3 == "abcdefgh" && s4 == "defg" {
		t.Log("Success")
	} else {
		t.Error("Failed Getrange")
	}
}
