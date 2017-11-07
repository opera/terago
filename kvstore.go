// +build prod

package terago

/*
#cgo LDFLAGS: -ltera_c
#include "c/kvstore.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type KvStore struct {
	Name   string
	CTable *C.tera_table_t
}

func (p KvStore) Close() {
	fmt.Println("close table: " + p.Name)
	if p.CTable != nil {
		C.tera_table_close(p.CTable)
	}
}

// ttl(time-to-live)
// Key-value will expired after <ttl> seconds. -1 means never expired.
func (p KvStore) Put(key, value string, ttl int) (err error) {
	if p.CTable == nil {
		return errors.New("table not open: " + p.Name)
	}
	ret := C.table_put_kv_sync(p.CTable, C.CString(key), C.int(len(key)),
		C.CString(value), C.int(len(value)), C.int(ttl))
	if !ret {
		err = errors.New("put kv error")
	}
	return
}

// Async put key-value into tera. Return success immediately and run put operation at background.
// Caution: If put failed, specify kv would be dump to error log.
func (p KvStore) PutAsync(key, value string, ttl int) (err error) {
	if p.CTable == nil {
		return errors.New("table not open: " + p.Name)
	}
	C.table_put_kv_async(p.CTable, C.CString(key), C.int(len(key)),
		C.CString(value), C.int(len(value)), C.int(ttl))
	return
}

func (p KvStore) Get(key string) (value string, err error) {
	if p.CTable == nil {
		err = errors.New("table not open: " + p.Name)
		return
	}
	var vallen C.int
	vc := C.table_get_kv_sync(p.CTable, C.CString(key), C.int(len(key)), (*C.int)(&vallen))
	if vallen >= 0 {
		value = C.GoStringN(vc, vallen)
		C.free(unsafe.Pointer(vc))
	} else {
		err = errors.New("key not found")
		value = ""
	}
	return
}

func (p KvStore) BatchPut(keys, values []string) (errs []error) {
	return
}

func (p KvStore) BatchGet(keys []string) (values []string, errs []error) {
	return
}

func (p KvStore) RangeGet(start, end string, maxNum int) (keys, values []string, err error) {
	return
}

func (p KvStore) Delete(key string) (err error) {
	if p.CTable == nil {
		return errors.New("table not open: " + p.Name)
	}
	ret := C.table_delete_kv_sync(p.CTable, C.CString(key), C.int(len(key)))
	if !ret {
		err = errors.New("put kv error")
	}
	return
}
