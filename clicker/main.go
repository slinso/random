package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func main() {
	robotgo.MouseSleep = 10
	execute := true

	clicks := 8001

	hook.Register(hook.KeyUp, nil, func(e hook.Event) {
		if hook.RawcodetoKeychar(uint16(e.Rawcode)) == "x" {
			fmt.Println("press x")
			execute = false
		}
	})
	s := hook.Start()
	hook.Process(s)

	for execute {
		numClick := 100
		if clicks < 1000 {
			numClick = 10
		} else if clicks == 10000 {
			numClick = 500
		}

		numUpgrades := 0
		if clicks < 100 {
			numUpgrades = 10
		} else if clicks < 1000 {
			numUpgrades = 6
		} else if clicks < 10000 {
			numUpgrades = 3
		} else {
			numUpgrades = 1
		}

		// middle of the game
		robotgo.Move(5400, 800)
		for i := 0; i < numClick; i++ {
			robotgo.Click("left", false)
			if clicks < 10000 {
				clicks++
			}
		}

		// upgrades
		robotgo.Move(6200, 500)
		for i := 0; i < numUpgrades; i++ {
			robotgo.Click("left", false)
		}

		// 2x bonus
		robotgo.Move(4850, 620)
		for i := 0; i < 3; i++ {
			robotgo.Click("left", false)
		}
	}
}
