package word

import (
	"fmt"
	"testing"

	"github.com/emanuela/go-programming/code_samples/sect29-chp202-ex2/quote"
)

func TestUseCount(t *testing.T) {
	m := UseCount("one two three three three")
	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("expected", 1, "got", v)
			}
		case "two":
			if v != 1 {
				t.Error("expected", 1, "got", v)
			}
		case "three":
			if v != 3 {
				t.Error("expected", 3, "got", v)
			}
		}
	}
}

func TestCount(t *testing.T) {
	wordCount := Count("We ask ourselves, Who am I to be brilliant, gorgeous, talented, fabulous?")
	if wordCount != 12 {
		t.Error("Expected", 12, "But got", wordCount)
	}
}

func ExampleCount() {
	fmt.Println(Count("We ask ourselves, Who am I to be brilliant, gorgeous, talented, fabulous?"))
	// Output:
	// 12
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
