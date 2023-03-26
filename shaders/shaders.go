package shaders

import (
	// "github.com/xLeDocteurx/go-opengl-playground/types"
	"github.com/xLeDocteurx/go-opengl-playground/utils"
	
	"fmt"
	"strings"

	// "github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/gl/v4.6-core/gl"
)

// func InitAndLoadShaders(program *uint32) ([]uint32, []uint32) {
// 	var vertexShaders []uint32
// 	var fragmentShaders []uint32

// 	redFragmentShader := GetFragmentShaderSource("FF0000", 1.0)
// 	fragmentShaders = append(fragmentShaders, redFragmentShader)
//     gl.AttachShader((*program), redFragmentShader)

// 	greenFragmentShader := GetFragmentShaderSource("00FF00", 1.0)
// 	fragmentShaders = append(fragmentShaders, greenFragmentShader)
//     gl.AttachShader((*program), greenFragmentShader)

// 	blueFragmentShader := GetFragmentShaderSource("0000FF", 1.0)
// 	fragmentShaders = append(fragmentShaders, blueFragmentShader)
//     gl.AttachShader((*program), blueFragmentShader) 

// 	basicVertexShader := GetCompiledVertexShader(1.0)
// 	vertexShaders = append(vertexShaders, basicVertexShader)
//     gl.AttachShader((*program), basicVertexShader) 
	
// 	return fragmentShaders, vertexShaders
// } 

func GetVertexShaderSource(float float32) string {
	var vertexShaderSource string = fmt.Sprintf("#version 410\n in vec3 vp; void main() {	gl_Position = vec4(vp, %f); }\x00", float)

    // compiledVertexShader, err := CompileShader(vertexShaderSource, gl.VERTEX_SHADER)
    // if err != nil {
    //     panic(err)
    // }

	return vertexShaderSource
}

func GetFragmentShaderSource(hexColor string, alpha float32) string {
	r, g, b := utils.HexColorToRGB(hexColor)
	var fragmentShaderSource string = fmt.Sprintf("#version 410\n out vec4 frag_colour; void main() {	frag_colour = vec4(%f, %f, %f, %f); }\x00", r, g, b, alpha)
	
    // compiledFragmentShader, err := CompileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
    // if err != nil {
    //     panic(err)
    // }
	
	return fragmentShaderSource
}

func CompileShader(source string, shaderType uint32) (uint32, error) {
    shader := gl.CreateShader(shaderType)
    
    csources, free := gl.Strs(source)
    gl.ShaderSource(shader, 1, csources, nil)
    free()
    gl.CompileShader(shader)
    
    var status int32
    gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
    if status == gl.FALSE {
        var logLength int32
        gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
        
        log := strings.Repeat("\x00", int(logLength+1))
        gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))
        
        return 0, fmt.Errorf("failed to compile %v: %v", source, log)
    }
    
    return shader, nil
}