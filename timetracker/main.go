package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	// get the current process id
	pid := robotgo.GetPID()
	fmt.Println("pid----", pid)

	// get current Window Handle
	hwnd := robotgo.GetHandle()
	fmt.Println("hwnd---", hwnd)

	// get current Window Handle
	bhwnd := robotgo.GetHandle()
	fmt.Println("bhwnd---", bhwnd)

	// get current Window title
	title := robotgo.GetTitle()
	fmt.Println("title-----", title)

	// screen
	x, y := robotgo.GetMousePos()
	fmt.Println("pos:", x, y)
	color := robotgo.GetPixelColor(100, 200)
	fmt.Println("color----", color)

	// window
	robotgo.ActiveName("brave")

	isExist, err := robotgo.PidExists(100)
	if err == nil && isExist {
		fmt.Println("pid exists is", isExist)

		robotgo.Kill(100)
	}

	title = robotgo.GetTitle(100)
	fmt.Println("title@@@", title)
}
