package types

import (
	// "github.com/xLeDocteurx/go-opengl-playground/shaders"
	// "github.com/xLeDocteurx/go-opengl-playground/utils"
	
	// "fmt"
	// "time"

	"github.com/veandco/go-sdl2/sdl"
	// "github.com/go-gl/glfw/v3.3/glfw"
	// "github.com/go-gl/gl/v4.6-core/gl"
)

var WindowWidth int = 320
var WindowHeight int = 240

type Color struct {
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
	A uint8 `json:"a"`
}
func NewColor(r uint8, g uint8, b uint8, a uint8) Color {
	return Color{R: r, G: g, B: b, A: a}
}

type JsonMap struct {
    Width int `json:"width"`
    Height int `json:"height"`
	Walls []Wall `json:"walls"`
	Pickables []Pickable `json:"pickables"`
}
func NewJsonMap(width int, height int, walls []Wall) JsonMap {
	return JsonMap{Width: width, Height: height, Walls: walls}
}

type Wall struct {
	Color Color `json:"color"`
	Points []sdl.Point `json:"points"`
}

type Pickable struct {
	Name string `json:"name"`
	ImagePath string `json:"imagePath"`
	Hight float32 `json:"hight"`
	X int32 `json:"x"`
	Y int32 `json:"y"`
	Scale float32 `json:"scale"`
}

type SquareShape struct {
	Renderer *sdl.Renderer
	Texture *sdl.Texture
	X int32
	Y int32
	Width int32
	Height int32

	Color Color
}
func NewSquareShape(renderer *sdl.Renderer, texture *sdl.Texture, x int32, y int32, cellWidth int32, cellHeight int32, color Color) SquareShape {
	return SquareShape{Renderer: renderer, Texture: texture, X: x, Y: y, Width: cellWidth, Height: cellHeight, Color: color}
}
func (s *SquareShape) Draw() {
	rectangle := sdl.Rect{s.X, s.Y, s.Width, s.Height}
	s.Renderer.SetDrawColor(s.Color.R, s.Color.G, s.Color.B, s.Color.A)
	s.Renderer.FillRect(&rectangle)
	s.Renderer.DrawRect(&rectangle)
}

type LineShape struct {
	Renderer *sdl.Renderer
	Texture *sdl.Texture
	XA int32
	YA int32
	XB int32
	YB int32

	Color Color
}
func NewLineShape(renderer *sdl.Renderer, texture *sdl.Texture, xa int32, ya int32, xb int32, yb int32, color Color) LineShape {

	return LineShape{Renderer: renderer, Texture: texture, XA: xa, YA: ya, XB: xb, YB: yb, Color: color}
}
func (s *LineShape) Draw() {
	s.Renderer.SetDrawColor(s.Color.R, s.Color.G, s.Color.B, s.Color.A)
	s.Renderer.DrawLine(s.XA, s.YA, s.XB, s.YB)
}

type Player struct {
	Renderer *sdl.Renderer
	Texture *sdl.Texture
	X int32
	Y int32
	Width int32
	Height int32

	Color Color
}
func NewPlayer(renderer *sdl.Renderer, texture *sdl.Texture, x int32, y int32, cellWidth int32, cellHeight int32, color Color) Player {
	cellWidth = cellWidth
	cellHeight = cellHeight
	centeredX := x - ( cellWidth / 2 ) + (cellWidth / 2)
	centeredY := y - ( cellHeight / 2 ) + (cellHeight / 2)
	return Player{Renderer: renderer, Texture: texture, X: centeredX, Y: centeredY, Width: cellWidth, Height: cellHeight, Color: color}
}
func (p *Player) Draw() {
	rectangle := sdl.Rect{p.X - (p.Width / 2), p.Y - (p.Height / 2), p.Width, p.Height}
	p.Renderer.SetDrawColor(p.Color.R, p.Color.G, p.Color.B, p.Color.A)
	p.Renderer.DrawRect(&rectangle)

	p.Renderer.SetDrawColor(127, 127, 127, 255)
	p.Renderer.DrawLine(p.X, p.Y, p.X, p.Y - (p.Height / 2))
	p.Renderer.DrawLine(p.X, p.Y, p.X + 25, p.Y - p.Height)
	p.Renderer.DrawLine(p.X, p.Y, p.X - 25, p.Y - p.Height)

	lineOfSightXA := p.X + 25
	lineOfSightYA := p.Y - p.Height
	lineOfSightXB := p.X - 25
	lineOfSightYB := p.Y - p.Height

	p.Renderer.SetDrawColor(255, 127, 0, 255)
	p.Renderer.DrawLine(lineOfSightXA, lineOfSightYA, lineOfSightXB, lineOfSightYB)

	// for i := 0; i < WindowWidth; i++ {
	// 	// lineOfSightLength := lineOfSightXB - lineOfSightXA
	// 	// originX := p.X
	// 	// originY := p.Y
	// 	// directionX := lineOfSightXA + (int32(i) * lineOfSightLength / int32(WindowWidth))
	// 	// directionY := lineOfSightYA

	// 	didHitSomething := false
	// 	for !didHitSomething {
	// 		didHitSomething = true
	// 	}

	// 	time.Sleep(1 * time.Second)
	// }
}