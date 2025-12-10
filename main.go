package main

import (
	"fmt"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jdfincher/blank-slate/internal/gamestate"
	"github.com/jdfincher/blank-slate/internal/player"
	"github.com/jdfincher/blank-slate/internal/ui"
)

const (
	screenWidth  = 800
	screenHeight = 400
)

func DebugOverlay(gS *gamestate.State) {
	var (
		mouseX  = rl.GetMouseX()
		mouseY  = rl.GetMouseY()
		Yoffset = 125
	)
	rl.DrawRectangle(int32(rl.GetScreenWidth()-150), 0, 150, 125, rl.Black)
	rl.DrawFPS(int32(rl.GetScreenWidth()-Yoffset), 25)
	rl.DrawText(DebuggerState(gS), int32(rl.GetScreenWidth()-Yoffset), 50, 18, rl.DarkGreen)
	rl.DrawText(fmt.Sprintf("Mouse-X: %d\nMouse-Y: %d", mouseX, mouseY), int32(rl.GetScreenWidth()-Yoffset), 70, 18, rl.DarkGreen)
}

func DebuggerState(gS *gamestate.State) string {
	if gS.Menu {
		return "Menu State"
	} else if gS.Create {
		return "Create State"
	} else if gS.Idle {
		return "Idle State"
	} else if gS.Combat {
		return "Combat State"
	}
	return "No State Found"
}

func MouseView() {
	rl.DrawLine(rl.GetMouseX(), 0, rl.GetMouseX(), int32(rl.GetScreenHeight()), rl.Red)
	rl.DrawLine(0, rl.GetMouseY(), int32(rl.GetScreenWidth()), rl.GetMouseY(), rl.Red)
}

func main() {
	gS := gamestate.State{
		Menu:        true,
		Create:      false,
		Idle:        false,
		Combat:      false,
		Player:      new(player.Player),
		CreateState: new(gamestate.CreateState),
	}
	gS.CreateState.Reset()

	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(screenWidth, screenHeight, "Test")
	bg := rl.LoadTexture("internal/res/temp/BGMenu.jpg")
	defer rl.UnloadTexture(bg)
	defer rl.CloseWindow()

	rg.LoadStyle("internal/styles/cyber.rgs")
	rl.SetTargetFPS(60)

	debug := false

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		gS.HandleStyle()

		if gS.Menu {
			rl.DrawTexture(bg, 0, 0, rl.White)
			gS.StartMenu()

		} else if gS.Create {
			rl.DrawTexture(bg, 0, 0, rl.White)
			gS.NewGameCreate()
		} else if gS.Idle {
			if clicked := rg.Button(rl.NewRectangle(ui.CenterX(200), ui.BottomAlign(50), 200, 50), "Back"); clicked {
				gS.Menu = true
			}
		}

		if clicked := rg.Button(rl.NewRectangle(float32(rl.GetScreenWidth())-55, ui.BottomAlign(30), 30, 30), "#142#"); clicked {
			debug = !debug
		}
		if debug {
			MouseView()
			DebugOverlay(&gS)

		}

		if pressed := rl.IsKeyPressed(rl.KeyLeftControl); pressed {
			debug = !debug
		}
		rl.EndDrawing()
	}
	rl.UnloadTexture(bg)
}
