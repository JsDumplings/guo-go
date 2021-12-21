package dbs
import (
	"fmt"
)

func Insert1(){
	// 插入数据
	ret,_ := MysqlDb.Exec("insert INTO register(name,password,phone) values(?,?,?)","小红","12345",123)

  //插入数据的主键id
  lastInsertID,_ := ret.LastInsertId()
  fmt.Println("插入数据的主键:",lastInsertID)

  //影响行数
  rowsaffected,_ := ret.RowsAffected()
  fmt.Println("影响行数:",rowsaffected)

}