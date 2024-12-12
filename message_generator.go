package mewtocol

import (
	"fmt"
	"strings"
)

// formatReadIOSingle 生成并返回读取触点的报文
func formatReadIOSingle(dstAddress uint, contactCode string, contactNo uint) string {
	if !isValidCode(contactCode, []string{"X", "Y", "R", "L", "T", "C"}) {
		panic(fmt.Sprintln("invalid code:", contactCode))
	}

	command := "RCS" + contactCode + fmt.Sprintf("%04d", contactNo)
	return format(dstAddress, command)
}

// formatReadIOWord  生成并返回以指定的文件单位读取触点的报文
func formatReadIOWord(dstAddress uint, contactCode string, startWordNo uint, endWordNo uint) string {
	if !isValidCode(contactCode, []string{"X", "Y", "R", "L", "T", "C"}) {
		panic(fmt.Sprintln("invalid code:", contactCode))
	}

	command := "RCC" + contactCode + fmt.Sprintf("%04d", startWordNo) + fmt.Sprintf("%04d", endWordNo)
	return format(dstAddress, command)
}

// 生成并返回写入指定触点的报文
func formatWriteIOSingle(dstAddress uint, contactCode string, contactNo uint, state bool) string {
	if !isValidCode(contactCode, []string{"Y", "R", "L"}) {
		panic(fmt.Sprintln("invalid code:", contactCode))
	}
	contactData := ""

	if state {
		contactData = "1"
	} else {
		contactData = "0"
	}

	command := "WCS" + contactCode + fmt.Sprintf("%04d", contactNo) + contactData
	return format(dstAddress, command)
}

// formatReadDataArea 生成并返回读取以指定的地址长度的数据报文
func formatReadDataArea(dstAddress uint, dataCode string, startWordNo uint, endWordNo uint) string {
	if !isValidCode(dataCode, []string{"D", "L", "F"}) {
		panic(fmt.Sprintln("invalid code:", dataCode))
	}

	command := "RD" + dataCode + fmt.Sprintf("%05d", startWordNo) + fmt.Sprintf("%05d", endWordNo)
	return format(dstAddress, command)
}

// formatWriteDataArea 生成并返回写入以指定的地址长度的数据报文
func formatWriteDataArea(dstAddress uint, dataCode string, startWordNo uint, values []uint32) string {
	if !isValidCode(dataCode, []string{"D", "L", "F"}) {
		panic(fmt.Sprintln("invalid code:", dataCode))
	}

	valuesBin := make([]byte, 0, len(values))
	index := 0
	for i, val := range values {
		hex := []byte(fmt.Sprintf("%04x", val))
		upper := hex[:2]
		lower := hex[2:]
		valBin := append(lower, upper...)
		valuesBin = append(valuesBin, valBin...)
		index = i
	}

	endWordNo := startWordNo + uint(index)
	command := "WD" + dataCode + fmt.Sprintf("%05d", startWordNo) + fmt.Sprintf("%05d", endWordNo)
	command += strings.ToUpper(string(valuesBin))
	return format(dstAddress, command)
}
