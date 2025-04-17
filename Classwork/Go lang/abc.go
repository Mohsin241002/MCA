package main

import "fmt"

func main() {
	// arr := [5]int{1, 2, 3, 4, 5}
	// fmt.Println(arr)
	// for _, v := range arr {
	// 	fmt.Println(v)
	// }

	var myMap map[string]int = map[string]int{
		"pear": 4,
	}

	myMap["apple"] = 2
	myMap["banana"] = 3
	myMap["orange"] = 5

	// print the map
	fmt.Println(myMap)
	delete(myMap, "apple")
	fmt.Println(myMap)

	// check if the key exists
	val, ok := myMap["apple"]
	fmt.Println("Apple:", val, "Present:", ok)
	// output: Apple: 0 Present: false
}
