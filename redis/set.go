package redis

type SetOperations interface {
	Sadd(s string) interface{}
	Scard(s string) int
}

type RedisIntSet struct {
	RedisSet []int
}

type RedisStringSet struct {
	RedisSet []string
}

func (r *RedisStringSet) Sadd(s string) interface{} {

	//check if s is alreay present in set
	for _, v := range r.RedisSet {
		if s == v {
			return r.RedisSet
		}
	}

	r.RedisSet = append(r.RedisSet, s)

	return r.RedisSet

}

func (r *RedisStringSet) Scard(s string) int {
	return len(r.RedisSet)
}
