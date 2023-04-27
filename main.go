package main

import (
	"fmt"

	"github.com/pedroernestomolina/testmod/quotes"
	"github.com/pedroernestomolina/testmod/strings"
)

func main() {
	p, i := strings.CountEvenOdd("hola pedro, que tal estas?")
	fmt.Println(p, i)

	emoji := quotes.GetEmoji("turtle")
	fmt.Println(emoji)
}
