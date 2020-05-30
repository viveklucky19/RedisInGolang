package redis

import (
	"errors"
	"fmt"

	"github.com/spf13/cast"
)

//Operations interface that can be perform by data structures
type Operations interface {
	GetSize() int
	Lpush(data interface{}) error
	Rpush(data interface{}) error
	Lpop() error
	Rpop() error
	Lrange(start, end int) (interface{}, error)
}

//for list of integers
type RedisListInt struct {
	RedisList []int
}

//for list of strings
type RedisListString struct {
	RedisList []string
}

func (r *RedisListInt) GetSize() int {
	return len(r.RedisList)
}

//Left push operation
func (r *RedisListInt) Lpush(data interface{}) error {
	d, err := cast.ToIntE(data)
	if err != nil {
		fmt.Println("Not Valid Integer")
		return err
	}
	newList := []int{d}
	r.RedisList = append(newList, r.RedisList...)
	return nil

}

//Right push operation
func (r *RedisListInt) Rpush(data interface{}) error {
	d, err := cast.ToIntE(data)
	if err != nil {
		fmt.Println("Not Valid Integer")
		return err
	}
	r.RedisList = append(r.RedisList, d)
	return nil
}

//Lpop  operation
func (r *RedisListInt) Lpop() error {
	if r.GetSize() == 0 {
		return errors.New("List is Empty")
	}
	r.RedisList = r.RedisList[1:r.GetSize()]
	return nil
}

//Rpop operation
func (r *RedisListInt) Rpop() error {
	if r.GetSize() == 0 {
		return errors.New("List is Empty")
	}
	r.RedisList = r.RedisList[0 : r.GetSize()-1]
	return nil
}

//Lrange operation
func (r *RedisListInt) Lrange(start, stop interface{}) (interface{}, error) {

	if start == nil || stop == nil {
		return nil, errors.New("Start or stop can't be empty")
	}
	s, err := cast.ToIntE(start)
	if err != nil {
		return nil, errors.New("Invalid Int Value for start")
	}

	e, err := cast.ToIntE(stop)
	if err != nil {
		return nil, errors.New("Invalid Int Value for stop")
	}

	listSize := r.GetSize()

	//if list is or empty  return list
	if listSize == 0 {
		return r.RedisList, errors.New("List is Empty")
	}

	//if start >=length list => return empty list
	if s >= listSize || s > e {
		return []int{}, nil
	}

	//is end >= listSize, return list from s to listSize
	if e >= listSize {
		return r.RedisList[s:], nil
	}

	return r.RedisList[s : e+1], nil
}

//similarly all the operation can be implemented by string list
