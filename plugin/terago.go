package main

import (
	"C"
	"github.com/opera/terago"
)

func NewClient(conf_path, log_prefix string) (client terago.ClientI, err error) {
	return terago.NewClient(conf_path, log_prefix)
}
