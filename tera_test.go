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
		panic("tera.OpenKvStore error: " + t_err.Error())
	}
	logExecTime(start, "OpenKvStore")

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

func TestTeraBatch(*testing.T) {
	fmt.Println("Hello terago batch!")
	start := time.Now()
	client, err := NewClient("tera.flag", "terago")
	defer client.Close()
	if err != nil {
		panic("tera.NewClient error: " + err.Error())
	}
	logExecTime(start, "NewClient")

	start = time.Now()
	kv, err := client.OpenKvStore("terago")
	defer kv.Close()
	if err != nil {
		panic("tera.OpenKvStore error: " + err.Error())
	}
	logExecTime(start, "OpenKvStore")

	//	for {
	keys := []string{"t", "e", "r", "a", "go"}
	values := []string{"tt", "ee", "rr", "aa", "gogo"}
	start = time.Now()
	var kvs []KeyValue
	for i, k := range keys {
		kvs = append(kvs, KeyValue{Key: k, Value: values[i], TTL: 10})
	}
	err = kv.BatchPut(kvs)
	if err != nil {
		panic("BatchPut error: " + err.Error())
	}
	logExecTime(start, "BatchPut")

	start = time.Now()
	// get an exist key value, return value
	kvs, err = kv.BatchGet(keys)
	if err != nil {
		panic("BatchGet error: " + err.Error())
	}
	fmt.Printf("BatchGet KeyValues[%v]\n", kvs)
	logExecTime(start, "BatchGet")

	start = time.Now()
	kvs, err = kv.RangeGet("a", "h", 10)
	if err != nil {
		panic("RangeGet error: " + err.Error())
	}
	fmt.Printf("RangeGet KeyValues[%v]\n", kvs)
	if len(kvs) != 3 || kvs[0].Key != "a" || kvs[1].Key != "e" || kvs[2].Key != "go" {
		e := fmt.Errorf("RangeGet err: %v", kvs)
		panic(e)
	}
	logExecTime(start, "RangeGet")

	start = time.Now()
	kvs, err = kv.RangeGet("a", "h", 2)
	if err != nil {
		panic("RangeGet error: " + err.Error())
	}
	fmt.Printf("RangeGet2 KeyValues[%v]\n", kvs)
	if len(kvs) != 2 || kvs[0].Key != "a" || kvs[1].Key != "e" {
		e := fmt.Errorf("RangeGet2 err: %v", kvs)
		panic(e)
	}
	logExecTime(start, "RangeGet2")
	//	}
}

func TestOpenFailed(*testing.T) {
	fmt.Println("Hello terago!")
	start := time.Now()
	client, err := NewClient("tera.flag", "terago")
	defer client.Close()
	if err != nil {
		panic("tera.NewClient error: " + err.Error())
	}
	logExecTime(start, "NewClient")

	start = time.Now()
	kv, err := client.OpenKvStore("teragooo")
	defer kv.Close()
	if err == nil {
		panic("tera.OpenKvStore should failed: ")
	}
	fmt.Printf("%v\n", err)
	logExecTime(start, "OpenKvStore")
}
