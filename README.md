# Mewtocol

基于golang实现的松下mewtocol协议组件库，内部已实现TCP连接，后续会添加串口COM通讯

## 例



```golang
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
```



## 触点和数据代码

### 触点代码 

| 触点代码 | 说明       |
| -------- | ---------- |
| X        | 外部输入   |
| Y        | 外部输出   |
| R        | 内部继电器 |
| T        | 定时器     |
| C        | 计数器     |
| L        | 链接继电器 |



### 数据代码

| 数据代码 | 说明                |
| -------- | ------------------- |
| D        | 数据寄存器 DT       |
| L        | 链接寄存器 LD       |
| F        | 文件寄存器 FL       |
| S        | 目标值 SV           |
| K        | 经过值 EV           |
| IX       | 索引寄存器 IX       |
| IY       | 索引寄存器 IY       |
| WX       | 字单位外部输入 WX   |
| WY       | 字单位外部输出 WY   |
| WR       | 字单位内部继电器 WR |
| WL       | 字单位链接继电器 WL |

