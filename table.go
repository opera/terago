package tera

/*
#cgo LDFLAGS: -ltera_c
#include <stdio.h>
#include <stdlib.h>
#include <tera_c.h>
bool table_put_kv_sync(tera_table_t* table, const char* key, int keylen,
                       const char* value, int vallen, int ttl) {
	char* err = NULL;
	bool ret = tera_table_put_kv(table, key, keylen, value, vallen, ttl, &err);
	if (!ret) {
		fprintf(stderr, "tera put kv error: %s.\n", err);
		free(err);
	}
	return ret;
}

char* table_get_kv_sync(tera_table_t* table, const char* key, int keylen, int* vallen) {
	uint64_t vlen = 0;
	char* err = NULL;
	char* value;
	bool ret = tera_table_get(table, key, keylen, "", "", 0, &value, &vlen, &err, 0);
	if (ret) {
	  *vallen = (int)vlen;
  } else {
		*vallen = -1;
		fprintf(stderr, "tera get kv error: %s.\n", err);
		free(err);
	}
  return value;
}

bool table_delete_kv_sync(tera_table_t* table, const char* key, int keylen) {
	bool ret = tera_table_delete(table, key, keylen, "", "", 0);
	if (!ret) {
		fprintf(stderr, "tera delete error.\n");
	}
	return ret;
}
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type Table struct {
	Name   string
	CTable *C.tera_table_t
}

func (t *Table) Close() {
	fmt.Println("close table: " + t.Name)
	if t.CTable != nil {
		C.tera_table_close(t.CTable)
	}
}

// ttl(time-to-live)
// Key-value will expired after <ttl> seconds. 0 means never expired.
func (t *Table) PutKV(key, value string, ttl int) (err error) {
	if t.CTable == nil {
		return errors.New("table not open: " + t.Name)
	}
	ret := C.table_put_kv_sync(t.CTable, C.CString(key), C.int(len(key)),
		C.CString(value), C.int(len(value)), C.int(ttl))
	if !ret {
		err = errors.New("put kv error")
	}
	return
}

func (t *Table) GetKV(key string) (value string, err error) {
	if t.CTable == nil {
		err = errors.New("table not open: " + t.Name)
		return
	}
	var vallen C.int
	vc := C.table_get_kv_sync(t.CTable, C.CString(key), C.int(len(key)), (*C.int)(&vallen))
	if vallen >= 0 {
		value = C.GoStringN(vc, vallen)
		C.free(unsafe.Pointer(vc))
	} else {
		err = errors.New("key not found")
		value = ""
	}
	return
}

func (t *Table) DeleteKV(key string) (err error) {
	if t.CTable == nil {
		return errors.New("table not open: " + t.Name)
	}
	ret := C.table_delete_kv_sync(t.CTable, C.CString(key), C.int(len(key)))
	if !ret {
		err = errors.New("put kv error")
	}
	return
}
