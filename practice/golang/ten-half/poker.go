package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Poker struct {
	Number int
	Type   int
}

func (p *Poker) GetName() string {
	var name = strconv.Itoa(p.Number)

	var flower = ""

	switch p.Type {
	case 1:
		flower = "黑桃"
	case 2:
		flower = "紅心"
	case 3:
		flower = "方塊"
	case 4:
		flower = "梅花"
	default:
		flower = "錯誤"
	}

	return flower + name
}

func NewPeck() []*Poker {
	var peck []*Poker

	for i := MinPokerType; i <= MaxPokerType; i++ {
		for j := MinPokerNumber; j <= MaxPokerNumber; j++ {
			peck = append(peck, &Poker{
				Number: j,
				Type:   i,
			})
		}
	}

	return peck
}

func RandomPeck(peck []*Poker) []*Poker {
	rand.Shuffle(len(peck), func(i, j int) { peck[i], peck[j] = peck[j], peck[i] })
	return peck
}

func DealForPlayer(peck []*Poker, player *Player) []*Poker {
	if len(peck) == 0 {
		log.Fatalln("老闆，沒牌啦！")
	}
	poker := peck[0]
	peck = peck[1:]

	player.Peck = append(player.Peck, poker)
	player.CountTotal()

	return peck
}

func PrintPeck(peck []*Poker) {
	for i, p := range peck {
		if i%13 == 0 && i != 0 {
			fmt.Println("")
		}
		fmt.Printf("%s ", p.GetName())
	}
	fmt.Println()
}
