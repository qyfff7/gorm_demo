package main

import (
	"fmt"
	"gorm_demo/config"
	"gorm_demo/repository/dao"
	"gorm_demo/repository/models"
)

func main() {
	config.InitConfig()
	db := dao.InitDB()
	user := models.User{
		UserName: "vhsj",
		Password: "123456",
	}
	db.Create(&user) // 插入数据
	user.Age = 11
	db.Save(&user)   // 修改数据
	db.Delete(&user) // 删除数据

	db.Save(&models.User{ // 插入数据
		UserName: "lisi",
		Password: "123456",
		Age:      10,
	})

	user2 := models.User{}
	db.First(&user2, 2) // 查询数据
	fmt.Println("user2的年龄是", user2.Age)
}
