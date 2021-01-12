package main

import "fmt"

func main() {
	a := [...]float64{67.7, 89.8, 21, 78}
	// fmt.Println("length of a is ", len(a))
	// for i := 0; i < len(a); i++ { // lopping from 0 to the length of the array
	// 	fmt.Printf("%d the element of a is %.2f\n", i, a[i])
	// }

	sum := float64(0)
	for i, v := range a {
		fmt.Printf("%d the element of a is %.2f\n", i, a[i])
		sum += v
	}
	fmt.Printf("\nsum of all elements of a ", sum)
}
