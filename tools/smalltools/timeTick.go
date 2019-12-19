package smalltools

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//日期转化为时间戳  https://www.cnblogs.com/-xuzhankun/p/10812048.html
//注意事项  https://blog.csdn.net/harleylau/article/details/82900213   cannot parse "'2019-12-21" as "2"
var timeLayout string = "2006-1-02 15:04:05" //转化所需模板
func timeToTick(datetime string) {
	st, err := time.Parse(timeLayout, datetime) //string转time
	if err != nil {
		fmt.Println(err)
	} else {
		WinColorPrintln(st.Unix(), 2) //转化为时间戳 类型是int64
	}
}

func tickToTime(timestamp int) {
	//日期转化为时间戳
	// loc, _ := time.LoadLocation("Local")    //获取时区
	// tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)
	// timestamp := tmp.Unix()    //转化为时间戳 类型是int64
	// fmt.Println(timestamp)
	//时间戳转化为日期
	datetime := time.Unix(int64(timestamp), 0).Format(timeLayout)
	WinColorPrintln(datetime, 2)
}

func Main() {
	for {
		WinColorPrintln("\nquit | timeToTick 2019-12-19 15:04:05 | tickToTime 1576719376", 1)
		fmt.Println("wait input...")
		choice := ""
		data := ""
		dataTime := ""
		fmt.Scanln(&choice, &data, &dataTime)
		switch strings.ToLower(choice) {
		case "quit":
			return
		case "timetotick":
			timeToTick(data + " " + dataTime)
		case "ticktotime":
			dt, _ := strconv.Atoi(data)
			tickToTime(dt)
		}
	}
}
