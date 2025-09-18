package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//db, _ := sql.Open("mysql", "root:q123q123@(127.0.0.1:3306)/test")
	db, _ := sql.Open("mysql", "root:q123q123@(127.0.0.1:3306)/test") // 设置连接数据库的参数

	drivers := sql.Drivers()
	strings := drivers
	fmt.Println(strings)
	defer db.Close()
	err := db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败")
	} else {
		fmt.Println("数据库连接成功")
	}

	//sql := "insert into stuinfo values (2, 'cat')"
	//result, err := db.Exec(sql)
	//n, err := result.RowsAffected()
	//fmt.Println("受影响的记录数是：", n)

	//stu := [2][2]string{{"3", "ketty"}, {"4", "rose"}}
	//stmt, err := db.Prepare("insert into stuinfo values(?, ?)")
	//for _, item := range stu {
	//	stmt.Exec(item[0], item[1])
	//}

	var id, name string
	row := db.QueryRow("select * from stuinfo")
	row.Scan(&id, &name)
	fmt.Println(id, "--", name)

	rows, err := db.Query("select * from stuinfo")
	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Println(id, "--", name)
	}
}
