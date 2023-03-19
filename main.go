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
	"github.com/go-gl/gl/v4.6-core/gl"
)

// type JsonMap struct {
//     width int
//     height int
// 	cells []int
// }

// Global vars
var windowWidth int = 320
var windowHeight int = 240
var window *glfw.Window
var program uint32

var mapPath string = "./maps/map.json"
var jsonMapFileData types.JsonMap

var square = []float32{
    -0.5, 0.5, 0,
    -0.5, -0.5, 0,
    0.5, -0.5, 0,

    -0.5, 0.5, 0,
    0.5, 0.5, 0,
    0.5, -0.5, 0,
}

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
	fmt.Printf("%+v\n", jsonFile)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
    runtime.LockOSThread()

    
    if err := glfw.Init(); err != nil {
		panic(err)
    }
    
    glfw.WindowHint(glfw.Resizable, glfw.False)
    glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
    glfw.WindowHint(glfw.ContextVersionMinor, 1)
    glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
    glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

    window, err := glfw.CreateWindow(windowWidth, windowHeight, "OpenGL Playground", nil, nil)
    if err != nil {
            panic(err)
    }
    window.MakeContextCurrent()
	
    defer glfw.Terminate()


    if err := gl.Init(); err != nil {
		panic(err)
    }
    version := gl.GoStr(gl.GetString(gl.VERSION))
    fmt.Println("OpenGL version", version)

    program = gl.CreateProgram()
    gl.LinkProgram(program)

    for !window.ShouldClose() {
		mainRoutine()
		// Do OpenGL stuff.
		window.SwapBuffers()
		glfw.PollEvents()
    }
}

func main() {
	fmt.Println("--------")
	fmt.Println("main()")
	fmt.Println("--------")

	// go mainRoutine()

}

func mainRoutine() {
	fmt.Println("--------")
	fmt.Println("mainRoutine()")
	fmt.Println("--------")

	drawMap()

	time.Sleep(1 * time.Second)
	// mainRoutine()
}

func drawMap() {
	fmt.Println("drawMap")
	// gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	// gl.UseProgram(program)
	
	// glfw.PollEvents()
	// window.SwapBuffers()
}