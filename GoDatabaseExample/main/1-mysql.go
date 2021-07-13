package main

import (
	"database/sql"
	//导入第三方包：先执行go get -t -v 第三方包路径
	//然后在本项目本目录下执行 go mod init 项目包名
	//最后，在项目的go文件中import第三方包即可
	_ "github.com/go-sql-driver/mysql"
)

var(
	Db *sql.DB
	err error
)
func main() {

	//驱动名要对应上面导入的驱动包名
	Db,err := sql.Open("mysql","root:1123456@tcp(localhost:3306)/db_test")

	//测试是否连通， 不同的话会返回err
	pingErr := Db.Ping()

	if err!= nil {
		panic(err.Error())
	}
	if pingErr != nil {
		panic(pingErr)
	}
}
