package snapsys

import "strconv"
import "strings"

// MULTI
// DECR key
// EXEC
func snapProduct(userId int64, productId int64) (success bool, message string) {
	// key: value
	// limit-user-userId     : count
	// snap-userId-productId : count
	limitCountPerProduct = 1
	limitCountPerUser = 1

	total, detail := getUserSnapCount(userId)

	if total >= limitCountPerUser {
		return false, "total snap count limit"
	}

	if v, ok := detail[productId]; ok {
		if v >= limitCountPerProduct {
			return false, "snap per product limit"
		}
	}

	return false, ""
}

func getUserSnapCount(userId int64) (total int, detail map[int64]int) {

	keys := getRedisKeys("snap-" + strconv.FormatInt(userId, 10) + "-*")

	for _, key := range keys {
		pid, _ := strconv.ParseInt(strings.Split(key, "-")[2], 10, 64)
		count := getRedisValueByKey(key).(int)

		total += count

		if _, ok := detail[pid]; ok {
			detail[pid] += count
		} else {
			detail[pid] = count
		}
	}

	return
}
