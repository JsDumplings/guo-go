package dbs
import (
	"fmt"
)

// 注册表结构体
type Register struct {
	name string `db:"name"`
	password string
	phone int
}
// 查询数据，取所有字段
func Search1() {
	// 通过切片储存
	res := make([]Register,0)

	rows, _:= MysqlDb.Query("select * from register")
	var reg Register
	// 遍历
  for rows.Next(){        //循环显示所有的数据
      rows.Scan(&reg.name,&reg.password,&reg.phone)
			res = append(res,reg)
  }
	fmt.Println(res)
}

// 查询多条数据
func Search2() {
	rows, _:= MysqlDb.Query("select name,password from register")
	var name,password string
	for rows.Next(){        //循环显示所有的数据
		rows.Scan(&name,&password)
		fmt.Println(name,password)
	}
	// return name,password
}

// 查询单行 使用QueryRow
func Search3(username string)(name string, password string)  {
	rows := MysqlDb.QueryRow("select name,password from register where name=?",username)
	rows.Scan(&name,&password)
	fmt.Println(name,password)
	return name,password
}