package snapsys

import . "github.com/smartystreets/goconvey/convey"
import "testing"

func TestProduct(t *testing.T) {
	Convey("get products", t, func() {
		Convey("success", func() {
			products := getProducts(0, 20)

			So(len(products), ShouldBeGreaterThan, 0)

			for _, v := range products {
				t.Log(v)

				So(len(v.Name), ShouldBeGreaterThan, 0)
				So(len(v.Desc), ShouldBeGreaterThan, 0)
				So(v.Price, ShouldBeGreaterThan, 0)
				So(v.Count, ShouldBeGreaterThan, 0)
			}
		})
	})

	Convey("load products", t, func() {
		Convey("success", func() {
			err := loadProducts(0, 20)

			So(err, ShouldBeNil)

			keys := getRedisKeys("product-*")

			So(len(keys), ShouldBeGreaterThan, 0)

			for _, v := range keys {
				So(v, ShouldStartWith, "product-")
			}
		})
	})
}
