package snapsys

import "encoding/json"
import "strconv"
import "strings"
import "time"

func snapProduct(userId int64, productId int64) (success bool, message string) {
	// key: value
	// snap-userId-productId        : count
	// snap-detail-userId-productId : snap detail

	total, detail := getUserSnapCount(userId)

	if total >= limitCountPerUser {
		return false, "total snap count limit"
	}

	if v, ok := detail[productId]; ok {
		if v >= limitCountPerProduct {
			return false, "snap per product limit"
		}
	}

	product := getProduct(productId)
	count, err := addSnapCount(productId)

	// warning: for now, the work flow has bugs

	if count >= int64(product.Count) || err != nil {
		return false, "product count is 0"
	}

	err = addSnap(userId, productId)

	if err != nil {
		return false, "add snap error"
	}

	return true, "success"
}

func getProduct(productId int64) (product *Product) {
	data := getRedisValueByKey("product-" + strconv.FormatInt(productId, 10))

	if data == nil {
		debug("get product data: nil")
		return new(Product)
	}

	product = new(Product)
	err := json.Unmarshal(data.([]byte), product)

	if err != nil {
		debug("parse product detail error: %s", err.Error())
		return new(Product)
	}

	return
}

func getUserSnapCount(userId int64) (total int, detail map[int64]int) {
	keys := getRedisKeys("snap-" + strconv.FormatInt(userId, 10) + "*")

	detail = make(map[int64]int)

	for _, key := range keys {
		debug("get value from key: %s", key)
		pid, _ := strconv.ParseInt(strings.Split(key, "-")[2], 10, 64)
		count, _ := strconv.Atoi(string(getRedisValueByKey(key).([]byte)[:]))

		total += count

		if _, ok := detail[pid]; ok {
			detail[pid] += count
		} else {
			detail[pid] = count
		}
	}

	return
}

func addSnapCount(productId int64) (count int64, err error) {
	// MULTI
	// DECR key
	// EXEC
	reply, err := redisPool.Get().Do("INCR", "product-"+strconv.FormatInt(productId, 10)+"-snap")

	if err != nil {
		debug("add snap count error: %s", err.Error())
		return 0, err
	}

	return reply.(int64), nil
}

func addSnap(userId int64, productId int64) (err error) {
	snap := Snap{
		UserId:     userId,
		ProductId:  productId,
		CreateTime: time.Now().UnixNano(),
	}

	err = setRedisKeyValue("snap-"+strconv.FormatInt(userId, 10)+"-"+strconv.FormatInt(productId, 10), 1)
	if err != nil {
		return
	}

	err = setRedisKeyValue("snap-detail-"+strconv.FormatInt(userId, 10)+"-"+strconv.FormatInt(productId, 10), snap)
	if err != nil {
		return
	}

	return nil
}
