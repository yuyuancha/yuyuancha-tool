package main

import (
	"errors"
	"fmt"
	"sort"
)

// Rooms 房間清單
type Rooms struct {
	Rooms []*Room
}

// NewRooms 新增房間清單
func NewRooms(room []*Room) *Rooms {
	r := &Rooms{}
	r.Rooms = room
	return r
}

// 驗證是否能分配
func (r *Rooms) validateGroup(group *Group) (bool, error) {
	// 判斷是否有小孩但沒有大人
	if group.ChildNumber > 0 && group.AdultNumber == 0 {
		return false, errors.New("小孩必須有大人陪同")
	}

	// 判斷是否總數大於房間總容量
	if r.getAllCapacity() < group.getAllNumber() {
		return false, errors.New("人數超過房間總容量")
	}

	// 判斷是否每間房間的小孩都有至少一個大人
	if group.ChildNumber > 0 && group.AdultNumber < r.getRoomLength() {
		r.sortByCapacity()
		amount := 0
		for i, room := range r.Rooms {
			if i+1 > group.AdultNumber {
				break
			}
			amount += room.Capacity
		}
		if amount < group.getAllNumber() {
			return false, errors.New("每間房間至少需要一個大人")
		}
	}

	return true, nil
}

// 分配房間
func (r *Rooms) assignRoom(group *Group) {
	type roomCalStruct struct {
		roomIndex int
		roomPrice int
		priceDiff int // 越大越適合給大人住
		tmpPrice  int // 暫時價格
	}

	adultNum := group.AdultNumber
	childNum := group.ChildNumber

	roomCals := make([]*roomCalStruct, 0)
	for index, room := range r.Rooms {
		roomCals = append(roomCals, &roomCalStruct{
			roomIndex: index,
			roomPrice: room.RoomPrice,
			priceDiff: room.AdultPrice - room.ChildPrice,
		})
	}

	clearRoomCalsTmpPrice := func() {
		for _, roomCal := range roomCals {
			roomCal.tmpPrice = 0
		}
	}

	for childNum > 0 {
		// 價格適合小孩住的放前面
		sort.Slice(roomCals, func(i, j int) bool {
			if roomCals[i].priceDiff == roomCals[j].priceDiff {
				return roomCals[i].roomPrice < roomCals[j].roomPrice
			}
			return roomCals[i].priceDiff < roomCals[j].priceDiff
		})

		for i, roomCal := range roomCals {
			room := r.Rooms[roomCal.roomIndex]
			if room.isFull() {
				roomCals = append(roomCals[:i], roomCals[i+1:]...)
				continue
			}

			cn := childNum
			if cn > room.Capacity-1 {
				cn = room.Capacity - 1
			}
			roomCal.tmpPrice = room.RoomPrice + room.AdultPrice + room.ChildPrice*cn
		}

		sort.Slice(roomCals, func(i, j int) bool {
			return roomCals[i].tmpPrice < roomCals[j].tmpPrice
		})

		room := r.Rooms[roomCals[0].roomIndex]
		room.addAdult()
		cn := (roomCals[0].tmpPrice - room.RoomPrice - room.AdultPrice) / room.ChildPrice
		for i := 0; i < cn; i++ {
			room.addChild()
		}

		childNum -= cn
		adultNum--

		if childNum == 0 && !room.isFull() {
			c := room.Capacity - room.CurrentAdult - room.CurrentChild
			if c > adultNum {
				c = adultNum
			}
			for i := 0; i < c; i++ {
				room.addAdult()
				adultNum--
			}
		}

		clearRoomCalsTmpPrice()
	}

	for adultNum > 0 {
		// 價格適合大人住的放前面
		sort.Slice(roomCals, func(i, j int) bool {
			if roomCals[i].priceDiff == roomCals[j].priceDiff {
				return roomCals[i].roomPrice < roomCals[j].roomPrice
			}
			return roomCals[i].priceDiff > roomCals[j].priceDiff
		})

		for i, roomCal := range roomCals {
			room := r.Rooms[roomCal.roomIndex]
			if room.isFull() {
				roomCals = append(roomCals[:i], roomCals[i+1:]...)
				continue
			}

			an := adultNum
			remainCapacity := room.Capacity - room.CurrentAdult - room.CurrentChild
			if an > remainCapacity {
				an = remainCapacity
			}

			roomCal.tmpPrice = room.RoomPrice + room.AdultPrice*an
		}

		sort.Slice(roomCals, func(i, j int) bool {
			return roomCals[i].tmpPrice < roomCals[j].tmpPrice
		})

		room := r.Rooms[roomCals[0].roomIndex]
		an := (roomCals[0].tmpPrice - room.RoomPrice) / room.AdultPrice
		for i := 0; i < an; i++ {
			room.addAdult()
		}

		adultNum -= an
		clearRoomCalsTmpPrice()
	}

	r.printInfo()
}

func (r *Rooms) printResult() {
	total := 0
	for _, room := range r.Rooms {
		total += room.calculatePrice()
	}
	fmt.Println("總價格:", total)
}

// 取得所有房間的總容量
func (r *Rooms) getAllCapacity() int {
	total := 0
	for _, room := range r.Rooms {
		total += room.Capacity
	}
	return total
}

// 取得房間數量
func (r *Rooms) getRoomLength() int {
	return len(r.Rooms)
}

// 計算總價格
func (r *Rooms) calculateTotalPrice() int {
	total := 0
	for _, room := range r.Rooms {
		total += room.calculatePrice()
	}
	return total
}

// 依照房間容量排序
func (r *Rooms) sortByCapacity() {
	sort.Slice(r.Rooms, func(i, j int) bool {
		return r.Rooms[i].Capacity > r.Rooms[j].Capacity
	})
}

// 依照房間ID排序
func (r *Rooms) sortByRoomID() {
	sort.Slice(r.Rooms, func(i, j int) bool {
		return r.Rooms[i].RoomID < r.Rooms[j].RoomID
	})
}

// 打印房間資訊
func (r *Rooms) printInfo() {
	for _, room := range r.Rooms {
		room.printInfo()
	}
}
