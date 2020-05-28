package redis

import (
	"errors"
	"fmt"

	"github.com/spf13/cast"
)

type Operations interface {
	Lpush(data interface{}) error
	Rpush(data interface{}) error
	Lpop() error
	GetSize() int
}

type Redis struct {
	RedisList []int
}

func (r *Redis) GetSize() int {
	return len(r.RedisList)
}

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

func (r *Redis) Rpush(data interface{}) error {
	d, err := cast.ToIntE(data)
	if err != nil {
		fmt.Println("Not Valid Integer")
		return err
	}
	r.RedisList = append(r.RedisList, d)
	return nil
}

func (r *Redis) Lpop() error {
	if r.GetSize() == 0 {
		return errors.New("List is Empty")
	}
	r.RedisList = r.RedisList[1:r.GetSize()]
	return nil
}

func (r *Redis) Rpop() error {
	if r.GetSize() == 0 {
		return errors.New("List is Empty")
	}
	r.RedisList = r.RedisList[0 : r.GetSize()-1]
	return nil
}
