package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	fmt.Println(hand)
	fmt.Println(remainingCards)
}
