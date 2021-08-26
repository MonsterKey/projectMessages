package ConnectSql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)
var db *sql.DB

type Comments struct {
	title string
	email string
	typex int
	likew string
	liker string
	liked string
	star string
	sex string
	desc string
	userId int
	//date string
}

func SqlOpen() {
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

func SqlSelect() map[string]string{
	//查询数据
	rows, err := db.Query("SELECT * FROM comment_list")
	checkErr(err)

	println("-----------")
	list := make(map[string]string)
	for rows.Next() {
		var id int
		var name string
		var email string
		var comment string
		err = rows.Scan(&id, &name, &email, &comment)
		checkErr(err)
		//list["id"] = string(id)
		list["name"] = name
		list["email"] = email
		list["comment"] = comment
		fmt.Println("id = ", id, "\nname = ", name, "\nemail= ",email,"\ncomment = ", comment, "\n-----------")
	}
	return list
}

func SelectComments() []map[string]interface{} {
	//查询数据
	rows, err := db.Query("SELECT * FROM comments")
	checkErr(err)

	println("-----------")
	//list := make(map[string]interface{})
	//lens := SelectLen()
	lists := make([]map[string]interface{},0)
	var p Comments
	for rows.Next() {
		err = rows.Scan(&p.title, &p.email, &p.typex, &p.likew, &p.liker, &p.liked, &p.star, &p.sex, &p.desc, &p.userId)
		//err = rows.Scan(&id, &title, &email, &typex, &likew, &liker, &liked, &star, &sex, &desc, &userId, &data)
		checkErr(err)
		//list["id"] = string(p.id)
		//list["email"] = p.email
		//list["title"] = p.title
		//list["likew"] = string(p.likew)
		//list["liker"] = string(p.liker)
		//list["liked"] = string(p.liked)
		//list["star"] = string(p.star)
		//list["sex"] = string(p.sex)
		//list["desc"] = p.desc
		//list["user_id"] = string(p.userId)

		list := make(map[string]interface{})
		list["email"] = p.email
		list["title"] = p.title
		list["likew"] = p.likew
		list["liker"] = p.liker
		list["liked"] = p.liked
		list["star"] = p.star
		list["sex"] = p.sex
		list["desc"] = p.desc
		list["user_id"] = string(p.userId)
		//list["date"] = p.date
		//fmt.Printf("list: %v\n",list)
		lists = append(lists, list)
	}
	return lists
}

func SqlInsert(m map[string]string) int64 {
	//插入数据
	//INSERT INTO comments values (2,'2','222',2,1,1,1,1,1,'hhh',1);
	//stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid")
	stmt, err := db.Prepare("INSERT INTO comments(title,email,typex,likew,liker,liked,star,sex,\"desc\",user_id) " +
		"values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)")
	checkErr(err)

	for k, v := range m {
		fmt.Printf("key: %s, value: %s\n",k,v)
	}
	res, err := stmt.Exec(
		m["title"],
		m["email"],
		m["typex"],
		m["like[write]"],
		m["like[read]"],
		m["like[dai]"],
		m["star"],
		m["sex"],
		m["desc"],
		1,
		)
	//这里的三个参数就是对应上面的$1,$2,$3了
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("rows affect:", affect)
	return affect
}

func SelectLen()  int{
	//查询数据
	rows, err := db.Query("SELECT \"count\"(user_id) FROM comments")
	checkErr(err)
	var lens int
	for rows.Next() {
		err = rows.Scan(&lens)
		checkErr(err)
	}
	return lens
}

func Test()  {
	//	sqlStr := "INSERT INTO doorindex_username(id,username,jurisdiction) VALUE(?,?,?)"
	//插入数据
	sql := "insert into comment_list(id,name,email,comment) values($1,$2,$3,$4)"
	//sql := "INSERT into comment_list VALUES(3,'zxj','3421','hhhh')"
	stmt, err := db.Prepare(sql)
	if err!=nil {
		fmt.Println(err)
	}
	checkErr(err)

	res, err := stmt.Exec(7, "软件开发部门","1121321","hhhhh")
	//res, err := stmt.Exec()

	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("rows affect:", affect)
}

func SqlClose() {
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}


//func main() {
//	sqlOpen()
//	sqlSelect()
//	sqlClose()
//}
