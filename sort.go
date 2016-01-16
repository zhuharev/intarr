package intarr

import (
	"github.com/cznic/sortutil"
)

func Sort(s Slice) {
	sortutil.Int32Slice(s).Sort()
}
