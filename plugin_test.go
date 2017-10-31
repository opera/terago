package terago

import (
	"fmt"
	"plugin"
	"reflect"
	"testing"
)

func Test(*testing.T) {
	p, err := plugin.Open("./terago.so")
	if err != nil {
		fmt.Printf("plugin.Open: %s\n", err)
		return
	}

	h, err := p.Lookup("NewClient")
	if err != nil {
		fmt.Printf("plugin lookup: %s\n", err)
		return
	}

	NewClient := h.(func(string, string) (ClientI, error))
	c, err := NewClient("./tera.flag", "terago_plugin")
	if err != nil {
		fmt.Printf("plugin NewClient: %s\n", err)
	}
	fmt.Printf("efin %v\n", reflect.TypeOf(c))
	client := c.(ClientI)
	defer client.Close()

	table, err := client.OpenTable("plugin")
	if err != nil {
		fmt.Printf("plugin OpenTable: %s\n", err)
	}
	defer table.Close()

	{
		err := table.PutKV("hello", "terago", 10)
		if err != nil {
			panic("put key value error: " + err.Error())
		}
	}

	{
		value, err := table.GetKV("hello")
		if err != nil {
			panic("get key value error: " + err.Error())
		}
		fmt.Printf("GetKV: hello:%s.\n", value)
	}
}
