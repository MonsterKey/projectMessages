//+build ignore

package UserSql

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
)
var db *sql.DB
func sqlOpen() {
	var err error
	db, err = sql.Open("postgres", "port=5432 user=postgres password=1234 dbname=message sslmode=disable")
	//port是数据库的端口号，默认是5432，如果改了，这里一定要自定义；
	//user就是你数据库的登录帐号;
	//dbname就是你在数据库里面建立的数据库的名字;
	//sslmode就是安全验证模式;

	//还可以是这种方式打开
	//db, err := sql.Open("postgres", "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full")
	checkErr(err)
}


func sqlSelect() {
	//查询数据
	rows, err := db.Query("SELECT * FROM comment_list")
	checkErr(err)

	println("-----------")
	for rows.Next() {
		var id int
		var name string
		var email string
		var comment string
		var comment_date string
		err = rows.Scan(&id, &name, &email, &comment, &comment_date)
		checkErr(err)
		fmt.Println("id = ", id, "\nname = ", name, "\nemail= ",email,"\ncomment = ", comment, "\ncreated = ", comment_date, "\n-----------")
	}
}


func sqlClose() {
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}


func main() {
	sqlOpen()
	sqlSelect()
	sqlClose()
}
