package types

import (
	// "github.com/xLeDocteurx/go-opengl-playground/utils"
	
	// "fmt"
	// "time"
	"math"

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
	Player Player `json:"player"`
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
	Hight float64 `json:"hight"`
	X int `json:"x"`
	Y int `json:"y"`
	Scale float64 `json:"scale"`
}

type SquareShape struct {
	Renderer *sdl.Renderer
	Texture *sdl.Texture
	X int
	Y int
	Width int
	Height int

	Color Color
}
func NewSquareShape(renderer *sdl.Renderer, texture *sdl.Texture, x int, y int, cellWidth int, cellHeight int, color Color) SquareShape {
	return SquareShape{Renderer: renderer, Texture: texture, X: x, Y: y, Width: cellWidth, Height: cellHeight, Color: color}
}
func (s *SquareShape) Draw() {
	rectangle := sdl.Rect{int32(s.X), int32(s.Y), int32(s.Width), int32(s.Height)}
	s.Renderer.SetDrawColor(s.Color.R, s.Color.G, s.Color.B, s.Color.A)
	s.Renderer.FillRect(&rectangle)
	s.Renderer.DrawRect(&rectangle)
}

type LineShape struct {
	Renderer *sdl.Renderer
	Texture *sdl.Texture
	XA int
	YA int
	XB int
	YB int

	Color Color
}
func NewLineShape(renderer *sdl.Renderer, texture *sdl.Texture, xa int, ya int, xb int, yb int, color Color) LineShape {

	return LineShape{Renderer: renderer, Texture: texture, XA: xa, YA: ya, XB: xb, YB: yb, Color: color}
}
func (s *LineShape) Draw() {
	s.Renderer.SetDrawColor(s.Color.R, s.Color.G, s.Color.B, s.Color.A)
	s.Renderer.DrawLine(int32(s.XA), int32(s.YA), int32(s.XB), int32(s.YB))
}

type Player struct {
	Renderer *sdl.Renderer
	Texture *sdl.Texture
	
	X int `json:"x"`
	Y int `json:"y"`
	Angle float64 `json:"angle"`

	Width int
	Height int

	Color Color
}
func NewPlayer(renderer *sdl.Renderer, texture *sdl.Texture, x int, y int, angle float64, cellWidth int, cellHeight int, color Color) Player {
	cellWidth = cellWidth
	cellHeight = cellHeight
	centeredX := x - ( cellWidth / 2 ) + (cellWidth / 2)
	centeredY := y - ( cellHeight / 2 ) + (cellHeight / 2)
	return Player{Renderer: renderer, Texture: texture, X: centeredX, Y: centeredY, Angle: angle, Width: cellWidth, Height: cellHeight, Color: color}
}
func (p *Player) Draw() {
	rectangle := sdl.Rect{int32(p.X) - (int32(p.Width) / 2), int32(p.Y) - (int32(p.Height) / 2), int32(p.Width), int32(p.Height)}
	p.Renderer.SetDrawColor(p.Color.R, p.Color.G, p.Color.B, p.Color.A)
	p.Renderer.DrawRect(&rectangle)

	lineOfSightXA := p.X + 30
	lineOfSightYA := p.Y - 30
	lineOfSightXB := p.X - 30
	lineOfSightYB := p.Y - 30

	// Line of sigth
	playerEndX := float64(p.X) + 500 * math.Cos(p.Angle)
	playerEndY := float64(p.Y) - 500 * math.Sin(p.Angle)

	p.Renderer.SetDrawColor(127, 127, 127, 255)
	p.Renderer.DrawLine(int32(p.X), int32(p.Y), int32(playerEndX), int32(playerEndY))

	p.Renderer.SetDrawColor(127, 127, 127, 255)
	p.Renderer.DrawLine(int32(p.X), int32(p.Y), int32(lineOfSightXA), int32(lineOfSightYA))
	p.Renderer.DrawLine(int32(p.X), int32(p.Y), int32(lineOfSightXB), int32(lineOfSightYB))

	p.Renderer.SetDrawColor(255, 127, 0, 255)
	p.Renderer.DrawLine(int32(lineOfSightXA), int32(lineOfSightYA), int32(lineOfSightXB), int32(lineOfSightYB))
}
func (p *Player) Move(upState, leftState, downState, rightState bool) {
	if upState {
		p.Y -= 5
	}
	if leftState {
		p.X -= 5
	}
	if downState {
		p.Y += 5
	}
	if rightState {
		p.X += 5
	}
}