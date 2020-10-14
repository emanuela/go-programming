package main

import (
	"fmt"

	"github.com/emanuela/go-programming/code_samples/sect27-chp192-ex1/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(5),
	}
	fmt.Println(fido)
}
