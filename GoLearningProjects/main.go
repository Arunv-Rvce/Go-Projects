package main

import "fmt"

func main() {
	fmt.Println(addition(10, 20))
	switch_examples()
	get_example()

	foo := map[string]interface{}{
		"Matt": 6,
	}
	timeMap(foo)
	fmt.Println(foo)

	s := &fakeString{"Fakestring for testing"}
	printString(s)
	printString("Actuall String Type")
}
