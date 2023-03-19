package main

import (
	"github.com/xLeDocteurx/go-opengl-playground/types"
	// "github.com/xLeDocteurx/go-opengl-playground/utils"

	"fmt"
	// "sync"
	"time"
	"os"
	"runtime"
	"github.com/go-gl/glfw/v3.3/glfw"
)

// type JsonMap struct {
//     width int
//     height int
// 	cells []int
// }

// Global vars
var windowWidth int = 320
var windowHeight int = 240

var mapPath string = "./maps/map.json"
var jsonMapFileData types.JsonMap

func init() {
	fmt.Println("--------")
	fmt.Println("init()")
	fmt.Println("--------")
	// Open our jsonFile
	jsonFile, err := os.Open(mapPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Successfully Opened %v \n", mapPath)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	fmt.Println("--------")
	fmt.Println("main()")
	fmt.Println("--------")
	// utils.Blbl()

	// ///// Launch RoutineManager that manages the routines of this executable : 
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func(wg *sync.WaitGroup) {
	// 	defer wg.Done()
	// 	mainRoutine()
	// }(&wg)
	

	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(windowWidth, windowHeight, "OpenGL Playground", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	// window.SetAspectRatio(16, 9)

	go mainRoutine()

	for !window.ShouldClose() {
		// Do OpenGL stuff.
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func mainRoutine() {
	fmt.Println("--------")
	fmt.Println("mainRoutine()")
	fmt.Println("--------")

	time.Sleep(1 * time.Second)
	mainRoutine()
}