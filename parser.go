package mewtocol

import (
	"fmt"
	"strconv"
)

type Mewtocol struct {
	Header  string
	Address uint16
	Code    string
	Body    []byte
}

// parseReadIOSingle 从PLC获取接点单体响应，并返回值
func parseReadIOSingle(str string) (bool, error) {
	res, err := parseHeader(str)
	if err != nil {
		return false, err
	}

	contactData := string(res.Body)

	if contactData == "1" {
		return true, nil
	} else {
		return false, nil
	}
}

// parseReadIOWord 解析来自PLC的接点字单体取得应答并返回
func parseReadIOWord(str string) ([]uint, error) {
	res, err := parseHeader(str)
	if err != nil {
		return nil, err
	}

	return parseListData(res.Body), nil
}

// parseWriteIOSingle 解析来自PLC的接点输出单体写入响应并返回
func parseWriteIOSingle(str string) (bool, error) {
	_, err := parseHeader(str)
	if err != nil {
		return false, err
	}

	return true, nil
}

// parseReadDataArea 解析来自PLC的数据区域获取响应并返回值
func parseReadDataArea(str string) ([]uint, error) {
	res, err := parseHeader(str)
	if err != nil {
		return nil, err
	}

	return parseListData(res.Body), nil
}

// parseWriteDataArea 解析来自PLC的接点输出单体写入响应并返回
func parseWriteDataArea(str string) (bool, error) {
	_, err := parseHeader(str)
	if err != nil {
		return false, err
	}

	return true, nil
}

// parseListData 将字节的切片转为2个字节的数值，并返回数值的切片
func parseListData(data []byte) []uint {
	count := len(data) / 4
	list := make([]uint, 0, count)

	for i := 0; i < count; i++ {
		n := i * 4

		// 1、2字节在后，3、4字节在前
		valLower := data[(n + 0):(n + 2)]
		valUpper := data[(n + 2):(n + 4)]
		val := make([]byte, 0, 4)
		val = append(val, valUpper...)
		val = append(val, valLower...)
		intVal, _ := strconv.ParseUint(string(val), 16, 32)
		list = append(list, uint(intVal))
	}

	return list
}

func parseHeader(str string) (*Mewtocol, error) {
	buff := []byte(str)
	success := string(buff[3])

	if success == "$" {
		header := string(buff[0])
		address, _ := strconv.ParseInt(string(buff[1:3]), 10, 16)
		code := string(buff[4:6])
		body := getReqBody(buff)
		return &Mewtocol{header, uint16(address), code, body}, nil
	} else if success == "!" {
		errNo, _ := strconv.ParseInt(string(buff[4:6]), 16, 16)
		return nil, fmt.Errorf(fmt.Sprintf("mewtocol error response:%x", errNo))
	} else {
		panic(fmt.Sprintf("invalid success code:%s", success))
	}
}

func getReqBody(buff []byte) []byte {
	headerSize := 6
	footerSize := 2
	bodySize := len(buff) - (headerSize + footerSize)
	bodyStart := headerSize
	bodyEnd := headerSize + bodySize
	return buff[bodyStart:bodyEnd]
}
