package utils

import (
	"database/sql"
	"fmt"

	//导入第三方包：先执行go get -t -v 第三方包路径
	//然后在本项目本目录下执行 go mod init 项目包名
	//最后，在项目的go文件中import第三方包即可
	_ "github.com/go-sql-driver/mysql"
)
var (
	Db *sql.DB
	err error
)

func init() {
	//驱动名要对应上面导入的驱动包名
	//Open函数：不建立数据库连接，只验证参数合法性
	//数据源参数：用户名:密码@tcp(ip:端口)/数据库名
	Db,err = sql.Open("mysql","root:123456@tcp(localhost:3306)/db_test")
	//测试数据库连接参数是否合法
	if err!= nil {
		panic(err.Error())
	}
	//测试是否连通，不通的话会返回err
	pingErr := Db.Ping()
	if pingErr != nil {
		fmt.Println("数据库Ping不通...")
		panic(pingErr)
	}
}