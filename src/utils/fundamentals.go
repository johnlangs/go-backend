package utils

import "fmt"

func main() {

	// booleans
	and := true && false
	fmt.Println(and)

	or := true || false
	fmt.Println(or)

	not := !true
	fmt.Println(not)
	// booleans

	// Strings
	str := "Hello\nWorld!\n"
	fmt.Println(str)
	// Strings

	// Variables
	// var name type = expression //
	var num int32 = 42
	var float float64 = 56.28
	const name string = "John"

	fmt.Println(num)
	fmt.Println(float)
	fmt.Println(name)
	// Variables

	// _ idenifier
	str = "world"
	// printf return two values (#bytes printed, error state)
	_, err := fmt.Printf("Hello %s\n", str)
	if err != nil {
		fmt.Println("Error!")
	}
	// _ idenifier

	// Container types

		// Arrays
		var a [3]int
		fmt.Println(a)
		fmt.Println(a[1])

		b := [...]int{1, 2, 3}
		fmt.Println(b)

		fmt.Println(len(b))
		fmt.Println(b[len(b)-1])
		// Arrays

		// Slices
		n := make([]int, 3)
		fmt.Println(n)
		fmt.Println(len(n))	
		// Slices

	// Container types
}