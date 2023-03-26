package main

import (
	"github.com/xLeDocteurx/go-opengl-playground/types"
	// "github.com/xLeDocteurx/go-opengl-playground/utils"
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
)

// type JsonMap struct {
//     width int
//     height int
// 	cells []int
// }

// Global vars
var WindowWidth int32 = 320
// var WindowWidth int32 = 640
var WindowHeight int32 = 240
// var WindowHeight int32 = 480
var cellWidth int32
var cellHeight int32

var frameCount int32 = 0

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
	window, err := sdl.CreateWindow("Golang SDL Playground", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(WindowWidth), int32(WindowHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	// window.SetGrab(true)

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_PRESENTVSYNC);
	if err != nil {
		panic(err)
	}
	
	// // func (renderer *Renderer) CreateTexture(format uint32, access int, w, h int32) (*Texture, error)
	texture, err = renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_STREAMING, 1, 1)
	// texture, err = renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_STREAMING, WindowWidth, WindowHeight)
	if err != nil {
		panic(err)
	}

	cellWidth = WindowWidth / int32(unmarshaledMapJson.Width)
	cellHeight = WindowHeight / int32(unmarshaledMapJson.Height)

	go mainRoutine()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
	}

}

func pollEvent() {
	
}

func mainRoutine() {
	fmt.Println("--------")
	fmt.Printf("mainRoutine(%+v)\n", frameCount)

	// //Clear screen
	// renderer.Clear();

	drawWalls()

	drawPickables()

	drawPlayer()

	// drawUI()

	// //Render texture to screen
	// renderer.Copy(texture, nil, nil);

	//Update screen
	renderer.Present();


	time.Sleep((1 / 24) * time.Second)
	frameCount += 1
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
			rect := types.NewSquareShape(renderer, texture, point.X - 2, point.Y - 2, 4, 4, color)
			rect.Draw()
		}
	}

	// window.UpdateSurface()
}

func drawPickables() {

    for i := 0; i < len(unmarshaledMapJson.Pickables); i++ {
		pickable := unmarshaledMapJson.Pickables[i]

		rectangle := sdl.Rect{pickable.X, pickable.Y, int32(10 * pickable.Scale), int32(10 * pickable.Scale)}
		renderer.SetDrawColor(255, 0, 127, 255)
		renderer.FillRect(&rectangle)
		renderer.DrawRect(&rectangle)
	}

	// window.UpdateSurface()
}

func drawPlayer() {
	color := types.NewColor(127, 127, 127, 127)
	player := types.NewPlayer(renderer, texture, int32(WindowWidth / 2), int32(WindowHeight - WindowHeight / 5), int32(cellWidth), int32(cellHeight), color)
	player.Draw()
}

func drawUI() {

}