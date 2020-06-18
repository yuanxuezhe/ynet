package main

import (
	"fmt"
	"strconv"

	//"gitee.com/yuanxuezhe/ynet/tcp"
	"gitee.com/yuanxuezhe/ynet/yconnpool"
	//"net"
	//"strconv"
	"github.com/garyburd/redigo/redis"
	//"io/ioutil"
	//"testing"
	"time"
)

func main() {
	var err error
	//cp, _ := yconnpool.NewConnPool(func() (yconnpool.ConnRes, error) {
	//	return net.Dial("tcp", ":8080")
	//}, 10, time.Second*10)
	//
	//for i := 0; i < 10; i++ {
	//	conn, _ := cp.Get()
	//	err = network.SendMsg(conn.(net.Conn), []byte("YUANSHUAI<==>WANYUAN  "+strconv.Itoa(i)))
	//	if err != nil {
	//		fmt.Printf("%s", err)
	//	}
	//	buff, err := network.ReadMsg(conn.(net.Conn))
	//	if err != nil {
	//		fmt.Printf("%s", err)
	//	}
	//	fmt.Println(conn.(net.Conn).LocalAddr(), "==>", conn.(net.Conn).RemoteAddr(), "    ", string(buff))
	//	cp.Put(conn)
	//}
	//
	//cp.Close()

	cp, _ := yconnpool.NewConnPool(func() (yconnpool.ConnRes, error) {
		return redis.Dial("tcp", "127.0.0.1:6379")
	}, 100, time.Second*10)
	for i := 0; i < 10; i++ {
		conn, _ := cp.Get()

		_, err = conn.(redis.Conn).Do("SET", "key"+strconv.Itoa(i), "duzhenxun"+strconv.Itoa(11*i))
		if err != nil {
			fmt.Println("redis set failed:", err)
		}
		rs, err := conn.(redis.Conn).Do("GET", "key"+strconv.Itoa(i))

		if err != nil {
			fmt.Println("redis get failed:", err)
		} else {
			fmt.Printf("Get mykey: %s \n", rs)
		}
		cp.Put(conn)
	}

	var p1 struct {
		Title  string `redis:"title"`
		Author string `redis:"author"`
		Body   string `redis:"body"`
	}

	m := map[string]string{
		"title":  "Example2",
		"author": "Steve",
		"body":   "Map",
	}

	for i := 0; i < 10; i++ {
		conn, _ := cp.Get()

		_, err = conn.(redis.Conn).Do("HMSET", redis.Args{}.Add("idd"+strconv.Itoa(i)).AddFlat(m)...)
		if err != nil {
			fmt.Println("redis set failed:", err)
		}
		cp.Put(conn)
	}
	conn, _ := cp.Get()
	v, err := redis.Values(conn.(redis.Conn).Do("HGETALL", "idd2"))
	if err != nil {
		panic(err)
	}

	if err := redis.ScanStruct(v, &p1); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", p1)
	cp.Put(conn)
	cp.Close()
}
