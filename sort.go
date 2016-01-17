package intarr

import (
	"github.com/cznic/sortutil"
)

func Sort(s Slice) {
	sortutil.Int32Slice(s).Sort()
}

func (s Slice) Sort() {
	sortutil.Int32Slice(s).Sort()
}
