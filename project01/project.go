package main

import (
	"awesomeProject/ConnectSql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"log"
	"net/http"
	_ "net/http"
)

//func dealMap(m map[string]string) map[string]string{
//	for k, v := range m {
//		if k=="sex" {
//			if v=="男" {
//				fmt.Println("nan")
//				m[k] = string(1)
//			} else {
//				m[k] = string(0)
//			}
//		}
//	}
//	return m
//}


func main() {
	r := gin.Default()
	//{"title":"1","email":"1193985703@qq.com","type":"1",
	//"like[write]":"on","like[read]":"on","like[dai]":"on","star":"on","sex":"男","desc":"123"}
	r.POST("/test", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")

		buf := make([]byte, 1024)
		n, _ := c.Request.Body.Read(buf)
		data := string(buf[0:n])
		fmt.Println(string(buf[0:n]))
		value := gjson.Get(data,"title")
		fmt.Printf("value: %s\n",value)
		json := make(map[string]interface{})
		fmt.Printf("%v\n",json)
		err := c.BindJSON(&json)
		if err==nil {
			log.Println(err)
			c.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
			return
		}
		fmt.Printf("%v\n",json)
		c.JSON(http.StatusOK,gin.H{
			"title":json["title"],
			"email":json["email"],
			"typex":json["typex"],
			"likew":json["like[write]"],
			"liker":json["like[read]"],
			"liked":json["like[dai]"],
			"star":json["star"],
			"sex":json["sex"],
			"desc":json["desc"],
		})
	})

	r.GET("/getMessage", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		ConnectSql.SqlOpen()
		res := ConnectSql.SqlSelect()
		//fmt.Printf("%v\n",res)
		ConnectSql.SqlClose()
		mjson, _ := json.Marshal(res)
		mString :=string(mjson)
		fmt.Println(mString)
		c.String(http.StatusOK,mString)
	})

	r.POST("/commit", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		//{"desc":null,"email":null,"liked":null,"liker":null,"likew":null,"sex":null,"star":null,"title":null,"typex":null}
		buf := make([]byte, 1024)
		n, _ := c.Request.Body.Read(buf)
		data := string(buf[0:n])
		fmt.Println(string(buf[0:n]))
		title := gjson.Get(data,"title")
		email := gjson.Get(data,"email")
		typex := gjson.Get(data,"typex")
		likew := gjson.Get(data,"like[write]")
		//fmt.Printf("likew: %s\n",likew)
		liker := gjson.Get(data,"like[read]")
		//fmt.Printf("likew: %s\n",liker)
		sex := gjson.Get(data,"sex")
		liked := gjson.Get(data,"like[dai]")
		//fmt.Printf("likew: %s\n",liked)
		star := gjson.Get(data,"star")
		desc := gjson.Get(data,"desc")

		m := make(map[string]string)
		if err := json.Unmarshal(buf[0:n],&m); err != nil {
			fmt.Println("Umarshal failed:", err)
			return
		}
		//业务操作：
		//ConnectSql.SqlInsert(m)

		fmt.Printf("%v\n",m)

		c.JSON(http.StatusOK,gin.H{
			"title":title.Str,
			"email":email.Str,
			"typex":typex.Str,
			"likew":likew.Str,
			"liker":liker.Str,
			"liked":liked.Str,
			"star":star.Str,
			"sex":sex.Str,
			"desc":desc.Str,
		})
	})

	r.POST("/testCommit", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		//{"desc":null,"email":null,"liked":null,"liker":null,"likew":null,"sex":null,"star":null,"title":null,"typex":null}
		buf := make([]byte, 1024)
		n, _ := c.Request.Body.Read(buf)
		data := string(buf[0:n])
		fmt.Println(data)

		m := make(map[string]string)
		if err := json.Unmarshal(buf[0:n],&m); err != nil {
			fmt.Println("Umarshal failed:", err)
			return
		}
		//fmt.Printf("%v\n",m)
		c.String(http.StatusOK,"good!")
	})

	r.GET("/get", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		ConnectSql.SqlOpen()
		res := ConnectSql.SelectComments()
		fmt.Printf("%v\n",res)
		//lens := ConnectSql.SelectLen()

		ConnectSql.SqlClose()
		mjson, err := json.Marshal(res)
		if err!=nil {
			panic(err)
		}
		mString :=string(mjson)
		c.String(http.StatusOK, mString)
	})

	r.POST("/insert", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		//json参数处理：
		buf := make([]byte, 1024)
		n, _ := c.Request.Body.Read(buf)
		data := string(buf[0:n])
		m := make(map[string]string)
		if err := json.Unmarshal(buf[0:n],&m); err != nil {
			fmt.Println("Umarshal failed:", err)
			return
		}
		fmt.Println(data)
		//map特殊数据处理：
		//m = dealMap(m)
		for k, v := range m {
			fmt.Printf("key: %s, value: %s\n",k,v)
		}

		ConnectSql.SqlOpen()
		res := ConnectSql.SqlInsert(m)
		defer ConnectSql.SqlClose()
		if res>=1 {
			c.JSON(http.StatusOK,gin.H{
				"status": "ok",
			})
			return
		}
		c.JSON(http.StatusBadRequest,gin.H{
			"status": "fail",
		})
	})

	r.GET("/getlens", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		ConnectSql.SqlOpen()
		lens := ConnectSql.SelectLen()
		fmt.Println(lens)
		defer ConnectSql.SqlClose()
		c.JSON(http.StatusOK, gin.H{
			"lens": lens,
		})
	})
	r.Run()
}
