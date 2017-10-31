package terago

import (
	"fmt"
)

type Client struct {
	ConfPath string
}

func NewClient(conf_path, log_prefix string) (client ClientI, err error) {
	fmt.Println("new mock client")
	client = Client{
		ConfPath: conf_path,
	}
	return
}

func (c Client) Close() {
	fmt.Println("close mock client")
}
