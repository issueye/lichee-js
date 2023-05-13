var boltdb = require("db/bolt")


var db = boltdb.open("test.db")

db.createBucket("boltjs")

db.update("boltjs", function (db) {
    db.put("test:0001", "这是一条测试数据")
})

db.view("boltjs", function (db) {
    let data = db.get("test:0001")
    console.log(data.value);
})
