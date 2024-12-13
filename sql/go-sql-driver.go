package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func test() {
	// 设置数据库连接字符串 (格式: 用户名:密码@tcp(地址:端口)/数据库)
	dsn := "root:password@tcp(127.0.0.1:3306)/testdb"

	// 打开数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 测试数据库连接
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("成功连接到 MySQL 数据库！")

	// 执行查询（示例：查询数据）
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 读取结果
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	// 检查是否有错误
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
