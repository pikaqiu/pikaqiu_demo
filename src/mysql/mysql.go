package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
)

func TestMysql() {
	db, err := sql.Open("mysql", "root:@/mysms?charset=utf8")
	checkErr(err)
	fmt.Println("ok")

	//插入数据： 测试成功
	stmt, err := db.Prepare("INSERT INTO user VALUES(?, ?, ?)")
	checkErr(err)

	_, err = stmt.Exec(2, "mc", "222222")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}