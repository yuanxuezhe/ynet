package main

import (
	"fmt"
	"gitee.com/yuanxuezhe/ynet/tcp"
	"gitee.com/yuanxuezhe/ynet/yconnpool"
	"net"
	"strconv"

	//"io/ioutil"
	//"testing"
	"time"
)

func main() {
	var err error
	cp, _ := yconnpool.NewConnPool(func() (yconnpool.ConnRes, error) {
		return net.Dial("tcp", ":8080")
	}, 10, time.Second*10)

	for i := 0; i < 10; i++ {
		conn, _ := cp.Get()
		err = network.SendMsg(conn.(net.Conn), []byte("YUANSHUAI<==>WANYUAN  "+strconv.Itoa(i)))
		if err != nil {
			fmt.Printf("%s", err)
		}
		buff, err := network.ReadMsg(conn.(net.Conn))
		if err != nil {
			fmt.Printf("%s", err)
		}
		fmt.Println(conn.(net.Conn).LocalAddr(), "==>", conn.(net.Conn).RemoteAddr(), "    ", string(buff))
		cp.Put(conn)
	}

	cp.Close()
}
