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
