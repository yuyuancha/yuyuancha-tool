package main

import (
	"fmt"
	"github.com/yuyuancha-tool/pokemon-card-tool/config"
	"github.com/yuyuancha-tool/pokemon-card-tool/logic"
	"github.com/yuyuancha-tool/pokemon-card-tool/model"
)

func main() {
	pokemonCardLogic := logic.NewPokemonCardLogic()
	results := pokemonCardLogic.GetUncollectedCards()

	fmt.Println("未收集卡牌：", results)

	if len(results) == 0 {
		fmt.Println("未收集卡牌數量為 0")
		return
	}

	var cardModel *model.Card
	cards, err := cardModel.GetCardsByNumbers(config.PokemonCard.UncollectedCardSeries, results)
	if err != nil {
		panic(err)
	}
	if len(cards) == 0 {
		fmt.Println("查詢卡牌數量為空")
		return
	}

	for _, card := range cards {
		s := card.GetCardString()
		fmt.Println(s)
	}
}
