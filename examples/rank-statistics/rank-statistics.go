package main

import (
	"fmt"

	"github.com/chehsunliu/poker"
)

const (
	randomTimes = 10000
)

func main() {
	randoms := new([randomTimes]string)
	results := map[string]int{
		poker.RankString(10):   0,
		poker.RankString(166):  0,
		poker.RankString(322):  0,
		poker.RankString(1599): 0,
		poker.RankString(1609): 0,
		poker.RankString(2467): 0,
		poker.RankString(3325): 0,
		poker.RankString(6185): 0,
		poker.RankString(7462): 0,
	}

	for i := 0; i < randomTimes; i++ {
		deck := poker.NewDeck()
		deck.Draw(18)
		hand := deck.Draw(5)
		rank := poker.Evaluate(hand)
		randoms[i] = poker.RankString(rank)
		results[randoms[i]]++
	}

	fmt.Println(results)
}
