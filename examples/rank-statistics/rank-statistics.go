package main

import (
	"fmt"

	gopoker "github.com/Josanshuo/go-poker"
)

const (
	randomTimes = 10000
)

func main() {
	randoms := new([randomTimes]string)
	results := map[string]int{
		gopoker.RankString(10):   0,
		gopoker.RankString(166):  0,
		gopoker.RankString(322):  0,
		gopoker.RankString(1599): 0,
		gopoker.RankString(1609): 0,
		gopoker.RankString(2467): 0,
		gopoker.RankString(3325): 0,
		gopoker.RankString(6185): 0,
		gopoker.RankString(7462): 0,
	}

	for i := 0; i < randomTimes; i++ {
		deck := gopoker.NewDeck()
		deck.Draw(18)
		hand := deck.Draw(5)
		rank := gopoker.Evaluate(hand)
		randoms[i] = gopoker.RankString(rank)
		results[randoms[i]]++
	}

	fmt.Println(results)
}
