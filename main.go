package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

var (
	cursorX, cursorY int
)

func update(screen *ebiten.Image) error {

	//Cursor position
	cursorX, cursorY = ebiten.CursorPosition()

	//Image to draw on
	image := ebiten.NewImage(screenWidth, screenHeight)

	//Function
	f := func(x float64) float64 {
		return 50 * math.Sin(x/50)
		//return math.Pow(2, x*0.05)
	}

	// Function draw
	drawFunction(image, f)

	// Tangent draw
	drawTangent(image, f, float64(cursorX), float64(cursorY))

	// Draw on screen
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(image, op)

	// Cursor coordinate debug
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Cursor: (%d, %d)", cursorX, cursorY))

	return nil
}

func drawTangent(image *ebiten.Image, f func(x float64) float64, x, y float64) {

	// Slope/derivative calculation
	r := 0.001              // range
	x2 := x - screenWidth/2 // proper x to place in function
	derivative := (f(x2+r) - f(x2-r)) / (2 * r)

	// Line draw
	scaleFactor := 1000.0
	ebitenutil.DrawLine(image, x-scaleFactor, f(x2)-derivative*scaleFactor+screenHeight/2, x+scaleFactor, f(x2)+derivative*scaleFactor+screenHeight/2, color.RGBA{0, 255, 0, 255})
}

func drawFunction(image *ebiten.Image, f func(x float64) float64) {
	for x := -screenWidth / 2; x <= screenWidth/2; x++ {
		y := f(float64(x))
		image.Set(int(x)+screenWidth/2, int(y)+screenHeight/2, color.White)
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Derivative Visualization")
	if err := ebiten.RunGame(&game{}); err != nil {
		panic(err)
	}
}

type game struct{}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	if err := update(screen); err != nil {
		panic(err)
	}
}

func (g *game) Layout(width, weight int) (screenWidth, screenHeight int) {
	return width, weight
}
