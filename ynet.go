package ynet

import (
	network "gitee.com/yuanxuezhe/ynet/tcp"
	"net"
)

type Callback func(conn net.Conn)

func NewTcpserver(addr string, callback Callback) *network.TCPServer {
	tcpServer := &network.TCPServer{
		Addr:            addr,
		MaxConnNum:      100,
		PendingWriteNum: 1000,
		Callback:        callback,
	}
	return tcpServer
}
