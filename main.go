package main

import (
	"image"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

const (
	screenWidth  = 960
	screenHeight = 600
)

type game struct {
	p *plot.Plot
}

func NewGame() *game {
	p := plot.New()
	p.X.Min = -10
	p.X.Max = 10
	p.Y.Min = -10
	p.Y.Max = 10

	graph := plotter.NewFunction(func(x float64) float64 { return math.Sin(x) })
	graph.Color = color.RGBA{0, 0, 255, 255}
	p.Add(graph)

	return &game{
		p,
	}
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	DrawPlot(screen, g.p)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Derivative plotter")
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
