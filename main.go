package main

import (
	"fmt"
	"strings"
)

func HashTest(sep string) (uint32, uint32) {
	var pow, sq uint32 = 1, 2
	for i := len(sep); i > 0; i >>= 1 {

		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return pow, sq
}

func main() {
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))

	fmt.Println(strings.ContainsAny("team", "e"))
	fmt.Println(strings.ContainsAny("failure", "u & i"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))

	fmt.Println(1 & 1)
	fmt.Println(2 & 1)
	fmt.Println(4 & 1)
	fmt.Println(5 & 1)

	pow, sq := HashTest("1234567")
	fmt.Println(pow)
	fmt.Println(sq)

}
