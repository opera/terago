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
	client, err := NewClient("./tera.flag", "terago_plugin")
	if err != nil {
		fmt.Printf("plugin NewClient: %s\n", err)
	}
	fmt.Printf("%v\n", reflect.TypeOf(client))
	defer client.Close()
}
