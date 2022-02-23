package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// db 是一个 sql.DB 类型的对象
	// 该对象线程安全，且内部已包含了一个连接池
	// 连接池的选项可以在 sql.DB 的方法中设置，这里为了简单省略了
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/ry_vue")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var (
		id int
		name string
	)
	rows, err := db.Query("select * from sys_role_menu")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	// 必须要把 rows 里的内容读完，或者显式调用 Close() 方法，
	// 否则在 defer 的 rows.Close() 执行之前，连接永远不会释放
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}