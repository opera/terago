// +build mock

package terago

import (
	"errors"
	"fmt"
)

type KvStore struct {
	Name string
	Data map[string]string
}

func (p KvStore) Close() {
	fmt.Println("close mock table: " + p.Name)
}

// discard ttl in mock table
func (p KvStore) PutKV(key, value string, ttl int) (err error) {
	p.Data[key] = value
	return nil
}

func (p KvStore) PutKVAsync(key, value string, ttl int) (err error) {
	p.Data[key] = value
	return nil
}

func (p KvStore) GetKV(key string) (value string, err error) {
	var found bool
	value, found = p.Data[key]
	if !found {
		err = errors.New("NotExist")
	}
	return
}

func (p KvStore) DeleteKV(key string) (err error) {
	delete(p.Data, key)
	return nil
}
