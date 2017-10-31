package main

import (
	"fmt"
	"github.com/opera/terago"
	"plugin"
)

func main() {
	p, err := plugin.Open("terago.so")
	if err != nil {
		fmt.Printf("plugin.Open: %s\n", err)
		return
	}

	h, err := p.Lookup("NewClient")
	if err != nil {
		fmt.Printf("plugin lookup: %s\n", err)
		return
	}

	NewClient, ok := h.(func(string, string) (terago.ClientI, error))
	if !ok {
		fmt.Printf("plugin assert error\n")
		return
	}
	client, err := NewClient("./tera.flag", "terago_plugin")
	if err != nil {
		fmt.Printf("plugin NewClient: %s\n", err)
	}
	defer client.Close()
}
