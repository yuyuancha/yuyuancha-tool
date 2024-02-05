package main

import "fmt"

type MockGame struct {
	peck   []*Poker
	player *Player
	banker *Player

	times      int     // 遊玩局數
	stopTotal  float32 // 玩家牌大於停牌點數
	isPlayFive bool    // 玩家是否判斷過五關

	countFive      int // 統計過五關次數
	countTenHalf   int // 統計10JQK次數
	countFiveAndTH int // 	統計過五關十點半次數
	countWin       int // 統計玩家獲勝次數
	countTie       int // 統計玩家平手次數
	countLose      int // 統計玩家賠錢次數

	betBalance int // 玩家餘額(單位: 押注金額)
}

func NewMockGame() *MockGame {
	return &MockGame{
		peck: RandomPeck(NewPeck()),
		player: &Player{
			Name:  "test",
			Money: 1000000000,
		},
		banker: &Player{
			Name:  "banker",
			Money: 1000000000,
		},
		times:          0,
		stopTotal:      0,
		isPlayFive:     false,
		countFive:      0,
		countTenHalf:   0,
		countFiveAndTH: 0,
		countWin:       0,
		countTie:       0,
		countLose:      0,
	}
}

// SetStopTotal 設定玩家停止補牌點數
func (g *MockGame) SetStopTotal(stopTotal float32) {
	g.stopTotal = stopTotal
}

// SetTimes 設定遊玩局數
func (g *MockGame) SetTimes(times int) {
	g.times = times
}

// SetIsPlayFive 設定玩家是否判斷過五關(預設: false)
func (g *MockGame) SetIsPlayFive(isPlay bool) {
	g.isPlayFive = isPlay
}

// Run 開始遊玩
func (g *MockGame) Run() {
	for i := 1; i <= g.times; i++ {
		g.resetGame()
		g.dealPoker()
		g.countStatistics()
	}

	// 最後計算玩家餘額
	g.countBalance()
}

// RunAndPrint 開始遊玩並打印結果
func (g *MockGame) RunAndPrint() {
	g.Run()
	g.PrintStatistics()
}

// PrintStatistics 打印統計結果
func (g *MockGame) PrintStatistics() {
	fmt.Println("--------------------------------------------------------")
	fmt.Printf("               👀 跑%d場結果 ⚡️\n", g.times)
	fmt.Println("--------------------------------------------------------")
	fmt.Printf(" 🎲【連五張】機率: %.2f%% (%d倍獎金)\n", FormatToPercent(g.countFive, g.times), fiveOdds)
	fmt.Printf(" 🎲【10JQK】機率: %.2f%% (%d倍獎金)\n", FormatToPercent(g.countTenHalf, g.times), tenHalfOdds)
	fmt.Printf(" 🎲【連五張且總和十點半】機率: %.2f%% (%d倍獎金)\n", FormatToPercent(g.countFiveAndTH, g.times), fiveAndTHOdds)
	fmt.Printf(" 🎲【玩家-勝】機率: %.2f%%\n", FormatToPercent(g.countWin, g.times))
	fmt.Printf(" 🎲【玩家-平】機率: %.2f%%\n", FormatToPercent(g.countTie, g.times))
	fmt.Printf(" 🎲【玩家-輸】機率: %.2f%%\n", FormatToPercent(g.countLose, g.times))
	fmt.Printf(" 💰 玩家最終餘額: %dx (設x=投注金額) \n", g.betBalance)
}

func (g *MockGame) resetGame() {
	g.player.Peck = []*Poker{}
	g.player.Total = 0
	g.player.IsDead = false
	g.player.IsWin = false

	g.banker.Peck = []*Poker{}
	g.banker.Total = 0
	g.banker.IsDead = false
	g.banker.IsWin = false

	g.peck = RandomPeck(NewPeck())
}

func (g *MockGame) dealPoker() {
	g.dealPokerToPlayer()
	g.dealPokerToBanker()
}

func (g *MockGame) dealPokerToPlayer() {
	var player = g.player

	for !player.IsDead && !player.IsWin {
		var total = player.Total
		var peckNum = len(player.Peck)

		if total > 6 && total > g.stopTotal {
			// 玩家點數必須大於6，且大於設定停止補牌點數，才能過補
			if !g.isPlayFive {
				// 不需要拼過五關，直接過補
				break
			}

			if peckNum <= 3 {
				break
			}
		}

		g.peck = DealForPlayer(g.peck, player)
	}
}

func (g *MockGame) dealPokerToBanker() {
	var banker = g.banker
	var playerOpenTotal = g.getPlayerOpenTotal()

	if g.player.IsDead || g.player.IsWin {
		return
	}

	for !banker.IsDead && !banker.IsWin {
		var total = banker.Total
		if total > 6 && total > playerOpenTotal {
			// 莊家點數必須大於6，且大於玩家明牌點數，才能過補
			break
		}

		g.peck = DealForPlayer(g.peck, banker)
	}
}

func (g *MockGame) getPlayerOpenTotal() float32 {
	var playerFirstNumber = float32(g.player.Peck[0].Number)
	if playerFirstNumber > 10 {
		playerFirstNumber = 0.5
	}
	return g.player.Total - playerFirstNumber
}

func (g *MockGame) countStatistics() {
	var player = g.player

	var playerTotal = player.Total
	var playerPeckNum = len(player.Peck)

	// 過五關統計
	if playerPeckNum == 5 {
		if playerTotal == MaxPoint {
			// 過五關十點半統計
			g.countFiveAndTH++
		} else if playerTotal < MaxPoint {
			g.countFive++
		}
	}

	// 10JQK 統計
	if playerTotal == MaxPoint && playerPeckNum == 2 {
		g.countTenHalf++
	}

	// 玩家輸贏
	g.countWinLose()
}

func (g *MockGame) countWinLose() {
	var player = g.player
	var banker = g.banker
	var bankerTotal = banker.Total
	var playerTotal = player.Total

	if player.IsWin {
		g.countWin++
		return
	}

	if player.IsDead {
		g.countLose++
		return
	}

	if banker.IsDead {
		g.countWin++
		return
	}

	if playerTotal > bankerTotal {
		g.countWin++
		return
	}
	if playerTotal == bankerTotal {
		g.countTie++
		return
	}

	g.countLose++
}

func (g *MockGame) countBalance() {
	// 總共花費
	var totalBet = g.times
	// 總共獲勝場次(不包含特殊倍數場次)
	var normalWinTimes = g.countWin - g.countFive - g.countTenHalf - g.countFiveAndTH

	// 普通倍數
	// 總共獲勝獎金(包含賭金)
	var totalNormalWin = 2 * normalWinTimes

	// 特殊倍數
	// 總共過五關十點半獎金(包含賭金)
	var totalFiveAndTH = g.countFiveAndTH*fiveAndTHOdds + g.countFiveAndTH
	// 總共過五關獎金(包含賭金)
	var totalFiveWin = g.countFive*fiveOdds + g.countFive
	// 總共10JQK獎金(包含賭金)
	var totalTenHalfWin = g.countTenHalf*tenHalfOdds + g.countTenHalf

	g.betBalance = 0 - totalBet + totalNormalWin + totalFiveWin + totalTenHalfWin + totalFiveAndTH
}
