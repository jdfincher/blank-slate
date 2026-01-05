package main

import (
	"fmt"
	"runtime"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jdfincher/blank-slate/internal/gamestate"
	"github.com/jdfincher/blank-slate/internal/player"
	"github.com/jdfincher/blank-slate/internal/ui"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

func DebugOverlay(gS *gamestate.State) {
	var (
		pos = rl.GetMousePosition()

		mouseX  = int32(pos.X)
		mouseY  = int32(pos.Y)
		Yoffset = 125
	)
	rl.DrawRectangle(int32(rl.GetScreenWidth()-150), 0, 150, 150, rl.Black)
	rl.DrawFPS(int32(rl.GetScreenWidth()-Yoffset), 25)
	rl.DrawText(DebuggerState(gS), int32(rl.GetScreenWidth()-Yoffset), 50, 18, rl.DarkGreen)
	rl.DrawText(DebuggerCreateState(gS), int32(rl.GetScreenWidth()-Yoffset), 75, 18, rl.DarkGreen)
	rl.DrawText(fmt.Sprintf("Mouse-X: %d\nMouse-Y: %d", mouseX, mouseY), int32(rl.GetScreenWidth()-Yoffset), 100, 18, rl.DarkGreen)
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

func DebuggerCreateState(gS *gamestate.State) string {
	if gS.CreateState.Name {
		return "Name"
	} else if gS.CreateState.Race {
		return "Race"
	} else if gS.CreateState.Type {
		return "Type"
	} else if gS.CreateState.Skill {
		return "Skills"
	} else if gS.CreateState.Method {
		return "Methods"
	} else if gS.CreateState.Inv {
		return "Inventory"
	}
	return "No CreateState Found"
}

func MouseView() {
	rl.DrawLine(rl.GetMouseX(), 0, rl.GetMouseX(), int32(rl.GetScreenHeight()), rl.Red)
	rl.DrawLine(0, rl.GetMouseY(), int32(rl.GetScreenWidth()), rl.GetMouseY(), rl.Red)
}

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	gS := gamestate.State{
		Menu:        true,
		Create:      false,
		Idle:        false,
		Combat:      false,
		Quit:        false,
		Player:      new(player.Player),
		CreateState: new(gamestate.CreateState),
	}
	gS.CreateState.Reset()

	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.SetConfigFlags(rl.FlagWindowUndecorated)
	rl.InitWindow(screenWidth, screenHeight, "Test")
	image := rl.LoadImage("internal/res/temp/BGMenu.jpg")

	rl.ImageResize(image, 1920, 1080)

	bg := rl.LoadTextureFromImage(image)
	defer rl.UnloadTexture(bg)
	defer rl.UnloadImage(image)
	defer rl.CloseWindow()

	rg.LoadStyle("internal/styles/cyber.rgs")
	rl.SetTargetFPS(60)

	debug := false

	for !rl.WindowShouldClose() {

		virtmouse := rl.GetMousePosition()

		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			rl.SetMousePosition(int(virtmouse.X), int(virtmouse.Y))
		}

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
		} else if gS.Quit {
			break
		}
		rw := rl.GetRenderWidth()
		rh := rl.GetRenderHeight()
		sw := rl.GetScreenWidth()
		sh := rl.GetScreenHeight()

		rl.DrawText(
			fmt.Sprintf("Render: x%d / y%d Screen: x%d / y%d", rw, rh, sw, sh), 20, 20, 20, rl.Yellow)

		if debug {
			MouseView()
			DebugOverlay(&gS)

		}

		if pressed := rl.IsKeyPressed(rl.KeyLeftControl); pressed {
			debug = !debug
		}

		rl.EndDrawing()
	}
}
