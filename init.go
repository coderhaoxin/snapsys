package snapsys

import "github.com/garyburd/redigo/redis"
import "github.com/jmoiron/sqlx"
import _ "github.com/lib/pq"
import "time"
import "os"

var redisPool *redis.Pool
var psqlPool *sqlx.DB

var limitCountPerProduct int
var limitCountPerUser int

// all the init of package: snapsys here
func init() {
	var err error

	// limit
	limitCountPerProduct = 1
	limitCountPerUser = 2

	// redis
	redisPool = &redis.Pool{
		MaxIdle:     30,
		MaxActive:   300,
		IdleTimeout: 1 * time.Second,
		Dial: func() (conn redis.Conn, err error) {
			conn, err = redis.Dial("tcp", "localhost:6379")

			if err != nil {
				panic(err)
			}

			return
		},
	}

	// postgre
	psqlPool, err = sqlx.Connect("postgres", "user=hx dbname=hx sslmode=disable")
	panicError(err)

	// empty redis for test or development
	if os.Getenv("GOLANG_ENV") != "production" {
		_, err = redisPool.Get().Do("FLUSHDB")
		panicError(err)
	}

	// load products
	err = loadProducts(0, 20)
	panicError(err)
}
