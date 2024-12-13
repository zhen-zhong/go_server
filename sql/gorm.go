package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义模型（User 表）
type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:255"`
	Age  uint
}

func test1() {
	// 1. 设置 MySQL 连接字符串
	dsn := "root:password@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	fmt.Println("成功连接到 MySQL 数据库！")

	// 2. 自动迁移（创建表）
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("迁移失败:", err)
	}
	fmt.Println("表已创建/更新！")

	// 3. 创建数据
	user := User{Name: "Alice", Age: 25}
	db.Create(&user)
	fmt.Println("插入用户:", user)

	// 4. 查询数据
	var foundUser User
	if err := db.First(&foundUser, "name = ?", "Alice").Error; err != nil {
		log.Fatal("查询失败:", err)
	}
	fmt.Println("查询到的用户:", foundUser)

	// 5. 更新数据
	db.Model(&foundUser).Update("Age", 26)
	fmt.Println("更新后的用户:", foundUser)

	// 6. 删除数据
	db.Delete(&foundUser)
	fmt.Println("用户已删除")
}
