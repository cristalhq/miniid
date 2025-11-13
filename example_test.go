package miniid_test

import (
	"fmt"

	"github.com/cristalhq/miniid"
)

func Example() {
	const num = uint64(1234)
	const long = uint64(405160095143)
	const str1 = "Ju"
	const str2 = "078FWT6F"

	s := miniid.Encode(num)
	d1 := miniid.Decode(str1)
	fmt.Printf("num: %d -> %s -> %d\n", num, s, d1)

	f := miniid.EncodeFixed(long, 8)
	d2 := miniid.Decode(str2)
	fmt.Printf("fix: %d -> %s -> %d\n", long, f, d2)

	gen := miniid.New(1200)

	for range 10 {
		fmt.Print(gen.Next(), " ")
	}
	fmt.Println()

	// Output:
	// num: 1234 -> Ju -> 1234
	// fix: 405160095143 -> 078FWT6F -> 405160095143
	// JN JO JP JQ JR JS JT JU JV JW
}
