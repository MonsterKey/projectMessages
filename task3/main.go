package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

func main() {
	//发送get请求
	resp, err := http.Get("https://api.help.bj.cn/apis/weather/?id=101060101 ")
	if err != nil {
		fmt.Println("get failed, err", err)
	}
	//关闭Body
	defer resp.Body.Close()
	//读取body内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read from resp.Body failed, err", err)
		return
	}
	//输出字符串内容
	fmt.Println(string(body))
	json := string(body)
	//    "status": "0",
	//    "cityen": "changchun",
	//    "city": "长春",
	//    "citycode": "101060101",
	//    "temp": "21",
	//    "tempf": "69",
	//    "wd": "南风",
	//    "wden": "S",
	//    "wdenimg": "//www.help.bj.cn/weather/images/wind/s.png",
	//    "wdforce": "2级",
	//    "wdspd": "6km/h",
	//    "humidity": "94%",
	//    "stp": "980",
	//    "wisib": "6km",
	//    "uptime": "17:55",
	//    "prcp": "0",
	//    "prcp24h": "0",
	//    "aqi": "43",
	//    "pm25": "43",
	//    "weather": "多云",
	//    "weatheren": "Cloudy",
	//    "weathercode": "d01",
	//    "weatherimg": "//www.help.bj.cn/weather/images/d01.png",
	//    "today": "2021-08-20 周五"
	value01 := gjson.Get(json, "city")
	value02 := gjson.Get(json, "weather")
	value03 := gjson.Get(json, "today")
	value04 := gjson.Get(json, "wdforce")
	value05 := gjson.Get(json, "wd")
	value06 := gjson.Get(json, "pm25")
	value07 := gjson.Get(json, "temp")
	value08 := gjson.Get(json,"humidity")
	fmt.Println("城市："+value01.String())
	fmt.Println("天气："+value02.String())
	fmt.Println("日期："+value03.String())
	fmt.Println("风向："+value05.String())
	fmt.Println("风级：："+value04.String())
	fmt.Println("PM2.5："+value06.String())
	fmt.Println("温度："+value07.String()+"℃")
	fmt.Println("湿度："+value08.String())


}
