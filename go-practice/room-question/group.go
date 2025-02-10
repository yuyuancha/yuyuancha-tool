package main

import "fmt"

// Group 組別
type Group struct {
	AdultNumber        int
	ChildNumber        int
	CurrentAdultNumber int
	CurrentChildNumber int
}

// NewGroup 新增組別
func NewGroup(adultNumber, childNumber int) *Group {
	g := &Group{}
	g.AdultNumber = adultNumber
	g.ChildNumber = childNumber
	g.CurrentAdultNumber = adultNumber
	g.CurrentChildNumber = childNumber
	return g
}

// 取得總人數
func (g *Group) getAllNumber() int {
	return g.AdultNumber + g.ChildNumber
}

// 打印組別資訊
func (g *Group) printInfo() {
	fmt.Printf("大人數量:%d, 小孩數量:%d\n", g.AdultNumber, g.ChildNumber)
}
