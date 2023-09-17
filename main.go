package main

import "fmt";


func main() {
	//Arrays
	var ages [4]int = [4]int{3, 5, 7, 6}

	var new_ages = [6]int{3, 5, 7, 6, 2, 2}

	names := [4]string{"allen", "eben", "ethereum", "nana"}

	fmt.Println(new_ages, len(new_ages)) 
	fmt.Println(ages, len(ages))
	fmt.Println(names)

	//Slices
}
