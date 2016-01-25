package intarr

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestEncodeDecode(t *testing.T) {

	var (
		s   = Slice{1, 2, 3, 4, 5}
		bts []byte
		e   error
	)

	Convey("test encode", t, func() {
		bts, e = is2b(s)
		So(e, ShouldBeNil)
		So(len(bts), ShouldEqual, 20)
	})

	Convey("test decode", t, func() {
		sl, e := b2is(bts)
		So(e, ShouldBeNil)
		So(len(sl), ShouldEqual, 5)
		for k, v := range sl {
			So(v, ShouldEqual, s[k])
		}
	})

	Convey("test gzip", t, func() {
		bts, e := PackGzip(s)
		So(e, ShouldBeNil)
		So(len(bts), ShouldBeGreaterThan, 20)

		sl, e := UnpackGzip(bts)
		So(e, ShouldBeNil)
		for i, v := range sl {
			So(v, ShouldEqual, s[i])
		}

	})
}
