package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type LineInfo struct {
	LineName string     `json:"line_name"`
	Time     []TimeInfo `json:"time"`
}
type TimeInfo struct {
	Direction interface{} `json:"direction"`
	First     interface{} `json:"first"`
	Second    interface{} `json:"second"`
}

type Responses struct {
	Code        string     `json:"code"`
	Message     string     `json:"message"`
	Data        []LineInfo `json:"data"`
	StationName string
}

var line2station = [...]string{
	"骑河",
	"富翔路",
	"高铁苏州北站",
	"大湾",
	"富元路",
	"蠡口",
	"徐图港",
	"阳澄湖中路",
	"陆慕",
	"平泷路东",
	"平河路",
	"苏州火车站",
	"山塘街",
	"石路",
	"广济南路",
	"三香广场",
	"劳动路",
	"胥江路",
	"桐泾公园",
	"友联",
	"盘蠡路",
	"新家桥",
	"石湖东路",
	"宝带桥南",
	"尹中路",
	"郭巷",
	"郭苑路",
	"尹山湖",
	"独墅湖南",
	"独墅湖邻里中心",
	"月亮湾",
	"松涛街",
	"金谷路",
	"金尚路",
	"桑田岛",
}
var wg sync.WaitGroup

func main() {
	resultChan := make(chan Responses, 10)

	for _, stationName := range line2station {
		wg.Add(1) // 启动一个goroutine就登记+1
		go fetch(stationName, resultChan)
	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for res := range resultChan {

		fmt.Println(res.StationName, res.Data)
	}
	fmt.Println("完成")
}

func fetch(currentStation string, resultChan chan Responses) {
	defer wg.Done()

	url := "https://szgj.2500.tv/api/v1/search/schedule?current_station=" + currentStation
	resp, err := http.Get(url)
	if err == nil {

		body, err := io.ReadAll(resp.Body)

		var res Responses

		err = json.Unmarshal(body, &res)
		if err != nil {
			fmt.Println("error:", err, currentStation)

			return
		}
		res.StationName = currentStation
		resultChan <- res
		//fmt.Printf("Code：%v\n", res.Code)
		//fmt.Printf("Message：%s\n", res.Message)
		//fmt.Printf("Data：%v\n", res.Data)
	}
}
