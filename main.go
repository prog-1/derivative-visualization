package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

const (
	screenWidth  = 960
	screenHeight = 600
	eps          = 0.000001
)

type game struct {
	f       func(float64) float64
	NewPlot func() *plot.Plot
	p       *plot.Plot
}

func NewGame() *game {

	return &game{
		func(x float64) float64 { return math.Tan(x) },
		func() *plot.Plot {
			p := plot.New()
			p.X.Min = -10
			p.X.Max = 10
			p.Y.Min = -10
			p.Y.Max = 10

			p.BackgroundColor = color.Black
			return p
		},
		nil,
	}
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	g.p = g.NewPlot()

	f := plotter.NewFunction(g.f)
	f.Color = color.RGBA{255, 255, 255, 255}
	g.p.Add(f)

	t := plotter.NewFunction(GetTangent(g.f, GetCursorX(g.p)))
	t.Color = color.RGBA{255, 255, 255, 255}
	g.p.Add(t)

	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	DrawPlot(screen, g.p)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("X: %v", GetCursorX(g.p)))
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Tangent plotter")
	g := NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func DrawPlot(screen *ebiten.Image, p *plot.Plot) {
	// https://github.com/gonum/plot/wiki/Drawing-to-an-Image-or-Writer:-How-to-save-a-plot-to-an-image.Image-or-an-io.Writer,-not-a-file.
	img := image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight))
	c := vgimg.NewWith(vgimg.UseImage(img))
	p.Draw(draw.New(c))

	screen.DrawImage(ebiten.NewImageFromImage(c.Image()), &ebiten.DrawImageOptions{})
}

func GetCursorX(p *plot.Plot) float64 {
	tmpx, _ := ebiten.CursorPosition()
	x := float64(tmpx)
	return x/(screenWidth/(p.X.Max-p.X.Min)) + p.X.Min
}

func GetTangent(f func(float64) float64, x0 float64) func(float64) float64 {
	// TODO:
	// 1. Get cursor's X coordinate +
	// 2. Draw an according tangent
	// 2.1. Calculate f(x) +
	// 2.2. Calculate f'(x)
	// 2.3. Calculate b
	// 2.5. Draw the line
	x1 := x0 + eps
	y0, y1 := f(x0), f(x1)
	k := (y1 - y0) / (x1 - x0)
	b := y0 - k*x0
	return func(x float64) float64 {
		return k*x + b
	}
}
