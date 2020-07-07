package ynet

import (
	conn "gitee.com/yuanxuezhe/ynet/Conn"
	httpnet "gitee.com/yuanxuezhe/ynet/http"
	tcpnet "gitee.com/yuanxuezhe/ynet/tcp"
)

type Callback func(conn conn.CommConn)

func NewTcpserver(addr string, callback Callback) *tcpnet.TCPServer {
	tcpServer := &tcpnet.TCPServer{
		Addr:            addr,
		MaxConnNum:      100,
		PendingWriteNum: 1000,
		Callback:        callback,
	}
	return tcpServer
}

func NewWsserver(addr string, callback Callback) *httpnet.WSServer {
	wsServer := &httpnet.WSServer{
		Addr:            addr,
		MaxConnNum:      100,
		PendingWriteNum: 1000,
		Callback:        callback,
	}
	return wsServer
}

func NewTcpclient(addr string) conn.CommConn {
	conn := &tcpnet.TCPClient{
		Addr: addr,
	}

	return conn.Connect()
}

func NewWsclient(addr string) conn.CommConn {
	conn := &httpnet.WSClient{
		Addr: addr,
	}

	return conn.Connect()
}
