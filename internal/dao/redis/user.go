package redis

// 打算将生成的 atoken 放在这里，每次访问需要登录的 token 的时候，取出保存的 token，如果不一致，就要从新登录，缓存时间就是 token 过期时间
