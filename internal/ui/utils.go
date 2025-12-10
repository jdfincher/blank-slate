// Package ui
package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	OFFSET_Y = 25
)

func CenterX(width int) (X float32) {
	return float32((rl.GetScreenWidth() / 2) - (width / 2))
}

func CenterY(height int) (Y float32) {
	return float32((rl.GetScreenHeight() / 2) - (height / 2))
}

func BottomAlign(height int) (Y float32) {
	return float32((rl.GetScreenHeight() - height) - OFFSET_Y)
}

func CenterTextX(text string, fontsize int) (X int32) {
	width := int(rl.MeasureText(text, int32(fontsize)))
	return int32((rl.GetScreenWidth() / 2) - (width / 2))
}

func CenterWithinPanelX(xPanel, wPanel, wRec float32) float32 {
	return xPanel + (wPanel / 2) - (wRec / 2)
}
