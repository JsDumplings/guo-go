package dbs
import (
	"fmt"
)
// 删除数据
func Delete1() {
	ret,_ := MysqlDb.Exec("delete from register where id=?",1)
	del_nums,_ := ret.RowsAffected() // 删除的条数
	fmt.Println("删除的条数RowsAffected:",del_nums)
}