package main

import (
	"fmt"
	"net/http"
	"sync"
)

type Dimension struct {
	Yesterday     []string `json:"yesterday"`
	FiveDod       []string `json:"fiveDod"`
	XAxis         []string `json:"xAxis"`
	Today         []string `json:"today"`
	TodayFive     []string `json:"todayFive"`
	Weekday       []string `json:"weekday"`
	Dod           []string `json:"dod"`
	YesterdayFive []string `json:"yesterdayFive"`
	FiveWow       []string `json:"fiveWow"`
	WeekdayFive   []string `json:"weekdayFive"`
	Wow           []string `json:"wow"`
}

type DimensionForm struct {
	InternationTPC Dimension `json:"0-新国际酒店TPC整体流量"`
}

type Actress struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    DimensionForm `json:"data"`
}

var wg sync.WaitGroup

func main() { // 普通JSON
	for i := 0; i < 100; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go fetch()
	}
	wg.Wait() // 等待所有登记的goroutine都结束

}

func fetch() {
	defer wg.Done()
	url := "https://www.elong.com"
	_, err := http.Get(url)
	if err == nil {
		fmt.Printf("获取%s成功\n", url)

	} else {

		fmt.Println(err)
	}

	//
	//url := "http://tcservice.17usoft.com/flowrules/boardInfo/result?dimension_1=" + "" + "&mainId=" + "113" + "&mainType=0&userId=1206970"
	//resp, err := http.Get(url)
	//if err == nil {
	//	fmt.Printf("获取%s成功\n", url)
	//	body, err := ioutil.ReadAll(resp.Body)
	//
	//	var actress Actress
	//
	//	err = json.Unmarshal(body, &actress)
	//	if err != nil {
	//		fmt.Println("error:", err)
	//		return
	//	}
	//
	//	fmt.Printf("Code：%v\n", actress.Code)
	//
	//	fmt.Printf("Message：%s\n", actress.Message)
	//	fmt.Printf("Data：%v\n", actress.Data.InternationTPC)
	//}
}
