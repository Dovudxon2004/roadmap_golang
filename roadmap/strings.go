package roadmap

import (
	"fmt"
	"unicode/utf8"
)

func Strings() {
	s := "ýllo"
	fmt.Println(s[0])
	fmt.Println(utf8.RuneCountInString(s))
	var r string = "⌘"
	fmt.Printf("%x", r)
	for i, r := range s {
		fmt.Println(i, r)
	}
}
