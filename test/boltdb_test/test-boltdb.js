var boltdb = require("db/bolt")


boltdb.open("test.db")

boltdb.createBucket("boltjs")

boltdb.update("boltjs", function(db) {
    db.put("test:0001", "这是一条测试数据")
})

boltdb.view("boltjs", function(db) {
    let data = db.get("test:0001")
    console.log(data.value);
})
