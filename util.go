package snapsys

func getRedisKeys(pattern string) []string {
	var result []string

	reply, err := redisPool.Get().Do("KEYS", pattern)

	if err != nil {
		return result
	}

	keys, ok := reply.([]interface{})

	if !ok {
		return result
	}

	for _, v := range keys {
		bytes, ok := v.([]byte)

		if ok {
			result = append(result, string(bytes))
		}
	}

	return result
}

func getRedisValueByKey(key interface{}) interface{} {
	reply, err := redisPool.Get().Do("GET", key)

	if err != nil {
		return nil
	}

	return reply
}
