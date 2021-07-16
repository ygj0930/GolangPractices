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

//查询一条记录
func (u *User) GetUserById() (*User,error) {
	//定义sql
	sqlStr := "select * from users where id = ?"

	//直接执行
	row := utils.Db.QueryRow(sqlStr,u.Id)

	//提取内容
	var id int
	var username,password,email string
	errScan := row.Scan(&id,&username,&password,&email)
	if errScan != nil {
		fmt.Println("GetUserById  Scan error:",errScan)
		return nil, errScan
	}

	//拼接结果返回
	res := &User{
		Id:id,
		UserName: username,
		PassWord: password,
		Email: email,
	}
	return res,nil
}

//查询多条记录
func (u *User) GetAllUsers() ([]*User,error){
	//定义sql
	sqlStr := "select * from users"

	//执行查询
	rows,errQuery := utils.Db.Query(sqlStr)
	if errQuery != nil {
		fmt.Println("GetAllUsers  Query error:",errQuery)
		return nil, errQuery
	}

	//遍历提取
	res := make([]*User,0)

	for rows.Next() {//游标判断是否有下一行
		var id int
		var username,password,email string
		errScan := rows.Scan(&id,&username,&password,&email)
		if errScan != nil {
			fmt.Println("GetAllUsers  Scan error:",errScan)
			return nil, errScan
		}
		rec := &User{
			Id:id,
			UserName: username,
			PassWord: password,
			Email: email,
		}
		res = append(res,rec)
	}

	//结果返回
	return res,nil
}