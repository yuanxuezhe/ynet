package ynet

import (
	"net"
	network "ynet/tcp"
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
