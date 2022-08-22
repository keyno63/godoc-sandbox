package main

import (
	"awesomeProject/pkg/mod"
	"fmt"
)

func main() {
	e := mod.NewExample(1, "example1")
	fmt.Println("created example is: ", e)
	fmt.Println("created example ID is: ", e.Id())
	fmt.Println("created example Name is: ", e.Name())
}
