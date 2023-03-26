package utils

import (
	// "github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/gl/v4.6-core/gl"
	// "strings"
	"strconv"
	"fmt"
)

func HexColorToRGB(hexColor string) (float32, float32, float32) {
	r, errR := strconv.ParseInt(hexColor[0:2], 16, 64)
	g, errG := strconv.ParseInt(hexColor[2:4], 16, 64)
	b, errB := strconv.ParseInt(hexColor[4:6], 16, 64)

	if (errR != nil || errG != nil || errB != nil) {
		fmt.Println("errR : ", errR)
		fmt.Println("errG : ", errG)
		fmt.Println("errB : ", errB)
		return 0.5, 0.5, 0.5
	}
	return float32(r) / 256, float32(g) / 256, float32(b) / 256
}

// makeVao initializes and returns a vertex array from the points provided.
func MakeVao(points []float32) uint32 {
    var vbo uint32
    gl.GenBuffers(1, &vbo)
    gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
    gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)
    
    var vao uint32
    gl.GenVertexArrays(1, &vao)
    gl.BindVertexArray(vao)
    gl.EnableVertexAttribArray(0)
    gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
    gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
    
    return vao
}