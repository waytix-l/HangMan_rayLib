package main

import (
	ge "game/gameEngine"
)

func main() {
	var g ge.GameEngine
	var m ge.Menu
	g.InitGameEngine(720, 480, "HangMan RayLib")
	g.RunningGameEngine(&m)
}