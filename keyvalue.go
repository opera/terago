package terago

type KeyValue struct {
	Key   string
	Value string
	TTL   int
	Err   error
}
