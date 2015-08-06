package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
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

	str := "df 世 界asd"
	s := "1界"
	index := strings.IndexAny(str, "界")
	fmt.Printf("%d......\n", index)
	for i, c := range str {
		fmt.Println(i)
		for _, m := range s {
			fmt.Printf("%c\t%c\n", c, m)
		}
	}

	for i := len(str); i > 0; {
		rune, size := utf8.DecodeLastRuneInString(str[0:i])
		fmt.Printf("%v\t", i)
		i -= size
		for _, m := range s {
			fmt.Printf("%c\t%v\t%v\n", rune, size, m)
		}
	}

	fmt.Println("bytes =", len(str))
	fmt.Println("runes =", utf8.RuneCountInString(str))

	for _, ss := range strings.Fields(str) {
		fmt.Println(ss)

	}
	fmt.Println(strings.Title(str))

}
