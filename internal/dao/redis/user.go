package redis

import "strconv"

// 打算将生成的 atoken 放在这里，每次访问需要登录的 token 的时候，取出保存的 token，如果不一致，就要从新登录，缓存时间就是 token 过期时

func SetToken(userId uint64, aToken string) (err error) {
	// 先简单实现,后续设置缓存时间等
	err = rdb.Set(strconv.FormatUint(userId, 10), aToken, 0).Err()
	return
}

func GetToken(userId uint64) (redisToken string, err error) {
	// 先简单实现,后续设置缓存时间等
	redisToken, err = rdb.Get(strconv.FormatUint(userId, 10)).Result()
	return
}
