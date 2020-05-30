package redis

import (
	"errors"
	"go-sample/pocket52/utility"

	"github.com/spf13/cast"
)

func GetStringLength(s string) int {
	return len(s)
}

func Append(s *string, data string) (length int) {
	*s += data
	return len(*s)
}

func SetNx(s *string, value string) int {

	//if key is not set, set to value
	if s == nil || *s == utility.ConstEMPTY {
		*s = value
		return 1
	}
	//else do nothing
	return 0
}

func INCR(s *string) (string, error) {

	intVal, err := cast.ToIntE(*s)

	if err != nil {
		return "Not Integer String", errors.New("Input is not integer string")
	}

	intVal = intVal + 1

	*s = cast.ToString(intVal)
	return *s, nil
}

func GetRange(s *string, st, e interface{}) (string, error) {

	var err error
	var start, end int
	start, err = cast.ToIntE(st)
	if err != nil {
		return utility.ConstEMPTY, errors.New("Non integer value for start")
	}

	end, err = cast.ToIntE(e)
	if err != nil {
		return utility.ConstEMPTY, errors.New("Non integer value for end")
	}

	if start >= 0 {

		if start > GetStringLength(*s) {
			return utility.ConstEMPTY, err
		}

		if end >= 0 {
			if start > end {
				return utility.ConstEMPTY, err
			}
			if end >= GetStringLength(*s) {
				return (*s)[start:], err
			}

			return (*s)[start : end+1], err
		}

		if end < 0 {
			if end*-1 > GetStringLength(*s) {
				return utility.ConstEMPTY, err
			}
			return (*s)[start : GetStringLength(*s)+end+1], err
		}

	}

	//if start <0

	if -1*start > GetStringLength(*s) {
		return utility.ConstEMPTY, err
	}
	if end >= 0 {
		return utility.ConstEMPTY, err
	}

	return (*s)[GetStringLength(*s)+start : GetStringLength(*s)+end+1], err

}
