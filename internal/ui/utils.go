// Package ui
package ui

import (
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	OFFSET_Y = 25
)

func CenterX(width float32) (X float32) {
	return float32((rl.GetScreenWidth() / 2) - int(width/2))
}

func CenterY(height float32) (Y float32) {
	return float32((rl.GetScreenHeight() / 2) - int(height/2))
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

func LabelAlignLeft() {
	rg.SetStyle(rg.LABEL, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_LEFT))
}

func LabelAlignRight() {
	rg.SetStyle(rg.LABEL, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_RIGHT))
}

func LabelAlignCenter() {
	// Only aligns text within labels defined below
	rg.SetStyle(rg.LABEL, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_CENTER))
}

func SizeText(size int64) {
	rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, size)
}

func ResetText32() {
	rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 32)
}
