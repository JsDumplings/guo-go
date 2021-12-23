package main
import (
	"./dbs"
	"./https"
)

// func main() {
	// 初始化连接数据库
	// dbs.MysqlInit()
	// 插入数据
	// dbs.Insert1()
	// 查询数据
	// dbs.Search1()
	// dbs.Search2()
	// dbs.Search3()
	// 删除数据
	// dbs.Delete1()
	// 更新数据
	// dbs.Updata1()
// }


func main() {
	// 初始化连接数据库
	dbs.MysqlInit()
	https.ApiIndex()
}