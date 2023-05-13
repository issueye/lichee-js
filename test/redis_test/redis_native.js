let native_data = {
    name: 'test:001',
    code: 200,
}

let native_key = "redis-js:test:string:001"

// 字符串操作
licheeRedis.set(native_key, JSON.stringify(native_data), 10)

let native_rtData = licheeRedis.get(native_key)

console.log("redis return data", native_rtData.value);

let native_listKey = "redis-js:test:list:001"
licheeRedis.lPush(native_listKey, JSON.stringify(native_data))