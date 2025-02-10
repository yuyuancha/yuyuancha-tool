package main

import "fmt"

// Room 房間
type Room struct {
	RoomID       int
	RoomPrice    int
	AdultPrice   int
	ChildPrice   int
	Capacity     int
	CurrentAdult int
	CurrentChild int
}

// NewRoom 新增房間
func NewRoom(roomID, roomPrice, adultPrice, childPrice, capacity int) *Room {
	r := &Room{}
	r.RoomID = roomID
	r.RoomPrice = roomPrice
	r.AdultPrice = adultPrice
	r.ChildPrice = childPrice
	r.Capacity = capacity
	r.CurrentAdult = 0
	r.CurrentChild = 0
	return r
}

// 是否滿人
func (r *Room) isFull() bool {
	return r.CurrentAdult+r.CurrentChild >= r.Capacity
}

// 計算價格
func (r *Room) calculatePrice() int {
	if r.CurrentAdult == 0 && r.CurrentChild == 0 {
		return 0
	}
	return r.RoomPrice + r.CurrentAdult*r.AdultPrice + r.CurrentChild*r.ChildPrice
}

// 計算群組價格
func (r *Room) calculateGroupPrice(adultNumber, childNumber int) (price, remainAdult, remainChild int) {
	if childNumber > 0 && adultNumber == 0 {
		return 0, adultNumber, childNumber
	}

	if r.isFull() {
		return 0, adultNumber, childNumber
	}

	remainAdult = adultNumber
	remainChild = childNumber

	if r.isFull() {
		return 0, remainAdult, remainChild
	}

	// 小孩比較便宜
	if r.AdultPrice > r.ChildPrice {
		// 先確定至少有一個成人
		if r.CurrentAdult == 0 {
			r.addAdult()
			remainAdult--
		}

		for remainChild > 0 && !r.isFull() {
			r.addChild()
			remainChild--
		}
	}

	return r.calculatePrice(), remainAdult, remainChild
}

// 清空房間
func (r *Room) clear() {
	r.CurrentAdult = 0
	r.CurrentChild = 0
}

// 新增大人
func (r *Room) addAdult() {
	if r.isFull() {
		return
	}
	r.CurrentAdult++
}

// 新增小孩
func (r *Room) addChild() {
	if r.isFull() {
		return
	}
	r.CurrentChild++
}

// 打印房間資訊
func (r *Room) printInfo() {
	fmt.Printf("[房號: %d] 開房價格: %d, 大人價格(每位): %d, 小孩價格(每位): %d, 總容納人數: %d, 目前大人人數: %d, 目前小孩人數: %d, 價格: %d\n",
		r.RoomID, r.RoomPrice, r.AdultPrice, r.ChildPrice, r.Capacity, r.CurrentAdult, r.CurrentChild, r.calculatePrice())
}
