package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	group := NewGroup(16, 0)
	rooms := NewRooms([]*Room{
		//NewRoom(1, 100, 50, 30, 5),
		//NewRoom(2, 120, 20, 40, 6),
		//NewRoom(3, 500, 60, 60, 8),

		//NewRoom(4, 1000, 200, 100, 4),
		//NewRoom(5, 0, 500, 500, 4),
		//NewRoom(6, 500, 300, 200, 4),

		NewRoom(7, 500, 500, 300, 4),
		NewRoom(8, 500, 500, 300, 4),
		NewRoom(9, 0, 500, 300, 8),
		NewRoom(10, 500, 1000, 600, 2),
	})

	group.printInfo()
	rooms.printInfo()

	if ok, err := rooms.validateGroup(group); !ok {
		fmt.Println("無法分配房間，原因:", err)
		return
	} else {
		fmt.Println("通過驗證，可以分配房間。")
	}

	rooms.assignRoom(group)
	rooms.printResult()
}
