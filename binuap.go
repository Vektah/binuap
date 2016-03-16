//go:generate wget https://raw.githubusercontent.com/ua-parser/uap-core/master/regexes.yaml
//go:generate go-bindata -o data/data.go -pkg data regexes.yaml
package binuap

import (
	"github.com/ua-parser/uap-go/uaparser"
	"github.com/vektah/binuap/data"
)

var UAP *uaparser.Parser

func init() {
	var err error
	UAP, err = uaparser.NewFromBytes(data.MustAsset("regexes.yaml"))
	if err != nil {
		panic(err)
	}
}
