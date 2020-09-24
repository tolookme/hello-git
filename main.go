package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello git!")

	fmt.Println("")

	fmt.Println("git status")
	fmt.Println("git log")

	fmt.Println("")

	fmt.Println("git add main.go")
	fmt.Println("git commit -m \"add main.go\"")
	fmt.Println("git push")

	fmt.Println("")

	fmt.Println("git commit -m \"modify main.go\"")
	fmt.Println("git push")

	fmt.Printf("time = %v.\n", time.Now())

}
