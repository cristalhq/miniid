package miniid_test

import (
	"fmt"

	"github.com/cristalhq/miniid"
)

func Example() {
	s := miniid.Encode(1234)
	fmt.Printf("num: %d -> %s\n", 1234, s)

	f := miniid.EncodeFixed(1234, 8)
	fmt.Printf("fix: %d -> %s\n", 1234, f)

	gen := miniid.New(1200)

	for range 10 {
		fmt.Print(gen.Next(), " ")
	}
	fmt.Println()

	// Output:
	// num: 1234 -> Ju
	// fix: 1234 -> 000000Ju
	// JN JO JP JQ JR JS JT JU JV JW
}
