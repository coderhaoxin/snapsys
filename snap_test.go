package snapsys

import . "github.com/smartystreets/goconvey/convey"
import "testing"

func TestSnap(t *testing.T) {
	Convey("pass t", t, func() {
		x := 1

		Convey("test", func() {
			x++
			So(x, ShouldEqual, 2)
		})
	})
}
