package main

import (
	"fmt"
)

func main() {

	var Input1 int
	fmt.Print("Input:")
	fmt.Scanf("%d", &Input1)
	for i := 1; i <= Input1; i++ {

		// if i%i == 0 && i/i == i {
		// 		fmt.Println("Square")
		// 	}
		if square(i) {
			fmt.Println("Square")
		} else if cubic(i) {
			fmt.Println("Cube")
		} else if cubic(i) && square(i) {
			fmt.Println("SquareCube")
		} else {
			fmt.Println(i)
		}

	}
}

func square(num int) bool {
	for i := 1; i*i <= num; i++ {
		if num%i == 0 && num/i == i {
			return true
		}
	}
	return false
}

func cubic(num int) bool {
	for i := 1; i*i*i <= num; i++ {
		if num%i == 0 && num/i == i*i {
			return true
		}
	}
	return false
}
