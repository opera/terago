// +build !mock

package terago

/*
#cgo LDFLAGS: -ltera_c
#include "c/client.h"
*/
import "C"
import (
	"errors"
	"fmt"
)

type Client struct {
	ConfPath string
	CClient  *C.tera_client_t
}

func NewClient(conf_path, log_prefix string) (client Client, err error) {
	fmt.Println("new client")
	c := Client{
		CClient:  C.client_open(C.CString(conf_path), C.CString(log_prefix)),
		ConfPath: conf_path,
	}
	if c.CClient == nil {
		err = errors.New("Fail to create tera client")
	} else {
		client = c
	}
	return
}

func (c Client) Close() {
	fmt.Println("close client")
	if c.CClient != nil {
		C.tera_client_close(c.CClient)
	}
}

func (c Client) OpenKvStore(name string) (kv KvStore, err error) {
	fmt.Println("open table: " + name)
	if c.CClient == nil {
		err = errors.New("Fail to open table, client not available.")
		return
	}
	c_table := C.table_open(c.CClient, C.CString(name))
	if c_table == nil {
		err = fmt.Errorf("Fail to open table %s", name)
		return
	}
	kv = KvStore{Name: name, CTable: c_table}
	return
}
