package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/opera/terago/thrift/tera"
	"log"
)

func AssertEmpty(ss ...string) {
	for _, s := range ss {
		if s == "" {
			panic(ss)
		}
	}
}

func main() {
	server := flag.String("server", "127.0.0.1:8118", "host:port")
	op := flag.String("op", "", "operation")
	table := flag.String("table", "", "table")
	key := flag.String("key", "", "key")
	value := flag.String("value", "", "value")
	flag.Parse()

	sock, err := thrift.NewTSocket(*server)
	if err != nil {
		log.Panicf("fail to dail %s, %#v", server, err)
	}
	trans := thrift.NewTTransportFactory().GetTransport(sock)
	proto := thrift.NewTBinaryProtocolFactoryDefault()
	client := tera.NewProxyClientFactory(trans, proto)
	if err := sock.Open(); err != nil {
		log.Panicf("open failed, server %s err %v", server, err)
	}

	switch *op {
	case "get":
		AssertEmpty(*table, *key)
		v, _ := client.Get(*table, *key)
		fmt.Println(v)
	case "put":
		AssertEmpty(*table, *key, *value)
		s, e := client.Put(*table, *key, *value)
		fmt.Printf("%v %v", s, e)
	default:
		fmt.Printf("-op <get|put> -key <key> -value <value>")
	}
}
