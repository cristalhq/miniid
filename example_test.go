package miniid_test

import (
	"fmt"

	"github.com/cristalhq/miniid"
)

func Example() {
	s := miniid.Encode(1234)
	fmt.Printf("num: %d -> %s\n", 1234, s)

	gen := miniid.New(1200)

	for i := 0; i < 10; i++ {
		fmt.Print(gen.Next(), " ")
	}
	fmt.Println()

	// Output:
	// num: 1234 -> Ju
	// JN JO JP JQ JR JS JT JU JV JW
}
