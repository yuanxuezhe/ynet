package main

import (
	"fmt"
	"gitee.com/yuanxuezhe/ynet"
	conn "gitee.com/yuanxuezhe/ynet/Conn"
	"os"
	"os/signal"
	"time"
)

//func Handler1(conn net.Conn) {
//	for {
//
//		msgs, err := network.ReadMsgs(conn)
//		if err != nil {
//			break
//		}
//		for _, msg := range msgs {
//			fmt.Println(msg)
//			network.SendMsg(conn,[]byte("Hello,Recv msg:" + msg))
//		}
//
//		time.Sleep(10 * time.Millisecond)
//	}
//}

func Handler(conn conn.CommConn) {
	for {
		//buff, err := network.ReadMsg(conn)
		buff, err := conn.ReadMsg()
		if err != nil {
			break
		}

		fmt.Println(string(buff))

		conn.WriteMsg([]byte("Hello,Recv msg:" + string(buff)))
		//network.SendMsg(conn, []byte("Hello,Recv msg:"+string(buff)))

		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	tcpServer := ynet.NewTcpserver(":8080", Handler)
	tcpServer.Start()

	wsServer := ynet.NewWsserver(":8090", Handler)
	wsServer.Start()

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	fmt.Printf("System closing down (signal: %v)", sig)

	tcpServer.Close()
}
