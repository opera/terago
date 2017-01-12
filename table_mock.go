// +build mock

package tera

import (
	"fmt"
)

type Table struct {
	Name string
	Data map[string]string
}

func (t *Table) Close() {
	fmt.Println("close mock table: " + t.Name)
}

// discard ttl in mock table
func (t *Table) PutKV(key, value string, ttl int) (err error) {
	t.Data[key] = value
	return nil
}

func (t *Table) GetKV(key string) (value string, found bool, err error) {
	value, found = t.Data[key]
	return
}

func (t *Table) DeleteKV(key string) (err error) {
	delete(t.Data, key)
	return nil
}
