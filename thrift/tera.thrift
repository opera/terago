namespace go tera
namespace java tera

enum Status {
  Ok              = 0,
  NotFound        = 1,
  Corruption      = 2,
  NotSupported    = 3,
  InvalidArgument = 4,
  TableNotExist   = 5,
  IOError         = 6
}

struct KeyValue {
  1: string key,
  2: string value,
  3: Status status,
  4: i64    ttl,   // ttl only take effect on Put&BatchPut operation
}

service TeraProxy {
  KeyValue       Get(1:string table, 2:string key)
  Status         Put(1:string table, 2:KeyValue kv)
  list<KeyValue> BatchGet(1:string table, 2:list<string> keys)
  list<Status>   BatchPut(1:string table, 2:list<KeyValue> kvs)
}
