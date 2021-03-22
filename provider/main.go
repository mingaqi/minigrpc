package main

import (
	"fmt"
)

func main() {
	s := "sdfsdfsfsdf呵呵"
	fmt.Println([]byte(s))
	for i, v := range []rune(s) {
		fmt.Printf("%d, %c \n", i, v)

	}

}
