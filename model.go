package snapsys

type User struct{}

type Snap struct {
	UserId     int64
	ProductId  int64
	CreateTime int64
}

type Order struct{}

type Record struct{}

type Product struct {
	Id    int64  `db:"id"`
	Name  string `db:"name"`
	Desc  string `db:"description"`
	Price int32  `db:"price"`
	Count int32  `db:"count"`
}
