package snapsys

import . "github.com/smartystreets/goconvey/convey"
import "testing"

func TestUtil(t *testing.T) {
	Convey("util", t, func() {
		Convey("getRedisValueByKey", func() {
			So(getRedisValueByKey("not-exsit"), ShouldEqual, nil)
		})

		Convey("setRedisKeyValue", func() {
			So(setRedisKeyValue("test", "result"), ShouldEqual, nil)
			So(string(getRedisValueByKey("test").([]byte)[:]), ShouldEqual, "result")
		})
	})
}
