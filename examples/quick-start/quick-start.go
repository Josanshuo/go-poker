package main

import (
	"fmt"
	"os"
	"strconv"

	gopoker "github.com/Josanshuo/go-poker"
)

func main() {
	args := os.Args[1:]

	deck := gopoker.NewDeck()
	var hand []gopoker.Card
	if len(args) > 0 {
		n, _ := strconv.Atoi(args[0])
		hand = deck.Draw(n)
	} else {
		hand = deck.Draw(5)
	}
	fmt.Println(hand)

	rank := gopoker.Evaluate(hand)
	fmt.Println(rank)
	fmt.Println(gopoker.RankString(rank))
}
