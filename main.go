package main

import (
	"math"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	screenWidth := int32(1920)
	screenHeight := int32(1080)

	rl.InitWindow(screenWidth, screenHeight, "MONSQUAZ")
	rl.SetTargetFPS(60)

	render := Arend{
		SWidth:  screenWidth,
		SHeight: screenHeight,
	}

	render.Init()

	for !rl.WindowShouldClose() {
		render.Update()

		render.Draw()
	}

	rl.CloseWindow()
}

type Arend struct {
	SWidth  int32
	SHeight int32

	StartTime time.Time
	LastFrame time.Time

	ArrowTexture rl.Texture2D
	Background   rl.Texture2D

	Arrows []Arrow
}

type Arrow struct {
	x     int
	y     int
	speed int
}

func (a *Arend) Init() {
	greenArrow := rl.LoadImage("green.png")
	a.ArrowTexture = rl.LoadTextureFromImage(greenArrow)
	rl.UnloadImage(greenArrow)

	gradients := a.GenGradient()
	a.Background = rl.LoadTextureFromImage(gradients)
	a.StartTime = time.Now()

	for i := 0; i < 10; i++ {
		a.GenerateArrow()
	}

}

func (a *Arend) GenerateArrow() {
	a.Arrows = append(a.Arrows, Arrow{
		x:     -int(a.ArrowTexture.Width) + rand.Intn(int(a.SWidth)),
		y:     int(a.ArrowTexture.Height) + int(a.SHeight) + rand.Intn(500),
		speed: 1 + rand.Intn(10),
	})
}

func (a *Arend) Update() {
	// t := time.Now()
	// elapsed := t.Sub(a.LastFrame)
	// start := t.Sub(a.StartTime)
	a.LastFrame = time.Now()

	for i, ar := range a.Arrows {
		a.Arrows[i].y = ar.y - ar.speed

	}

	for i, ar := range a.Arrows {
		if ar.y < (0 - int(a.ArrowTexture.Height)) {
			a.Arrows = append(a.Arrows[:i], a.Arrows[i+1:]...)

			if len(a.Arrows) < 30 {
				for m := 0; m < 1+rand.Intn(5); m++ {
					a.GenerateArrow()
				}
			}
		}
	}

}

func (a *Arend) Draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)

	rl.DrawTexture(a.Background, 0, 0, rl.White)

	for _, ar := range a.Arrows {
		rl.DrawTexture(a.ArrowTexture, int32(ar.x), int32(ar.y), rl.White)
	}

	rl.EndDrawing()
}

type GC struct {
	From rl.Color
	To   rl.Color
}

func (a *Arend) GenGradient() *rl.Image {
	w := int(a.SWidth)
	h := int(a.SHeight)
	m := []GC{
		GC{
			From: rl.Color{R: 0, G: 22, B: 7},
			To:   rl.Color{R: 0, G: 12, B: 5},
		},
		GC{
			From: rl.Color{R: 0, G: 49, B: 20},
			To:   rl.Color{R: 0, G: 51, B: 21},
		},
		GC{
			From: rl.Color{R: 0, G: 64, B: 30},
			To:   rl.Color{R: 0, G: 67, B: 30},
		},
		GC{
			From: rl.Color{R: 0, G: 80, B: 35},
			To:   rl.Color{R: 0, G: 84, B: 37},
		},
		GC{
			From: rl.Color{R: 0, G: 99, B: 45},
			To:   rl.Color{R: 0, G: 100, B: 46},
		},
		GC{
			From: rl.Color{R: 0, G: 118, B: 56},
			To:   rl.Color{R: 0, G: 112, B: 57},
		},
		GC{
			From: rl.Color{R: 0, G: 134, B: 63},
			To:   rl.Color{R: 0, G: 143, B: 66},
		},
		GC{
			From: rl.Color{R: 0, G: 154, B: 74},
			To:   rl.Color{R: 0, G: 165, B: 77},
		},
		GC{
			From: rl.Color{R: 0, G: 180, B: 87},
			To:   rl.Color{R: 0, G: 186, B: 88},
		},
		GC{
			From: rl.Color{R: 0, G: 201, B: 96},
			To:   rl.Color{R: 0, G: 203, B: 95},
		},
		GC{
			From: rl.Color{R: 0, G: 213, B: 102},
			To:   rl.Color{R: 0, G: 220, B: 102},
		},
		GC{
			From: rl.Color{R: 0, G: 235, B: 112},
			To:   rl.Color{R: 0, G: 240, B: 115},
		},
		GC{
			From: rl.Color{R: 0, G: 253, B: 123},
			To:   rl.Color{R: 0, G: 251, B: 121},
		},
		GC{
			From: rl.Color{R: 9, G: 255, B: 131},
			To:   rl.Color{R: 15, G: 254, B: 134},
		},
		GC{
			From: rl.Color{R: 28, G: 255, B: 140},
			To:   rl.Color{R: 30, G: 253, B: 139},
		},
		GC{
			From: rl.Color{R: 45, G: 254, B: 147},
			To:   rl.Color{R: 48, G: 252, B: 148},
		},
		GC{
			From: rl.Color{R: 62, G: 253, B: 157},
			To:   rl.Color{R: 60, G: 253, B: 156},
		},
		GC{
			From: rl.Color{R: 82, G: 253, B: 164},
			To:   rl.Color{R: 82, G: 253, B: 162},
		},
		GC{
			From: rl.Color{R: 95, G: 253, B: 179},
			To:   rl.Color{R: 102, G: 253, B: 175},
		},
		GC{
			From: rl.Color{R: 115, G: 253, B: 184},
			To:   rl.Color{R: 130, G: 253, B: 191},
		},
		GC{
			From: rl.Color{R: 132, G: 253, B: 192},
			To:   rl.Color{R: 134, G: 253, B: 192},
		},
		GC{
			From: rl.Color{R: 140, G: 255, B: 197},
			To:   rl.Color{R: 149, G: 254, B: 200},
		},
		GC{
			From: rl.Color{R: 165, G: 255, B: 207},
			To:   rl.Color{R: 170, G: 254, B: 209},
		},
	}

	for i := len(m)/2 - 1; i >= 0; i-- {
		opp := len(m) - 1 - i
		m[i], m[opp] = m[opp], m[i]
	}

	im := rl.GenImageColor(w, h, rl.Color{R: 255, G: 255, B: 255, A: 255})
	gn := len(m)
	hh := int(math.Ceil(1080 / float64(gn)))
	for i, g := range m {
		g.From.A = 255
		g.To.A = 255
		rl.ImageDraw(
			im,
			rl.GenImageGradientV(w, hh, g.From, g.To),
			rl.NewRectangle(0, 0, float32(w), float32(hh)),
			rl.NewRectangle(0, float32(hh*i), float32(w), float32(hh)),
			rl.White,
		)
	}

	return im
}
