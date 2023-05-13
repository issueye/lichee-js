var redisDB = require('db/redis')

let data = {
    nativeName: 'test:001',
    code: 200,
}

let key = "redis-js:test:string:001"

let redis = redisDB.newClient("49.235.124.25:6379", "123456", 0)

// 字符串操作
redis.set(key, JSON.stringify(data), 10)

let rtData = redis.get(key)

console.log("redis return data", rtData.value);

let listKey = "redis-js:test:list:001"
redis.lPush(listKey, JSON.stringify(data))

let listData = redis.lRange(listKey, 0, -1)
console.log('list_data', JSON.stringify(listData));