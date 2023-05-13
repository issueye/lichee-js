

testBolt.createBucket("boltjs")

testBolt.update("boltjs", function (db) {
    db.put("test:0001", "这是一条测试数据")
})

testBolt.view("boltjs", function (db) {
    let data = db.get("test:0001")
    console.log(data.value);
})
