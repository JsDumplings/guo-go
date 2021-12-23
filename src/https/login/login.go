package login
import (
	"encoding/json"
	"fmt"
	"net/http"
	"../returnJson"
	"log"
	"io/ioutil"
	"../../dbs"
)
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型

	type LoginReq struct {
		Username string `json:"username"`
		Password string    `json:"password"`
	}
	body ,err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println( err)
	}
	log.Printf("%s",body)
	
	var data LoginReq

	json.Unmarshal([]byte(body),&data)
	result := returnJson.NewBaseJsonBean()
  fmt.Println(result)

	// 连接数据库查询 用户名与密码是否正确
	username,password := dbs.Search3(data.Username)
	if len(username) != 0 && password == data.Password { 
		result.Code = 200 
		result.Message = "登录成功" 
	} else { 
		result.Code = 101
		result.Message = "用户名或密码不正确" 
	}

	//向客户端返回JSON数据 
	bytes, _ := json.Marshal(result) 
	fmt.Fprint(w, string(bytes))
}
func LoginApi(){
	http.HandleFunc("/", IndexHandler)
	err := http.ListenAndServe("127.0.0.1:9000", nil)
	if(err != nil) {
		fmt.Println(err)
	}
}