package intarr

type Slice []int32

func New(arr interface{}) (s Slice) {
	switch arr.(type) {
	case []int:
		for _, v := range arr.([]int) {
			s = append(s, int32(v))
		}
	case []int32:
		return Slice(arr.([]int32))
	case []int64:
		for _, v := range arr.([]int64) {
			s = append(s, int32(v))
		}
	/*case []byte:
	s, e = Decode(arr.([]byte))
	return*/
	default:
		return
	}
	return
}

func (sl Slice) Encode() ([]byte, error) {
	return is2b(sl)
}

func Decode(b []byte) (Slice, error) {
	return b2is(b)
}

func (sl Slice) In(i int32) bool {
	for _, v := range sl {
		if i == v {
			return true
		}
	}
	return false
}

func (sl Slice) Int64() (res []int64) {
	for _, i := range sl {
		res = append(res, int64(i))
	}
	return
}

func (sl Slice) Remove(num int32) (res Slice) {
	for i, v := range sl {
		if v == num {
			res = append(sl[:i], sl[i+1:]...)
			return
		}
	}
	return sl
}
