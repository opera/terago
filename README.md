# terago
go client for [baidu/tera](https://github.com/baidu/tera)

## Run test

 * If there is a tera environment(sdk & server), run `go test --tags prod`
 * Else run `go test --mock`

## Prepare

 * Setup an active tera cluster or onebox. ([How?](https://github.com/baidu/tera/blob/master/doc/en/onebox.md))
 * Install tera develop environment. We need `tera_c.h`, `libtera_c.so`.
