package main

import (
	"fmt"
	"github.com/opera/terago"
	"github.com/opera/terago/thrift/tera"
)

type Handler struct {
	client *terago.Client
	tables map[string]*terago.KvStore
}

func main() {
	fmt.Println("Hello terago!")

	// New a tera client
	// One client support multiple kvstore
	client, c_err := terago.NewClient("tera.flag", "terago")
	defer client.Close() // Donot forget
	if c_err != nil {
		panic("tera.NewClient error: " + c_err.Error())
	}

	// Open a tera kvstore.
	// One kvstore can be used in many goroutines.
	kv, t_err := client.OpenKvStore("terago")
	defer kv.Close() // Donot forget
	if t_err != nil {
		panic("tera.OpenTable error: " + t_err.Error())
	}

	// Put a key-value synchronously.
	p_err := kv.Put("hello", "terago", 10)
	if p_err != nil {
		panic("put key value error: " + p_err.Error())
	}

	p_err = kv.PutAsync("helloasync", "terago", 10)
	if p_err != nil {
		panic("put key value async error: " + p_err.Error())
	}

	// Get a key-value synchronously.
	value, g_err := kv.Get("hello")
	if g_err != nil {
		panic("get key value error: " + g_err.Error())
	}
	fmt.Printf("Get: hello:%s.\n", value)

	// Delete a key-value synchronously.
	d_err := kv.Delete("hello")
	if d_err != nil {
		panic("delete key value error: " + d_err.Error())
	}
}
