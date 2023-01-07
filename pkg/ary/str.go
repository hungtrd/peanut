package ary

type T interface {
	int64 | string | int
}

func InArray[V T](val V, array []V) bool {
	for _, v := range array {
		if val == v {
			return true
		}
	}
	return false
}
