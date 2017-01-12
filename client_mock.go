// +build mock

package tera

import (
	"fmt"
)

type Client struct {
	ConfPath string
}

func NewClient(conf_path, log_prefix string) (client Client, err error) {
	fmt.Println("new mock client")
	client.ConfPath = conf_path
	return
}

func (c *Client) Close() {
	fmt.Println("close mock client")
}

func (c *Client) OpenTable(table_name string) (table Table, err error) {
	fmt.Println("open mock table: " + table_name)
	table = Table{Name: table_name, Data: make(map[string]string)}
	return
}
