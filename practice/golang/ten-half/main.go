package main

import (
	"fmt"
)

func main() {
	main2()
}

func main2() {
	var rounds = []int{1000, 5000, 10000, 50000, 100000}

	for _, times := range rounds {
		var game = NewMockGame()

		game.SetStopTotal(7)
		game.SetTimes(times)
		game.SetIsPlayFive(true)

		game.RunAndPrint()
	}
}

func main1() {
	//var playerList = []string{"玩家A", "玩家B", "玩家C", "玩家D", "玩家E", "玩家F"}
	var playerList = []string{"玩家A", "玩家B"}
	var totalFiveCount = 0
	var tenHalfCount = 0
	const times = 5
	const round = 1000

	for j := 1; j <= times; j++ {
		var game = NewGame()

		for _, p := range playerList {
			game.IntoPlayer(p)
		}

		game.SetBanker(len(game.Players) - 1)

		for _, player := range game.Players {
			player.AddMoney(1000000)
		}

		for i := 1; i <= round; i++ {
			game.ResetGame()

			for index := range game.Players {
				game.BetOrder(index, 100)
			}

			game.DealPoker()
		}

		fmt.Println("======================================")
		fmt.Printf("第 %d 次跑 %d 場結果.....\n", j, round)
		//game.PrintPlayerMoney()
		totalFiveCount += game.CountFive
		tenHalfCount += game.CountTenHalf
		fmt.Printf("連五張的機率: %.2f%%\n", float64(game.CountFive)*100/float64(round*len(playerList)))
		fmt.Printf("10JQK牌型機率: %.2f%%\n", float64(game.CountTenHalf)*100/float64(round*len(playerList)))
	}

	fmt.Println("---")
	fmt.Printf("全部連五張的機率: %.2f%%\n", float64(totalFiveCount)*100/float64(times*round*len(playerList)))
	fmt.Printf("全部10JQK牌型機率: %.2f%%\n", float64(tenHalfCount)*100/float64(times*round*len(playerList)))
}
