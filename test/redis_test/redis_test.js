var redis = require("db/redis")


let data = {
    name: 'test:001',
    code: 200,
}

let key = "redis-js:test:string:001"

// 字符串操作
redis.set(key, JSON.stringify(data), 10)

let rtData = redis.get(key)

console.log("redis return data", rtData.value);

let listKey = "redis-js:test:list:001"
redis.lPush(listKey, JSON.stringify(data))