package main

import (
	"fmt"

	"github.com/emanuela/go-programming/code_samples/sect29-chp202-ex2/quote"
	"github.com/emanuela/go-programming/code_samples/sect29-chp202-ex2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
