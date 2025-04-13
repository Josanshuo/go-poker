package main

import (
	"fmt"
	"os"
	"strings"

	gopoker "github.com/Josanshuo/go-poker"
)

func main() {
	args := os.Args[1:]

	var hand []gopoker.Card
	if len(args) >= 5 && len(args) <= 7 {
		for _, c := range args {
			hand = append(hand, gopoker.NewCard(strings.ToUpper(c)))
		}
	} else {
		panic("Only support 5, 6 and 7 cards.")
	}
	fmt.Println(hand)

	rank := gopoker.Evaluate(hand)
	fmt.Println(rank)
	fmt.Println(gopoker.RankString(rank))
}
