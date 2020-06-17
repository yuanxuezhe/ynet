package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
	network "ynet/tcp"
)

func Handler(conn net.Conn) {
	for {
		//_,err := io.Copy(conn, conn)
		//buff, err := network.ReadMsg(conn)
		buff, err := network.ReadMsgs(conn)
		if err != nil {
			break
		}
		fmt.Println(buff)

		//network.SendMsg(conn,[]byte("Hello,Recv msg:" + string(buff)))

		//network.SendMsg(conn,buff)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	tcpServer := &network.TCPServer{
		Addr:            ":8080",
		MaxConnNum:      100,
		PendingWriteNum: 1000,
		Callback:        Handler,
	}
	tcpServer.Start()

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	fmt.Printf("System closing down (signal: %v)", sig)

	tcpServer.Close()
}
