package mewtocol

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	MAX_FLAME_SIZE int = 118
)

const DEFAULT_RESPONSE_TIMEOUT = 2000 // ms

type TcpClient struct {
	conn  *net.TCPConn
	ReqCh chan string
	ResCh chan string
	sync.Mutex
	// closed            bool
	responseTimeout time.Duration
	byteOrder       binary.ByteOrder
}

func NewTCPConn(ip, port string) (*TcpClient, error) {
	c := new(TcpClient)

	c.responseTimeout = DEFAULT_RESPONSE_TIMEOUT
	c.byteOrder = binary.BigEndian

	tcpAddr, err := net.ResolveTCPAddr("tcp", ip+":"+port)
	if err != nil {
		return nil, err
	}

	// raddr := &net.TCPConn{
	// 	IP:   net.ParseIP(remoteAddr),
	// 	Port: 5010,
	// }

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}

	if err = conn.SetKeepAlive(true); err != nil {
		return nil, err
	}

	c.conn = conn

	// todo：是否需要缓冲区
	c.ReqCh = make(chan string) // for request
	c.ResCh = make(chan string) // for response

	go c.listenLoop()
	return c, nil
}

func (c *TcpClient) listenLoop() {
	defer c.Close()

	// 在这里触发读取的逻辑
	for {
		sendStr := <-c.ReqCh
		err := send(c, sendStr)
		if err != nil {
			fmt.Println("报文发送失败：", err)
		}

		buf := make([]byte, 2048)
		n, err := bufio.NewReader(c.conn).Read(buf)

		if err != nil {
			fmt.Println("报文接收失败：", err)
			c.ResCh <- ""
		} else {
			// if isValidBCC(buf) {
			fmt.Println("接收报文为：", buf[0:n])
			fmt.Println("接收报文string结构为：", string(buf[0:n]))
			c.ResCh <- string(buf[0:n])

			// } else {
			// fmt.Println(fmt.Sprintf("invalid BCC:%s", string(buf)))
			// c.ResCh <- ""
			// }
		}
	}
}

func (c *TcpClient) Close() error {
	return c.conn.Close()
}

// send 发送请求报文
func send(c *TcpClient, sendStr string) error {
	sendData := []byte(sendStr + "\r")
	_, err := (*c.conn).Write(sendData)
	if err != nil {
		return err
	}

	return nil
}

// 如果数据附带的BCC和计算出的BCC一致则返回true
func isValidBCC(buff []byte) bool {
	lengthBeforeBCC := len(buff) - 2
	command := string(buff[:lengthBeforeBCC])
	bcc := string(buff[lengthBeforeBCC:])
	okBcc := getBcc(command)
	return (bcc == okBcc)
}
