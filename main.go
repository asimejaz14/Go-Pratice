package main

import "fmt"

func main() {
	number := "2"
	floatNumber := 2.45

	// format specifiers
	// simple print
	fmt.Println("Number is", number)
	fmt.Printf("Number is %q\n", number) // prints in double quotes
	fmt.Printf("Number is %T\n", number) // prints type of variable
	fmt.Printf("Float Number is %f\n", floatNumber)
	fmt.Printf("Float Number with 1 decimal point is %.01f\n", floatNumber)

	// sprintf
	str := fmt.Sprintf("Float Number with 1 decimal point is %.01f\n", floatNumber)
	println("Saved string is:", str)

	// Arrays
	var array1 [3]int = [3]int{3, 6, 1}
	var array2 = [3]int{2, 8, 4}
	array3 := [3]int{0, 9, 5}

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)

	fmt.Println("Length of array 1", len(array1))
}
