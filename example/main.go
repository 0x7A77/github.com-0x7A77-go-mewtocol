package main

import (
	"fmt"

	"github.com/0x7A77/go-mewtocol"
)

// 35.10
func main() {
	c, err := mewtocol.NewTCPConn("10.2.35.11", "9094")
	if err != nil {
		fmt.Println("网络连接：", err)
	}

	var dst uint = 1 // 目标地址，通常固定1

	// 读取D区数据,地址4到地址10
	areaDataListRead, err := mewtocol.ReadDataArea(c, dst, "D", 4, 10)
	if err != nil {
		fmt.Println("数据读取：", err)
	}
	fmt.Println("dataArea->", areaDataListRead)
}
