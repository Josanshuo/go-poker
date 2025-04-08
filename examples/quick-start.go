package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/chehsunliu/poker"
)

func main() {
	args := os.Args[1:]

	deck := poker.NewDeck()
	var hand []poker.Card
	if len(args) > 0 {
		n, _ := strconv.Atoi(args[0])
		hand = deck.Draw(n)
	} else {
		hand = deck.Draw(5)
	}
	fmt.Println(hand)

	rank := poker.Evaluate(hand)
	fmt.Println(rank)
	fmt.Println(poker.RankString(rank))
}
