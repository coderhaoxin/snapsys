package snapsys

import "encoding/json"
import "strconv"

func loadProducts(skip, limit int) error {
	// load products from postgre to redis
	// key: product-productId
	// value: json string
	products := getProducts(skip, limit)

	for _, v := range products {
		bytes, err := json.Marshal(v)

		if err != nil {
			return err
		}

		_, err = redisPool.Get().Do("SET", "product-"+strconv.FormatInt(v.Id, 10), bytes)

		if err != nil {
			return err
		}
	}

	return nil
}

func getProducts(skip, limit int) []Product {
	products := []Product{}

	if skip < 0 || limit <= 0 {
		return products
	}

	_ = psqlPool.Select(&products, "SELECT id, name, description, price, count FROM product ORDER BY id DESC OFFSET $1 LIMIT $2", skip, limit)

	return products
}
