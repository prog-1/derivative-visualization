package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

type Game struct {
	width, height, cX, cY int
	tangent               float64
}

func (g *Game) Update() error {
	g.cX, g.cY = ebiten.CursorPosition()
	return nil
}

func givenFunction(x float64) float64 {
	return 100 * (math.Cos(x/100) + math.Sin(x/100))
}

func drawTangent(screen *ebiten.Image, x, y, ang, length float64, clr color.Color) {
	x2 := x + length*math.Cos(ang)
	y2 := y + length*math.Sin(ang)
	vector.StrokeLine(screen, float32(x), float32(y), float32(x2), float32(y2), 1, clr, false)
	x2r := x - length*math.Cos(ang)
	y2r := y - length*math.Sin(ang)
	vector.StrokeLine(screen, float32(x), float32(y), float32(x2r), float32(y2r), 1, clr, false)
}

func (g *Game) Draw(screen *ebiten.Image) {
	nx := 1.0
	var x float64
	for x < float64(screenWidth) {
		y := givenFunction(x) + float64(screenHeight)/2
		vector.DrawFilledCircle(screen, float32(x), float32(y), 1, color.White, false)
		if math.Abs(x-float64(g.cX)) < 1 {
			ang := math.Atan2(givenFunction(x+0.1)-givenFunction(x), 0.1)
			drawTangent(screen, x, y, ang, g.tangent, color.RGBA{255, 255, 0, 255})
		}
		x += nx
	}
}

func (g *Game) Layout(outWidth, outHeight int) (int, int) {
	return screenWidth, screenHeight
}

func NewGame(width, height int) *Game {
	return &Game{
		width:   width,
		height:  height,
		tangent: screenWidth,
	}
}

func main() {
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
