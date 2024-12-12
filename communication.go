package mewtocol

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

// ReadIOSingle 读取指定的1接点状态
func ReadIOSingle(c *TcpClient, dst uint, contactCode string, contactNo uint) (bool, error) {
	fmt.Println("channel!")
	sendStr := formatReadIOSingle(dst, contactCode, contactNo)
	c.ReqCh <- sendStr
	recvStr := <-c.ResCh

	if recvStr == "" {
		return false, errors.New("failed to communicate with PLC")
	} else {
		return parseReadIOSingle(recvStr)
	}
}

// ReadIOWord 读取指定接点的状态
func ReadIOWord(c *TcpClient, dst uint, contactCode string, startNo uint, endNo uint) ([]uint, error) {
	sendStr := formatReadIOWord(dst, contactCode, startNo, endNo)
	c.ReqCh <- sendStr
	recvStr := <-c.ResCh

	if recvStr == "" {
		return nil, errors.New("failed to communicate with PLC")
	} else {
		return parseReadIOWord(recvStr)
	}
}

// WriteIOSingle 写入指定的1节点输出，指定输出ON时true，OFF时状态为false
func WriteIOSingle(c *TcpClient, dst uint, contactCode string, contactNo uint, state bool) (bool, error) {
	sendStr := formatWriteIOSingle(dst, contactCode, contactNo, state)
	c.ReqCh <- sendStr
	recvStr := <-c.ResCh

	if recvStr == "" {
		return false, errors.New("failed to communicate with PLC")
	} else {
		return parseWriteIOSingle(recvStr)
	}
}

// ReadDataArea 读取指定数据区域的状态
func ReadDataArea(c *TcpClient, dst uint, dataCode string, startNo uint, endNo uint) ([]uint, error) {
	sendStr := formatReadDataArea(dst, dataCode, startNo, endNo)
	log.Println("sendStr", sendStr)
	c.ReqCh <- sendStr
	recvStr := <-c.ResCh

	if recvStr == "" {
		return nil, errors.New("failed to communicate with PLC")
	} else {
		return parseReadDataArea(recvStr)
	}
}

// 将作为匹配列的引数传递的值写入数据区域
func WriteDataArea(c *TcpClient, dst uint, dataCode string, startNo uint, values []uint32) (bool, error) {
	sendStr := formatWriteDataArea(dst, dataCode, startNo, values)
	c.ReqCh <- sendStr
	recvStr := <-c.ResCh

	if recvStr == "" {
		return false, errors.New("failed to communicate with PLC")
	} else {
		return parseWriteDataArea(recvStr)
	}
}

// 生成并返回用于水平奇偶校验的代码(2字节)
// 将从da到bcc之前的数据每取1byte的排他逻辑和的结果作为16进制字符串返回
func getBcc(str string) string {
	buff := []byte(str)
	result := buff[0]

	for i := 1; i < len(buff); i++ {
		result = result ^ (buff[i])
	}
	return strings.ToUpper(fmt.Sprintf("%02x", result))
}
