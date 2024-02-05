package main

import "fmt"

type Player struct {
	Name     string
	Peck     []*Poker
	Total    float32
	IsDead   bool
	IsWin    bool
	BetMoney int
	Money    int
}

func (p *Player) AddMoney(m int) {
	p.Money += m
}

func (p *Player) CountTotal() {
	p.Total = 0

	for _, poker := range p.Peck {
		var number = float32(poker.Number)
		if poker.Number > 10 {
			number = 0.5
		}
		p.Total += number
	}

	if p.Total > MaxPoint {
		p.IsDead = true
		return
	} else if p.Total == MaxPoint {
		p.IsWin = true
		return
	}

	if len(p.Peck) >= 5 {
		p.IsWin = true
	}
}

func (p *Player) PrintCurrentInfo() {
	fmt.Printf("玩家名稱：%s\n玩家餘額：%d\n玩家投注金額：%d\n玩家手牌：", p.Name, p.Money, p.BetMoney)
	PrintPeck(p.Peck)
	fmt.Printf("玩家總點數：%.1f\n", p.Total)
	if p.IsDead {
		fmt.Println("*** 總數超過 10.5 點，爆炸ㄌ！")
	}
	if p.IsWin {
		fmt.Println("*** 玩家獲勝！")
	}
}
