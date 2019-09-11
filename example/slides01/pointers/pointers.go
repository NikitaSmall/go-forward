package pointers

import "fmt"

func main() {
	i := 1 // value
	fmt.Println("initial:", i)

	zeroval(i) //pass as value
	fmt.Println("zeroval:", i)

	zeroptr(&i)                // pass by pointer
	fmt.Println("zeroptr:", i) // should be changed

	fmt.Println("pointer:", &i) // value pointer
}

func zeroval(ival int) {
	ival = 0 // operate with copied and passed value
}

func zeroptr(iptr *int) {
	*iptr = 0 // operate with original value, accessible through the pointer
}
