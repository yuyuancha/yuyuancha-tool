package main

import (
	"fmt"
)

type Game struct {
	Peck            []*Poker
	BankerIndex     int
	Players         []*Player
	CurrentMaxTotal float32
	CountFive       int
	CountTenHalf    int
}

func NewGame() *Game {
	return &Game{
		Peck:         RandomPeck(NewPeck()),
		CountFive:    0,
		CountTenHalf: 0,
	}
}

func (g *Game) IntoPlayer(name string) {
	//fmt.Printf("%s 進入遊戲...\n", name)
	g.Players = append(g.Players, &Player{Name: name})
}

func (g *Game) SetBanker(index int) {
	if index > len(g.Players)-1 {
		index = len(g.Players) - 1
	}

	g.BankerIndex = index
}

func (g *Game) ResetGame() {
	for _, player := range g.Players {
		player.BetMoney = 0
		player.Peck = []*Poker{}
		player.Total = 0
		player.IsDead = false
		player.IsWin = false
	}

	g.Peck = RandomPeck(NewPeck())
	g.CurrentMaxTotal = 0
}

func (g *Game) PrintGameInfo() {
	//fmt.Println("當前遊戲資訊：")
	for index, player := range g.Players {
		fmt.Println("-------------------------")
		if index == g.BankerIndex {
			fmt.Println("【莊家】")
		}
		player.PrintCurrentInfo()
	}
}

func (g *Game) PrintPlayerMoney() {
	for index, player := range g.Players {
		if index == g.BankerIndex {
			fmt.Printf("【莊家】")
		}
		fmt.Printf("%s 餘額：%d\n", player.Name, player.Money)
	}
}

func (g *Game) BetOrder(index, betMoney int) {
	if index == g.BankerIndex || g.Players[index].Money <= 0 {
		return
	}
	g.Players[index].Money -= betMoney
	g.Players[index].BetMoney = betMoney
}

func (g *Game) DealPoker() {
	for index := range g.Players {
		if g.Players[index].BetMoney == 0 && index != g.BankerIndex {
			continue
		}

		if index == g.BankerIndex {
			// 莊家先檢查是否全部爆炸
			var isAllDead = true
			for i, p := range g.Players {
				if p.IsDead == false && i != index {
					isAllDead = false
				}
			}

			if isAllDead {
				continue
			}
		}

		for !g.Players[index].IsDead && !g.Players[index].IsWin {
			if g.Players[index].Total > 6 {
				if index != g.BankerIndex {
					// 非莊家
					break
				} else {
					if g.Players[index].Total > g.CurrentMaxTotal {
						// 莊家額外判斷是否有大於最大點數
						break
					}
				}
			}
			g.Peck = DealForPlayer(g.Peck, g.Players[index])
		}

		var number = float32(g.Players[index].Peck[0].Number)
		if number > 10 {
			number = 0.5
		}
		var maxTotal = g.Players[index].Total - number

		if !g.Players[index].IsDead && maxTotal > g.CurrentMaxTotal {
			g.CurrentMaxTotal = maxTotal
		}

		if !g.Players[index].IsDead && g.Players[index].Total >= 5 {
			g.CountFive += 1
		}

		if g.Players[index].IsWin && len(g.Players[index].Peck) == 2 {
			// 10.5 雙倍獎金
			g.CountTenHalf++
		}
	}

	g.Settle()
}

func (g *Game) Settle() {
	var bankPoint = g.Players[g.BankerIndex].Total
	for index, player := range g.Players {
		if player.BetMoney == 0 || index == g.BankerIndex {
			// 沒投注或莊家直接跳過
			continue
		}

		if player.IsDead {
			// 輸了賠錢
			g.TransferBetMoneyToBanker(index)
			continue
		}

		if bankPoint > 10.5 {
			// 莊家爆炸，直接支付場上玩家獎金
			var winMoney = player.BetMoney

			// 10+0.5 牌型的雙倍獎勵
			if len(player.Peck) == 2 && player.Total == 10.5 {
				winMoney *= 2
			}

			g.TransferFromBanker(index, winMoney)

			continue
		}

		var winMoney = 0

		if !player.IsWin {
			if player.Total < bankPoint {
				// 輸了賠錢
				g.TransferBetMoneyToBanker(index)
				continue
			} else if player.Total == bankPoint && player.Total != 10.5 {
				// 平手，雙倍牌型需要加倍，這裡要額外判斷
				g.TransferFromBanker(index, 0)
				continue
			}
			player.IsWin = true
			winMoney = player.BetMoney
		} else {
			winMoney = player.BetMoney
		}

		// 10+0.5 牌型的雙倍獎勵
		if len(player.Peck) == 2 && player.Total == 10.5 {
			winMoney *= 2
		}

		g.TransferFromBanker(index, winMoney)
	}
}

func (g *Game) TransferFromBanker(index, money int) {
	// 先收回原本的投注
	g.Players[index].Money += g.Players[index].BetMoney

	// 獎金再從莊家支付
	g.Players[g.BankerIndex].Money -= money
	g.Players[index].Money += money
}

// TransferBetMoneyToBanker 莊家收走投注金額
func (g *Game) TransferBetMoneyToBanker(index int) {
	g.Players[g.BankerIndex].Money += g.Players[index].BetMoney
}
