package intarr

type Slice []int32

func Encode(s Slice) ([]byte, error) {
	return is2b(s)
}
