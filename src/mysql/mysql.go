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

func GetAns(ask string) string {
	db, err := sql.Open("mysql", "root:@/mysms?charset=utf8")
	checkErr(err)

	//查询数据库
	sql := "select ans from matchtable where ask like "
	sql += "'%" + ask + "%'"
	fmt.Println("测试数据：" + sql)
	rows, err := db.Query(sql)
	checkErr(err)

	var ans string
	for rows.Next() {
		var ansTemp string
		err = rows.Scan(&ansTemp)
		ans = ansTemp
		//fmt.Println(ansTemp)
	}

	if ans == "" {
		ans = "呵呵"
	}
	//return sql
	return ans
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
