// +build mock

package terago

import (
	"fmt"
)

type Client struct {
	ConfPath string
}

func NewClient(conf_path, log_prefix string) (client Client, err error) {
	fmt.Println("new mock client")
	client = Client{
		ConfPath: conf_path,
	}
	return
}

func (c Client) Close() {
	fmt.Println("close mock client")
}

func (c Client) OpenKvStore(name string) (kv KvStore, err error) {
	fmt.Println("open mock table: " + name)
	if name == "terago" {
		kv = KvStore{Name: name, Data: make(map[string]string)}
	} else {
		err = fmt.Errorf("table not exist: %s", name)
	}
	return
}
