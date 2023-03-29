package main

import (
	"github.com/xLeDocteurx/go-opengl-playground/types"
	"github.com/xLeDocteurx/go-opengl-playground/utils"
	// "github.com/xLeDocteurx/go-opengl-playground/shaders"

	"github.com/veandco/go-sdl2/sdl"
	// "github.com/tfriedel6/canvas/sdlcanvas"

	"fmt"
	// "sync"
	"time"
	"os"
	"runtime"
	"io/ioutil"
	"encoding/json"
	// "image/color"
	"math"
)

// type JsonMap struct {
//     width int
//     height int
// 	cells []int
// }

// Game State
var Player types.Player

// Inputs State
var UpState bool = false
var LeftState bool = false
var DownState bool = false
var RightState bool = false

// Global vars
var TimeFactor float64 = 0.005
var WallsFactor float64 = 0.1
var Ouverture int32 = 90
var WindowWidth int32 = 320
// var WindowWidth int32 = 640
// var WindowWidth int32 = 960
var WindowHeight int32 = 240
// var WindowHeight int32 = 480
// var WindowHeight int32 = 720
var cellWidth int32
var cellHeight int32

var framesCount int32 = 0

var window *sdl.Window
// var surface *sdl.Surface
var renderer *sdl.Renderer
var texture *sdl.Texture

// var vertexShaders []uint32
// var fragmentShaders []uint32

var mapPath string = "./maps/map.json"
var unmarshaledMapJson types.JsonMap 
// var jsonMapFileData types.JsonMap

func init() {
	fmt.Println("--------")
	fmt.Println("init()")

	// Open our jsonFile
	jsonFile, err := os.Open(mapPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

    byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(byteValue, &unmarshaledMapJson)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Printf("Successfully Opened %v \n", mapPath)
	fmt.Printf("jsonFile : %+v\n", *jsonFile)
	fmt.Printf("unmarshaledMapJson : %+v\n", unmarshaledMapJson)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
    runtime.LockOSThread()

}

func main() {
	fmt.Println("--------")
	fmt.Println("main()")

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()
	
	// windowFlags := uint32(sdl.WINDOW_SHOWN) | uint32(sdl.WINDOW_FULLSCREEN_DESKTOP)
	window, err := sdl.CreateWindow("Golang SDL Playground", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(WindowWidth * 2), int32(WindowHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	// window.SetGrab(true)

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_PRESENTVSYNC);
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	// // func (renderer *Renderer) CreateTexture(format uint32, access int, w, h int32) (*Texture, error)
	texture, err = renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_STREAMING, 1, 1)
	// texture, err = renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_STREAMING, WindowWidth, WindowHeight)
	if err != nil {
		panic(err)
	}

	cellWidth = WindowWidth / int32(unmarshaledMapJson.Width)
	cellHeight = WindowHeight / int32(unmarshaledMapJson.Height)

	color := types.NewColor(127, 127, 127, 127)
	Player = types.NewPlayer(renderer, texture, unmarshaledMapJson.Player.X, unmarshaledMapJson.Player.Y, utils.ToRadians(unmarshaledMapJson.Player.Angle), int(cellWidth), int(cellHeight), color)

	// sdl.StartTextInput()

	// go mainRoutine()
	mainRoutine()


	// renderer.Destroy()
	// window.Destroy()
	// sdl.Quit()
}

func pollEvents() {

	// event := sdl.PollEvent()
	keyStates := sdl.GetKeyboardState()
	// fmt.Println("keyStates : ", keyStates)
	// mouseStateX, mouseStateY, mouseState := sdl.GetMouseState()
	// fmt.Println("mouseState : ", mouseStateX, mouseStateY, mouseState)

	UpState = keyStates[26] == 1 || keyStates[82] == 1
	LeftState = keyStates[4] == 1 || keyStates[80] == 1
	DownState = keyStates[22] == 1 || keyStates[81] == 1
	RightState = keyStates[7] == 1 || keyStates[79] == 1

	Player.Move(UpState, LeftState, DownState, RightState)

	// fmt.Println(
	// 	"states : ",
	// 	UpState,
	// 	LeftState,
	// 	DownState,
	// 	RightState,
	// )
	
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		// fmt.Printf("event : %+v\n", event)
		switch event.(type) {
		case *sdl.QuitEvent:
			fmt.Println("QuitEvent")
			//renderer.Destroy()
			//window.Destroy()
			//sdl.Quit()
			break
		// case *sdl.KeyboardEvent:
		// 	fmt.Println("KeyboardEvent")
		// 	// eventText := event.GetText()
		// 	// fmt.Printf("eventText : %+v\n", eventText)
		// case *sdl.TextInputEvent:
		// 	fmt.Println("TextInputEvent")
		// 	// eventText := event.GetText()
		// 	// fmt.Printf("eventText : %+v\n", eventText)
		}
	}
}

func mainRoutine() {
	fmt.Println("--------")
	fmt.Printf("mainRoutine(%+v)\n", framesCount)

	pollEvents()

	// Clear screen
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear();

	drawWalls()

	// drawPickables()

	drawPlayer()

	// drawUI()

	// //Render texture to screen
	// renderer.Copy(texture, nil, nil);

	//Update screen
	renderer.Present();


	time.Sleep((1 / 24) * time.Second)
	// time.Sleep(1 * time.Second)
	framesCount += 1
	mainRoutine()
}

func drawWalls() {
	// rectangle := sdl.Rect{WindowWidth / 2, WindowHeight / 2, 10, 10}
	// renderer.SetDrawColor(255, 255, 255, 255)
	// renderer.FillRect(&rectangle)
	// renderer.DrawRect(&rectangle)

    for i := 0; i < len(unmarshaledMapJson.Walls); i++ {
		wall := unmarshaledMapJson.Walls[i]
		points := unmarshaledMapJson.Walls[i].Points

		renderer.SetDrawColor(wall.Color.R, wall.Color.G, wall.Color.B, wall.Color.A)
		renderer.DrawLines(points)

		for j := 0; j < len(points); j++ {
			point := points[j]
			color := types.NewColor(255, 0, 0, 255)
			rect := types.NewSquareShape(renderer, texture, int(point.X - 2), int(point.Y - 2), 4, 4, color)
			rect.Draw()
		}
	}

	// window.UpdateSurface()
}

func drawPickables() {
    for i := 0; i < len(unmarshaledMapJson.Pickables); i++ {
		pickable := unmarshaledMapJson.Pickables[i]

		rectangle := sdl.Rect{int32(pickable.X) - int32(10 * pickable.Scale / 2), int32(pickable.Y) - int32(10 * pickable.Scale / 2), int32(10 * pickable.Scale), int32(10 * pickable.Scale)}
		renderer.SetDrawColor(255, 0, 127, 255)
		renderer.FillRect(&rectangle)
		renderer.DrawRect(&rectangle)
	}
	// window.UpdateSurface()
}

func drawPlayer() {
	Player.Draw()

	// for k := int(Player.Angle + float64(Ouverture / 2)); k < int(Player.Angle + float64(Ouverture / 2)); k++ {
	for k := 0; k < int(WindowWidth); k++ {

		// kAngle := utils.ToRadians(float64(k) * float64(Ouverture) / float64(WindowWidth)) - float64(framesCount) * TimeFactor
		kAngle := utils.ToRadians(float64(k) * float64(Ouverture) / float64(WindowWidth))
		kRayEndX := float64(Player.X) + 500 * math.Cos(kAngle + Player.Angle / 2)
		kRayEndY := float64(Player.Y) - 500 * math.Sin(kAngle + Player.Angle / 2)

		if k % int(WindowWidth / 20) == 0 {
			renderer.SetDrawColor(127, 127, 127, 255)
			renderer.DrawLine(int32(Player.X), int32(Player.Y), int32(kRayEndX), int32(kRayEndY))
		}

		for i := 0; i < len(unmarshaledMapJson.Walls) ; i++ {
			wall := unmarshaledMapJson.Walls[i]
			
			for j := 0; j < len(wall.Points) ; j++ {
				start := wall.Points[j]
				var end sdl.Point
				if j != len(wall.Points) - 1 {
					end = wall.Points[j + 1]
				} else {
					end = wall.Points[0]
				}
				x, y, didHitSomething := utils.GetIntersection(Player.X, Player.Y, int(kRayEndX), int(kRayEndY), int(start.X), int(start.Y), int(end.X), int(end.Y))
				// fmt.Printf("%+v / %+v / %+v : \n", x, y, didHitSomething)
				
				
				if didHitSomething {
					
					rectangle := sdl.Rect{int32(x) - 5, int32(y) - 5, 10, 10}
					renderer.SetDrawColor(127, 0, 255, 255)
					renderer.FillRect(&rectangle)
					renderer.DrawRect(&rectangle)

					// var rectangleWidth int32 = 1
					rectangleHeight := float64(WindowHeight) / (utils.GetDistance(float64(Player.X), float64(Player.Y), float64(x), float64(y)) * WallsFactor)
					// fmt.Printf("%v\n", rectangleHeight)
					// rectangle := sdl.Rect{WindowWidth + int32(k), 0 + (WindowHeight - int32(rectangleHeight)) / 2, rectangleWidth, int32(rectangleHeight)}
					renderer.SetDrawColor(wall.Color.R, wall.Color.G, wall.Color.B, wall.Color.A)
					// renderer.FillRect(&rectangle)
					// renderer.DrawRect(&rectangle)
					renderer.DrawLine(WindowWidth + int32(k), WindowHeight / 2 - int32(rectangleHeight / 2), WindowWidth + int32(k), WindowHeight / 2 + int32(rectangleHeight / 2))
				}
			}
		}

	}

}

func drawUI() {

}