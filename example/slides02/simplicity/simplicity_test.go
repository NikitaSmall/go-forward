package simplicity

// the testing package is a part of standard library.
// it's unnecessary to use any other package (but it's helpful)
// basic asserts and comparasments is just `for`s and `if`s
import "testing"

// test functions start with Test* and have a single argument
func TestDefaultValue(t *testing.T) {
	// start with the test
	count := Counter{
		desc: "test counter",
	}

	// basic check
	if count.String() != "test counter: 0" {
		// rise a test error if something bad or unexpected happened in test
		t.Error("unexpected result after simple initialization")
	}

	count.increment()
	if count.String() != "test counter: 1" {
		t.Error("unexpected result after increment")
	}
}

// func TestToFail(t *testing.T) {
// 	t.Error("just a fail to see it works")
// }
