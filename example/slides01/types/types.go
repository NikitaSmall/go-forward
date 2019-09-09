package types

import "log"

var (
	number      int     = 42  // also int8, int16, int32, int64
	floatNumber float32 = 4.2 // or float64

	str       string = "hello!"
	byteSlice []byte = []byte("it's not moon!")

	truth bool = true
	// and more of them! (chan, pointer, array etc.)
)

func types() {
	var test bool // each value has default value

	log.Print(test)
}
