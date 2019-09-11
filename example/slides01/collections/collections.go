package collections

import "fmt"

func demo() {
	// classic array with fixed size (size can't be changed)
	var array [3]int = [3]int{1, 2, 3}

	fmt.Println(len(array)) // can check size

	// can access array
	array[1] = 42
	fmt.Println(array[1])

	// just a declaration, no initialization
	var slice []int // slice is, basically, array without fixed size

	// empty slice initialization
	slice = make([]int, 0) // or make([]int, 0, 0)

	for index, element := range array { // can also iterate over any collection
		// can access both index (key for map) and value
		fmt.Printf("move %d'th array element to slice", index)
		slice = append(slice, element) // in case the addition is required
	}

	fmt.Println(slice[1]) // in other terms - same as array

	// declaration and initialization, just a shortcut
	simpleMap := make(map[string]int)

	simpleMap["question"] = -1
	simpleMap["answer"] = 42

	// access directly
	fmt.Println(simpleMap["answer"])

	// the way to check key
	if value, keyExists := simpleMap["answer"]; keyExists {
		fmt.Printf("the key exists and holds following value: %d", value)
	}

	// the way to delete by key
	delete(simpleMap, "question")

	fmt.Println(array, slice, simpleMap)
}
