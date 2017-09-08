package main

import (
	"fmt"
	"github.com/opera/terago"
)

func main() {
	fmt.Println("Hello terago!")

	// New a tera client
	// One client support multiple tables
	client, c_err := terago.NewClient("tera.flag", "terago")
	defer client.Close() // Donot forget
	if c_err != nil {
		panic("tera.NewClient error: " + c_err.Error())
	}

	// Open a tera table.
	// One table can be used in many goroutines.
	table, t_err := client.OpenTable("terago")
	defer table.Close() // Donot forget
	if t_err != nil {
		panic("tera.OpenTable error: " + t_err.Error())
	}

	// Put a key-value synchronously.
	p_err := table.PutKV("hello", "terago", 10)
	if p_err != nil {
		panic("put key value error: " + p_err.Error())
	}

	p_err = table.PutKVAsync("helloasync", "terago", 10)
	if p_err != nil {
		panic("put key value async error: " + p_err.Error())
	}

	// Get a key-value synchronously.
	value, g_err := table.GetKV("hello")
	if g_err != nil {
		panic("get key value error: " + g_err.Error())
	}
	fmt.Printf("GetKV: hello:%s.\n", value)

	// Delete a key-value synchronously.
	d_err := table.DeleteKV("hello")
	if d_err != nil {
		panic("delete key value error: " + d_err.Error())
	}
}
