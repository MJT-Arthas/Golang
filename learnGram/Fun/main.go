package main

import (
	"fmt"
	"strings"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin(users, coins)
	fmt.Println("剩下：", left)
	fmt.Println(distribution)
}

func dispatchCoin(arr []string, coins int) int {
	sum := 0
	for _, name := range arr {
		distribution[name] = 0
		for _, s := range name {

			switch strings.ToLower(string(s)) {
			case "e":
				sum += 1
				distribution[name] += 1
			case "i":
				sum += 2
				distribution[name] += 2
			case "o":
				sum += 3
				distribution[name] += 3
			case "u":
				sum += 4
				distribution[name] += 4
			default:

			}

		}
	}
	return coins - sum
}
