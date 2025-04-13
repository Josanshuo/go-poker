package main

import (
	"fmt"

	gopoker "github.com/Josanshuo/go-poker"
)

func main() {

	keepGoing := true

	for keepGoing {
		fmt.Println("Please look at the following 5 cards, which two hole cards could be its nuts hand (press any key to reveal the answer):")
		deck := gopoker.NewDeck()
		flop := deck.Draw(5)
		fmt.Println(flop)

		answer := ""
		fmt.Println("Type two hole cards ('quit' for exit): ")
		fmt.Scan(&answer)

		possibleHands := deck.Peek(52 - 5)
		var previousRank int32 = 10000
		bestHand := make([]gopoker.Card, 2)
		for i, h1 := range possibleHands {
			for _, h2 := range possibleHands[i+1:] {
				showdown := make([]gopoker.Card, 5)
				copy(showdown, flop)
				showdown = append(showdown, h1, h2)
				rank := gopoker.Evaluate(showdown)
				if rank < previousRank && !(rank > 10 && rank < 167) { // Exclude four of kind from the result
					bestHand[0] = h1
					bestHand[1] = h2
					previousRank = rank
				}
			}
		}

		if answer == "quit" {
			fmt.Println("Program will quit after this question.")
			keepGoing = false
		}

		fmt.Println("The two hole cards are:")
		fmt.Println(bestHand)
	}
}
