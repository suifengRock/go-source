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

func Replace(s, old, new string, n int) string {

	if old == new || n == 0 {
		return s // avoid allocation
	}
	// Compute number of replacements.
	if m := strings.Count(s, old); m == 0 {
		fmt.Println("...1")

		return s // avoid allocation
	} else if n < 0 || m < n {
		n = m
	}
	fmt.Println("...3")
	// Apply replacements to buffer.
	t := make([]byte, len(s)+n*(len(new)-len(old)))
	w := 0
	start := 0
	for i := 0; i < n; i++ {
		j := start
		if len(old) == 0 {
			if i > 0 {
				_, wid := utf8.DecodeRuneInString(s[start:])
				j += wid
			}
		} else {
			j += strings.Index(s[start:], old)
		}
		w += copy(t[w:], s[start:j])
		w += copy(t[w:], new)
		start = j + len(old)
	}
	w += copy(t[w:], s[start:])
	return string(t[0:w])
}

type my interface {
	SayHello()
}

func hello(m my) {
	m.SayHello()
}

type huluwa struct {
}

func (h huluwa) SayHello() {
	fmt.Println("hello huluwa !")
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

	str := "df 世 界 asd"
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
	fmt.Println(strings.ToTitle(str))

	str = "Hello, 世界"

	// for len(str) > 0 {
	// 	r, size := utf8.DecodeRuneInString(str)
	// 	fmt.Printf("%c %v\n", r, size)

	// 	str = str[size:]
	// }
	fmt.Println(Replace(str, "", "f", -1))

	hulu := new(huluwa)
	hello(hulu)
}
