package model

import(
	"databaseexcample/utils"
	"fmt"
)

type User struct {
	Id int
	UserName string
	PassWord string
	Email string
}

//用预编译执行sql语句
func (u *User) InsertByPrepare() error{
	//定义sql语句
	insertSql := "insert into users(username,password,email) values(?,?,?)"

	//进行预编译
	preStmt,errPre := utils.Db.Prepare(insertSql)
	if errPre != nil{
		fmt.Println("Sql prepare error:",errPre)
		return errPre
	}

	//通过Stmt.Exec进行插入
	_,errIns :=preStmt.Exec(u.UserName,u.PassWord,u.Email)
	if errIns != nil {
		fmt.Println("User insert error:",errIns)
		return errIns
	}

	//顺利执行，返回nil
	return nil
}

//直接用DB对象执行sql语句
func (u *User) InsertByDB() error{
	//定义sql语句
	insertSql := "insert into users(username,password,email) values(?,?,?)"

	//使用dbutils下的DB进行数据库操作
	_,errIns :=utils.Db.Exec(insertSql,u.UserName,u.PassWord,u.Email)
	if errIns != nil {
		fmt.Println("User insert error:",errIns)
		return errIns
	}

	//顺利执行，返回nil
	return nil
}