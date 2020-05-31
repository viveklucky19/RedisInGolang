package redis

import (
	"errors"
	"fmt"
	"go-sample/pocket52/utility"

	"github.com/spf13/cast"
)

type SetOperations interface {
	Sadd(s string) interface{}
	Scard() int
	Srem(s string) int
	Sunion(m interface{}) error
}

type RedisIntSet struct {
	RedisSet []int
}

type RedisStringSet struct {
	RedisSet []string
}

func (r *RedisStringSet) Sadd(s string) interface{} {

	//check if s is already present in set

	if isPresent, _ := utility.IsValuePresentInSlice(r.RedisSet, s); isPresent {
		return r.RedisSet
	}

	r.RedisSet = append(r.RedisSet, s)

	return r.RedisSet

}

func (r *RedisStringSet) Scard() int {
	return len(r.RedisSet)
}

func (r *RedisStringSet) Srem(s string) int {

	newSet := []string{}
	index := -1
	//check if s is already present in set

	if isPresent, i := utility.IsValuePresentInSlice(r.RedisSet, s); isPresent {
		index = i
	}

	if index != -1 {
		newSet = append(newSet, r.RedisSet[0:index]...)
		newSet = append(newSet, r.RedisSet[index+1:]...)

		r.RedisSet = newSet

		return 1
	}

	return 0

}

func (r *RedisStringSet) Sunion(m interface{}) error {

	fmt.Println("m = ", m)

	data, err := cast.ToStringMapStringSliceE(m)
	if err != nil {
		return errors.New("Input not map of string and string slice from String ")
	}

	for _, v := range data {
		for _, val := range v {
			r.Sadd(val)
		}
	}

	return nil
}
