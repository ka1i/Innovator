package win

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/ka1i/innovator/internal/app/events"
	"github.com/ka1i/innovator/internal/app/graphical"
	"github.com/ka1i/innovator/internal/pkg/usage/utils"
)

const (
	Title string = "Anime Engine"
)

var (
	BackgroundColor mgl32.Vec4 = mgl32.Vec4{0.55, 0.55, 0.55, 0.0} // background color
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func initWindow() *glfw.Window {
	// glfw hint setup
	hint := graphical.WindowHint()
	hint.Title(Title)
	hint.Size(1024, int(1024/utils.AspectRatio))
	hint.Resizable()

	glfw.WindowHint(glfw.ContextVersionMajor, 4)                //OpenGL大版本
	glfw.WindowHint(glfw.ContextVersionMinor, 1)                //OpenGl小版本
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile) //明确核心模式
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)    //Mac使用

	// init glfw window
	window := graphical.CreateWindow(hint)
	w, err := window()
	if err != nil {
		panic(err)
	}

	w.MakeContextCurrent()
	w.SetSizeLimits(208, int(1024/utils.AspectRatio), gl.DONT_CARE, gl.DONT_CARE)
	w.SetAspectRatio(16, 9)

	// display env version
	log.Printf("GLFW: %s \n", glfw.GetVersionString())
	log.Printf("openGL: %s \n", gl.GoStr(gl.GetString(gl.VERSION)))

	// openGL viewport init
	width, height := w.GetFramebufferSize()
	gl.Viewport(0, 0, int32(width), int32(height))

	// events register
	w.SetKeyCallback(events.KeyCallback)
	w.SetFramebufferSizeCallback(events.FramebufferSizeCallback)

	// disable vsync : default :0
	glfw.SwapInterval(0)

	return w
}

func makeVAO() uint32 {
	//连接顶点属性
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// position attribute
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(0, 4, gl.FLOAT, false, 4*4, 0)

	return vao
}

func MainLoop() {
	w := initWindow()

	// Render Loop
	program, err := graphical.NewProgram(vertexShaderSource, fragmentShaderSource)
	if err != nil {
		panic(err)
	}

	vao := makeVAO()

	// Load the texture
	texture, err := graphical.NewTexture("example.png")
	if err != nil {
		log.Fatalln(err)
	}

	width, height := w.GetFramebufferSize()

	gl.UseProgram(program)

	projection := mgl32.Ortho(0, float32(width), float32(height), 0, -1, 1) //mgl32.Mat4{}

	//线框模式(Wireframe Mode)
	//gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	gl.Enable(gl.COLOR_WRITEMASK)
	gl.Enable(gl.DEPTH)
	gl.DepthFunc(gl.LESS)

	gl.BindFragDataLocation(program, 0, gl.Str("fragmentColor\x00"))

	for !w.ShouldClose() {
		// glfw background
		gl.ClearColor(BackgroundColor.Elem())               //状态设置
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT) //状态使用

		// render window
		gl.UseProgram(program)

		// texture unit
		gl.Uniform1i(gl.GetUniformLocation(program, gl.Str("imgtexture\x00")), 0)

		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, texture)

		// projection
		projection = mgl32.Ortho(0, float32(width), float32(height), 0, -1, 1)

		projectionLoc := gl.GetUniformLocation(program, gl.Str("projection\x00"))
		gl.UniformMatrix4fv(projectionLoc, 1, false, &projection[0])

		model := mgl32.Ident4()
		model = model.Mul4(mgl32.Scale3D(float32(width)/2, float32(height)/2, 1))

		modelLoc := gl.GetUniformLocation(program, gl.Str("model\x00"))
		gl.UniformMatrix4fv(modelLoc, 1, false, &model[0])

		// bind vao
		gl.BindVertexArray(vao)

		gl.DrawArrays(gl.TRIANGLES, 0, 6)

		//检查调用事件，交换缓冲
		w.SwapBuffers()
		glfw.PollEvents()
	}
	glfw.Terminate()
}
