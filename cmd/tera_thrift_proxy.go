package main

import (
	"errors"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/opera/terago"
	"github.com/opera/terago/thrift/tera"
	"log"
)

type Handler struct {
	client   *terago.Client
	kvStores map[string]*terago.KvStore
}

func (p *Handler) getTable(name string) *terago.KvStore {
	if kv, ok := p.kvStores[name]; ok {
		return kv
	}

	kv, e := p.client.OpenKvStore(name)
	if e != nil {
		return nil
	}
	p.kvStores[name] = &kv
	return &kv
}

func (p *Handler) Get(table string, key string) (r *tera.KeyValue, err error) {
	kvStore := p.getTable(table)
	if kvStore == nil {
		return &tera.KeyValue{Key: key, Status: tera.Status_TableNotExist}, errors.New("table not exist")
	}
	value, e := kvStore.Get(key)
	if e == nil {
		r = &tera.KeyValue{Key: key, Value: value, Status: tera.Status_Ok}
	} else {
		r = &tera.KeyValue{Key: key, Status: tera.Status_NotFound}
	}
	return
}

func (p *Handler) Put(table string, kv *tera.KeyValue) (r tera.Status, err error) {
	kvStore := p.getTable(table)
	if kvStore == nil {
		return tera.Status_TableNotExist, errors.New("table not exist")
	}
	if kv.TTL == 0 {
		kv.TTL = -1
	}
	e := kvStore.Put(kv.Key, kv.Value, int(kv.TTL))
	if e != nil {
		log.Println(e)
	}
	return tera.Status_Ok, nil
}

func (p *Handler) BatchGet(table string, keys []string) (r []*tera.KeyValue, err error) {
	return
}

func (p *Handler) BatchPut(table string, kvs []*tera.KeyValue) (r []tera.Status, err error) {
	return
}

func main() {
	log.Println("Hello terago!")

	// New a tera client
	// One client support multiple kvstore
	client, err := terago.NewClient("tera.flag", "terago")
	defer client.Close() // Donot forget
	if err != nil {
		panic(err)
	}
	handler := Handler{
		client:   &client,
		kvStores: make(map[string]*terago.KvStore),
	}
	addr := ":8118"
	processor := tera.NewProxyProcessor(&handler)
	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	trans, err := thrift.NewTServerSocket(addr)
	if err != nil {
		panic(err)
	}
	server := thrift.NewTSimpleServer4(processor, trans, transportFactory, protocolFactory)
	log.Printf("Tera proxy start serving on %s ...", addr)
	server.Serve()
}
