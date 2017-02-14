# terago
Golang client for [baidu/tera](https://github.com/baidu/tera)

## Run unit test

 * `go test -tags prod` If there is a tera environment(sdk & server)
 * `go test -tags mock`
 
## Prepare

 * Setup an active tera cluster or onebox. ([How?](https://github.com/baidu/tera/blob/master/doc/en/onebox.md))
 * Install tera develop environment. We need `tera_c.h`, `libtera_c.so`.
