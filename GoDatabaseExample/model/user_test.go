package model

import (
	"fmt"
	"testing"
)


func TestUser_GetUserById(t *testing.T) {
	fmt.Println("测试查询单条用户")
	userCondition := &User{
		Id:5,
	}
	res,err := userCondition.GetUserById()
	if err != nil{
		t.Error(err)
	}else{
		fmt.Println("查询结果为：",res)
	}

}

func TestUser_GetAllUsers(t *testing.T) {
	fmt.Println("测试查询所有用户")
	userCondition := &User{}
	res,err := userCondition.GetAllUsers()
	if err != nil{
		t.Error(err)
	}else{
		fmt.Println("查询结果为：")
		for _,val := range res {
			fmt.Println(val)
		}
	}
}

func TestUser_InsertByDB(t *testing.T) {
	fmt.Println("测试直接通过DB添加用户")
	user1 := &User{
		UserName: "user4",
		PassWord: "1234",
		Email: "user1@1234.com",
	}
	err := user1.InsertByDB()
	if err != nil{
		t.Error(err)
	}
}

func TestUser_InsertByPrepare(t *testing.T) {
	fmt.Println("测试通过PrepareStatement添加用户")
	user2 := &User{
		UserName: "user5",
		PassWord: "1234",
		Email: "user1@1234.com",
	}
	err := user2.InsertByPrepare()
	if err != nil{
		t.Error(err)
	}
}

func BenchmarkUser_InsertByDB(b *testing.B) {
	fmt.Println("测试DB插入性能")
	user1 := &User{
		UserName: "user9",
		PassWord: "1234",
		Email: "user1@1234.com",
	}
	err := user1.InsertByDB()
	if err != nil{
		b.Error(err)
	}
}

func BenchmarkUser_InsertByPrepare(b *testing.B) {
	fmt.Println("测试PrepareStatement插入性能")
	user2 := &User{
		UserName: "user10",
		PassWord: "1234",
		Email: "user1@1234.com",
	}
	err := user2.InsertByPrepare()
	if err != nil{
		b.Error(err)
	}
}