package main

import "fmt";


func main() {
	//Arrays (arrays in go are fixed length. you cannot change the length)
	var ages [4]int = [4]int{3, 5, 7, 6}

	var new_ages = [6]int{3, 5, 7, 6, 2, 2}

	names := [4]string{"allen", "eben", "ethereum", "nana"}

	fmt.Println(new_ages, len(new_ages)) 
	fmt.Println(ages, len(ages))
	fmt.Println(names)

	//Slices
	var new_scores = []int{60, 67, 45}
	new_scores[2] = 100
	
	fmt.Println(new_scores, len(new_scores))

	new_scores = append(new_scores, 90)

	fmt.Println(new_scores, len(new_scores))
}
