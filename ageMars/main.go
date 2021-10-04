package main

import "fmt"

func main() {
	var age int

	fmt.Print("input age: ")
	fmt.Scanln(&age)

	mars := age * 365 / 687

	fmt.Println("umur kamu di mars ", mars)
}
