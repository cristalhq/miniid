package miniid_test

import (
	"fmt"

	"github.com/cristalhq/miniid"
)

func Example() {
	gen := miniid.New(1200)

	for i := 0; i < 10; i++ {
		fmt.Print(gen.Next(), " ")
	}
	fmt.Println()

	// Output:
	// JN JO JP JQ JR JS JT JU JV JW
}
