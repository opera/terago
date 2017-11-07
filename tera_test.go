package terago

import (
	"fmt"
	"testing"
	"time"
)

func logExecTime(start time.Time, prefix string) {
	elapsed_ms := time.Since(start) / time.Millisecond
	fmt.Printf("Performance: %s cost %d ms.\n", prefix, elapsed_ms)
}

func TestTera(*testing.T) {
	fmt.Println("Hello terago!")
	start := time.Now()
	client, c_err := NewClient("tera.flag", "terago")
	defer client.Close()
	if c_err != nil {
		panic("tera.NewClient error: " + c_err.Error())
	}
	logExecTime(start, "NewClient")

	start = time.Now()
	kv, t_err := client.OpenKvStore("terago")
	defer kv.Close()
	if t_err != nil {
		panic("tera.OpenTable error: " + t_err.Error())
	}
	logExecTime(start, "OpenTable")

	start = time.Now()
	p_err := kv.Put("hello", "terago", 10)
	if p_err != nil {
		panic("put key value error: " + p_err.Error())
	}
	logExecTime(start, "Put")

	start = time.Now()
	p_err = kv.PutAsync("helloasync", "terago", 10)
	if p_err != nil {
		panic("put key value async error: " + p_err.Error())
	}
	logExecTime(start, "PutAsync")

	start = time.Now()
	// get an exist key value, return value
	value, g_err := kv.Get("hello")
	if g_err != nil {
		panic("get key value error: " + g_err.Error())
	}
	fmt.Printf("get key[%s] value[%s].\n", "hello", value)
	logExecTime(start, "Get")

	start = time.Now()
	// get a not-exist key value, return "not found"
	_, g_err = kv.Get("hell")
	if g_err == nil {
		panic("get key value should fail: " + g_err.Error())
	}
	logExecTime(start, "Get_NotExist")

	start = time.Now()
	d_err := kv.Delete("hello")
	if d_err != nil {
		panic("delete key value error: " + d_err.Error())
	}
	logExecTime(start, "Delete")

	_, g_err = kv.Get("hello")
	if g_err == nil {
		panic("get key value should fail: " + g_err.Error())
	}
}
