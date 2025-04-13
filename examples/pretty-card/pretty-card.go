package main

import (
	"fmt"
	"os"
	"strings"

	gopoker "github.com/Josanshuo/go-poker"
)

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func main() {
	args := os.Args[1:]

	var hand []gopoker.Card
	for _, c := range args {
		hand = append(hand, gopoker.NewCard(strings.ToUpper(c)))
	}

	outputStrings := Map(hand, func(item gopoker.Card) string { return fmt.Sprintf(item.PrettyString()) })
	fmt.Println(outputStrings)
	fmt.Println(gopoker.NewCard("Ah").PrettyString())
}
