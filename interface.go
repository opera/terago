package terago

import ()

type ClientI interface {
	Close()
	//	OpenTable(table_name string) (table TableI, err error)
}

type TableI interface {
	Close()
	PutKV(key, value string, ttl int) (err error)
	PutKVAsync(key, value string, ttl int) (err error)
	GetKV(key string) (value string, err error)
	DeleteKV(key string) (err error)
}
