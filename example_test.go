package norm_test

import (
	"fmt"

	"github.com/goloop/norm"
)

func ExampleEmail() {
	value, ok := norm.Email("  Example @ Gmail.com ")
	fmt.Println(value, ok)
	// Output: Example@gmail.com true
}

func ExampleIBAN() {
	value, ok := norm.IBAN("de89 3704 0044 0532 0130 00")
	fmt.Println(value, ok)
	// Output: DE89370400440532013000 true
}

func ExampleMAC() {
	value, ok := norm.MAC("0011.2233.4455")
	fmt.Println(value, ok)
	// Output: 00:11:22:33:44:55 true
}

func ExampleRemove() {
	// Strip zero-width and control characters that break validation.
	clean := norm.Remove("\u200bhi\ufeff", norm.Invisible)
	fmt.Printf("%q\n", clean)
	// Output: "hi"
}

func ExampleKeep() {
	fmt.Println(norm.Keep("a1-b2 c3!", norm.Letters|norm.Digits))
	// Output: a1b2c3
}
