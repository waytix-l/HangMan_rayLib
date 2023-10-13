package main

import (
	ge "game/gameEngine"
)

func main() {
	var g ge.GameEngine
	var m ge.Menu
	g.InitGameEngine(1920, 1080, "HangMan RayLib")
	g.RunningGameEngine(&m)
}