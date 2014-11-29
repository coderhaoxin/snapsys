package snapsys

import . "github.com/smartystreets/goconvey/convey"
import "testing"

func TestSnap(t *testing.T) {
	Convey("getProduct", t, func() {
		Convey("success", func() {
			p := getProduct(1)
			So(p.Id, ShouldEqual, 1)
			So(p.Price, ShouldEqual, 100)
		})

		Convey("addSnapCount", func() {
			count, err := addSnapCount(1)

			So(count, ShouldEqual, 1)
			So(err, ShouldEqual, nil)

			count, err = addSnapCount(1)
			count, err = addSnapCount(1)
			So(count, ShouldEqual, 3)
			So(err, ShouldEqual, nil)

			count, err = addSnapCount(1)
			count, err = addSnapCount(1)
			So(count, ShouldEqual, 5)
			So(err, ShouldEqual, nil)
		})

		Convey("snapProduct", func() {
			success, message := snapProduct(1, 1)
			So(success, ShouldEqual, true)
			So(message, ShouldEqual, "success")

			success, message = snapProduct(1, 1)
			So(success, ShouldEqual, false)
			So(message, ShouldEqual, "snap per product limit")

			success, message = snapProduct(1, 2)
			So(success, ShouldEqual, true)
			So(message, ShouldEqual, "success")

			success, message = snapProduct(1, 3)
			So(success, ShouldEqual, false)
			So(message, ShouldEqual, "total snap count limit")

			for i := 10; i < 53; i++ {
				success, message = snapProduct(int64(i), 1)
				So(success, ShouldEqual, true)
				So(message, ShouldEqual, "success")
			}

			success, message = snapProduct(10086, 1)
			So(success, ShouldEqual, false)
			So(message, ShouldEqual, "product count is 0")
		})
	})
}
