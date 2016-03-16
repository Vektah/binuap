package binuap

import "testing"

// The test suite on the underlying library is really good, we
// just want to make sure that everything is working.
func TestThatDatabaseLoadsAndParses(t *testing.T) {
	res := UAP.Parse("Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_3; en-us; Silk/1.1.0-80) AppleWebKit/533.16 (KHTML, like Gecko) Version/5.0 Safari/533.16 Silk-Accelerated=true")

	if res.Device.Family != "Kindle" {
		t.Error("Parse error")
	}
}
