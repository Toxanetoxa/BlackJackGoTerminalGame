package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	name      string
	value     int
	value2    int
	suite     string
	suiteIcon rune
}

type Suite struct {
	name string
	icon rune
}

func main() {
	fmt.Println("Начало игры")

	var cards []Card

	player, casino := make([]Card, 0), make([]Card, 0)
	playerCash, casinoCash, count := 0, 0, 8

	var suits []Suite = []Suite{
		{name: "Diamonds", icon: '♦'},
		{name: "Hearts", icon: '♥'},
		{name: "Clubs", icon: '♣'},
		{name: "Spades", icon: '♠'},
	}

	for _, s := range suits {
		for _, c := range generateCardsDeck(s) {
			cards = append(cards, c)
		}
	}

	deck := shuffleCards(generateDesk(count, cards))

	for len(deck) > 4 {
		fmt.Print("----------------Новый раунд----------------\n")
		player, deck = dealingCards(player, deck, 2)
		playerCash = countSum(player)
		fmt.Println(playerCash)
		var more string = "no"
		fmt.Print("Eщё карту? yes/no\n")
		_, err := fmt.Scan(&more)
		if err != nil {
			fmt.Println("Error occurred while scanning input:", err)
			return
		}
		for more != "no" && more != "n" {
			player, deck = dealingCards(player, deck, 1)
			playerCash = countSum(player)
			fmt.Println(playerCash)
			if playerCash > 21 {
				fmt.Println("!!!!!!!!!!!!Казино выйграло!!!!!!!!!!!!")
				break
			}
			fmt.Print("Eщё карту? yes/no\n")
			_, err := fmt.Scan(&more)
			if err != nil {
				fmt.Println("Error occurred while scanning input:", err)
				return
			}
		}

		casino, deck = dealingCards(casino, deck, 2)
		casinoCash = countSum(casino)
		fmt.Print("Казино набрало карты\n")
		fmt.Println(casinoCash)
		for casinoCash < 21 {
			if playerCash <= 21 {
				fmt.Print("Казино набирает ещё картy\n")
				casino, deck = dealingCards(casino, deck, 1)
				casinoCash = countSum(casino)
				fmt.Println(playerCash)
			}
		}

		switch {
		case playerCash > 21:
			break
		case casinoCash > playerCash && casinoCash < 22:
			fmt.Println("!!!!!!!!!!!!Казино выйграло!!!!!!!!!!!!")
		case casinoCash > 21 || playerCash > casinoCash && playerCash < 21:
			fmt.Println("!!!!!!!!!!!!Игрок выйграл!!!!!!!!!!!!!!")
		default:
			fmt.Println("!--------------Ничья--------------!")
		}

		playerCash, casinoCash = 0, 0

		playerCash, casinoCash = 0, 0
		player, casino = make([]Card, 0), make([]Card, 0)
	}

	if len(deck) < 4 {
		fmt.Println("Игра закончена")
	}

}

func generateCardsDeck(s Suite) []Card {
	return []Card{
		{name: "Ace", value: 11, value2: 1, suite: s.name, suiteIcon: s.icon},
		{name: "2", value: 2, value2: 0, suite: s.name, suiteIcon: s.icon},
		{name: "3", value: 3, value2: 0, suite: s.name, suiteIcon: s.icon},
		{name: "4", value: 4, value2: 0, suite: s.name, suiteIcon: s.icon},
		{name: "5", value: 5, value2: 0, suite: s.name, suiteIcon: s.icon},
		{name: "6", value: 6, value2: 0, suite: s.name, suiteIcon: s.icon},
		{name: "7", value: 7, value2: 0, suite: s.name, suiteIcon: s.icon},
		{name: "8", value: 8, value2: 0, suite: s.name, suiteIcon: s.icon},
		{name: "9", value: 9, value2: 0, suite: s.name, suiteIcon: s.icon},
		{name: "10", value: 10, value2: 0, suite: s.name, suiteIcon: s.icon},
		{name: "Jack", value: 10, value2: 0, suite: s.name, suiteIcon: s.icon},
		{name: "Queen", value: 10, value2: 0, suite: s.name, suiteIcon: s.icon},
		{name: "King", value: 10, value2: 0, suite: s.name, suiteIcon: s.icon},
	}
}

func generateDesk(c int, cards []Card) []Card {
	var cardsDeck []Card

	for i := 0; i < c; i++ {
		for _, c := range cards {
			cardsDeck = append(cardsDeck, c)
		}
	}

	return cardsDeck
}

func shuffleCards(cards []Card) []Card {
	rand.Seed(time.Now().UnixNano())

	for i := len(cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}

	return cards
}

func dealingCards(player []Card, deck []Card, count int) ([]Card, []Card) {
	for i := 0; i < count; i++ {
		player = append(player, (deck)[i])
	}
	deck = (deck)[count:]

	for _, c := range player {
		fmt.Printf("%s%c ", c.name, c.suiteIcon)
	}
	fmt.Printf("\n")

	return player, deck
}

func countSum(cards []Card) int {
	sum := 0
	for _, i := range cards {
		sum += i.value
	}
	return sum
}

//func startRound(deck *[]Card) {
//	player, casino := make([]Card, 0), make([]Card, 0)
//	playerCash, casinoCash := 0, 0
//	dealingCards(deck, &player, 2)
//	playerCash = countSum(player)
//
//}
