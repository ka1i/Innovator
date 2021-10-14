package win

import (
	"log"
	"math"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/ka1i/innovator/internal/app/events"
	"github.com/ka1i/innovator/internal/app/graphical"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func initWindow() *glfw.Window {
	// glfw hint setup
	hint := graphical.WindowHint()
	hint.Title("Innovator: Hello World")
	hint.Size(800, 600)
	hint.Resizable()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)                //OpenGL大版本
	glfw.WindowHint(glfw.ContextVersionMinor, 3)                //OpenGl小版本
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile) //明确核心模式
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)    //Mac使用

	// init glfw window
	window := graphical.CreateWindow(hint)
	w, err := window()
	if err != nil {
		panic(err)
	}

	w.MakeContextCurrent()

	// display env version
	log.Printf("GLFW: %s \n", glfw.GetVersionString())
	log.Printf("openGL: %s \n", gl.GoStr(gl.GetString(gl.VERSION)))

	// openGL viewport init
	width, height := w.GetFramebufferSize()
	gl.Viewport(0, 0, int32(width), int32(height))

	// events register
	w.SetFramebufferSizeCallback(events.FramebufferSizeCallback)
	w.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
	w.SetKeyCallback(events.KeyCallback)
	w.SetMouseButtonCallback(events.MouseButtonCallback)
	w.SetCursorPosCallback(events.CursorPosCallback)

	// disable vsync
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

	var ebo uint32
	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	// position attribute
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)
	// color attribute
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)
	// texture coord attribute
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(6*4))
	gl.EnableVertexAttribArray(2)

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
	texture1, err := graphical.NewTexture("container.jpeg")
	if err != nil {
		log.Fatalln(err)
	}
	texture2, err := graphical.NewTexture("awesomeface.png")
	if err != nil {
		log.Fatalln(err)
	}

	width, height := w.GetFramebufferSize()

	gl.UseProgram(program)

	projection := mgl32.Mat4{}
	camera := mgl32.Mat4{}

	//线框模式(Wireframe Mode)
	//gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	gl.BindFragDataLocation(program, 0, gl.Str("fragmentColor\x00"))

	var fps uint = 0
	fpsTracker := glfw.GetTime()
	for !w.ShouldClose() {
		// fps
		currentTime := glfw.GetTime()
		if currentTime-fpsTracker >= 1.0 {
			log.Printf("*** Exit Press Esc *** fps:%d/s\n", fps)
			fpsTracker = currentTime
			fps = 0
		}
		fps++

		// glfw background
		gl.ClearColor(0.2, 0.3, 0.4, 1)                     //状态设置
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT) //状态使用

		// render window
		gl.UseProgram(program)

		// texture unit
		gl.Uniform1i(gl.GetUniformLocation(program, gl.Str("texture1\x00")), 0)
		gl.Uniform1i(gl.GetUniformLocation(program, gl.Str("texture2\x00")), 1)

		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, texture1)

		gl.ActiveTexture(gl.TEXTURE1)
		gl.BindTexture(gl.TEXTURE_2D, texture2)

		// projection * camera * model
		projection = mgl32.Ident4()
		camera = mgl32.Ident4()

		projection = mgl32.Perspective(mgl32.DegToRad(60), float32(width/height), 0.1, 100)
		projectionLoc := gl.GetUniformLocation(program, gl.Str("projection\x00"))
		gl.UniformMatrix4fv(projectionLoc, 1, false, &projection[0])

		radius := 100.0
		cameraX := math.Sin(glfw.GetTime()) * radius
		cameraZ := math.Cos(glfw.GetTime()) * radius
		camera = mgl32.LookAtV(mgl32.Vec3{float32(cameraX), 0, float32(cameraZ)}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
		cameraLoc := gl.GetUniformLocation(program, gl.Str("camera\x00"))
		gl.UniformMatrix4fv(cameraLoc, 1, false, &camera[0])

		// bind vao
		gl.BindVertexArray(vao)
		//gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.PtrOffset(0))
		for k, v := range cubePositions {
			model := mgl32.Ident4()
			model = model.Mul4(mgl32.Translate3D(v.Elem()))
			model = model.Mul4(mgl32.HomogRotate3D(float32(glfw.GetTime()*float64(k)), mgl32.Vec3{float32(k % 3), 0, 0}))

			modelLoc := gl.GetUniformLocation(program, gl.Str("model\x00"))
			gl.UniformMatrix4fv(modelLoc, 1, false, &model[0])

			gl.DrawArrays(gl.TRIANGLES, 0, 36)
		}

		//检查调用事件，交换缓冲
		w.SwapBuffers()
		glfw.PollEvents()
	}
	glfw.Terminate()
}
