package redis

import (
	"errors"
	"fmt"

	"github.com/spf13/cast"
)

//Operations interface that can be perform by data structures
type Operations interface {
	Lpush(data interface{}) error
	Rpush(data interface{}) error
	Lpop() error
	Rpop() error
	GetSize() int
}
type Redis struct {
	RedisList []int
}

func (r *Redis) GetSize() int {
	return len(r.RedisList)
}

//Left push operation
func (r *Redis) Lpush(data interface{}) error {
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
func (r *Redis) Rpush(data interface{}) error {
	d, err := cast.ToIntE(data)
	if err != nil {
		fmt.Println("Not Valid Integer")
		return err
	}
	r.RedisList = append(r.RedisList, d)
	return nil
}

//Lpop  operation
func (r *Redis) Lpop() error {
	if r.GetSize() == 0 {
		return errors.New("List is Empty")
	}
	r.RedisList = r.RedisList[1:r.GetSize()]
	return nil
}

//Rpop operation
func (r *Redis) Rpop() error {
	if r.GetSize() == 0 {
		return errors.New("List is Empty")
	}
	r.RedisList = r.RedisList[0 : r.GetSize()-1]
	return nil
}
