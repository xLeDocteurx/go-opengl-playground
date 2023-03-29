package utils

import (
	// "github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/gl/v4.6-core/gl"
	// "strings"
	"strconv"
	"fmt"
    "math"
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

func GetIntersection(x1, y1, x2, y2, x3, y3, x4, y4 int) (int, int, bool) {
    // Calculate slopes and y-intercepts of the two lines
    m1 := float64(y2-y1) / float64(x2-x1)
    b1 := float64(y1) - m1*float64(x1)
    m2 := float64(y4-y3) / float64(x4-x3)
    b2 := float64(y3) - m2*float64(x3)

    // Check if the lines are parallel
    if m1 == m2 {
        return 0, 0, false
    }

    // Calculate intersection point of the two lines
    x := int((b2 - b1) / (m1 - m2))
    y := int(m1*float64(x) + b1)

    return x, y, true
}

func GetDistance(x1, y1, x2, y2 float64) float64 {
    yLength := y2 - y1
    xLength := x2 - x1 
    return math.Sqrt(xLength * xLength + yLength * yLength)
}

func IntAbs(number int) int {
    if number < 0 {
        return -number
    }
    return number
}

func FloatAbs(number float64) float64 {
    if number < 0 {
        return -number
    }
    return number
}

func ToDegrees(angle float64) float64 {
    return angle * (180 / math.Pi);
}

func ToRadians(angle float64) float64 {
    return angle * (math.Pi / 180);
}