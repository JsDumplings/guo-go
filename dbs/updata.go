// 更新数据
package dbs
import (
	"fmt"
)

func Updata1() {
	ret,_ := MysqlDb.Exec("UPDATE register set name=? where id=?","大郭",2)
  upd_nums,_ := ret.RowsAffected() // 更新的条数

  fmt.Println("更新的条数RowsAffected:",upd_nums)
}