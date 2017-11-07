// +build mock

package terago

import (
	"errors"
	"fmt"
	"sort"
)

type KvStore struct {
	Name string
	Data map[string]string
}

func (p KvStore) Close() {
	fmt.Println("close mock table: " + p.Name)
}

// discard ttl in mock table
func (p KvStore) Put(key, value string, ttl int) (err error) {
	p.Data[key] = value
	return nil
}

func (p KvStore) PutAsync(key, value string, ttl int) (err error) {
	p.Data[key] = value
	return nil
}

func (p KvStore) Get(key string) (value string, err error) {
	var found bool
	value, found = p.Data[key]
	if !found {
		err = errors.New("NotExist")
	}
	return
}

func (p KvStore) BatchPut(kvs []KeyValue) (err error) {
	for _, kv := range kvs {
		p.Data[kv.Key] = kv.Value
	}
	return nil
}

func (p KvStore) BatchGet(keys []string) (result []KeyValue, err error) {
	for _, key := range keys {
		value, ok := p.Data[key]
		if ok {
			result = append(result, KeyValue{Key: key, Value: value})
		} else {
			result = append(result, KeyValue{Key: key, Err: errors.New("NotFound")})
			err = errors.New("NotFound")
		}
	}
	return
}

func (p KvStore) RangeGet(start, end string, maxNum int) (result []KeyValue, err error) {
	var keysSort []string
	for k, _ := range p.Data {
		keysSort = append(keysSort, k)
	}
	sort.Strings(keysSort)
	cnt := 0
	for _, k := range keysSort {
		if k >= start && k < end {
			result = append(result, KeyValue{Key: k, Value: p.Data[k]})
			cnt++
			if cnt >= maxNum {
				break
			}
		}
	}
	return
}

func (p KvStore) Delete(key string) (err error) {
	delete(p.Data, key)
	return nil
}
