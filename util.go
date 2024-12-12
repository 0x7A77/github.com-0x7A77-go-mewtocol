package mewtocol

import (
	"fmt"
	"log"
)

// format 将消息头%以及BCC加入到命令报文中
func format(dstAddress uint, body string) string {
	command := header() + address(dstAddress) + command() + body
	sendData := command + getBcc(command)
	return sendData
}

func header() string {
	return "%"
}

// 站号：1-255
func address(ad uint) string {
	if (ad < 1 || 32 < ad) && ad != 255 {
		panic(fmt.Sprintf("Invalid mewtocol address: %d", ad))
	}

	return fmt.Sprintf("%02d", ad)
}

// isValidCode 指定是否可以连接，并返回OK的布尔值
func isValidCode(code string, list []string) bool {
	for _, s := range list {
		if code == s {
			return true
		}
	}

	log.Fatal(fmt.Sprintf("invalid code:%s", code))
	return false
}

func command() string {
	return "#"
}
