ARCHIVED: https://github.com/ua-parser/uap-core now does this by default.

binuap
======

This is a trivial wrapper around https://github.com/ua-parser/uap-core
that includes the regexes.yaml database as go source code. This removes
the runtime dependency on regexes.yaml.


Usage
=====
```go
import "github.com/vektah/binuap"

func main() {
	ua := UAP.Parse("Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_3; en-us; Silk/1.1.0-80) AppleWebKit/533.16 (KHTML, like Gecko) Version/5.0 Safari/533.16 Silk-Accelerated=true")

	fmt.Println(res.Device.Family)
}
```

See the docs on https://github.com/ua-parser/uap-go/tree/master/uaparser for
details on `ua` itself.


Updating the data
=================

```bash
# Make sure you have go-bindata
go get -u github.com/jteeuwen/go-bindata/...

# Then generate the new data file
go generate

# And make sure you run the tests. It will validate all the regexes
go test
```
