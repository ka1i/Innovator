package win

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
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
	w.SetFramebufferSizeCallback(framebuffer_size_callback)

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
	gl.GenBuffers(1, &vbo)                                                            //创建顶点缓冲对象，绑定id
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)                                               //把新创建的缓冲绑定到GL_ARRAY_BUFFER目标
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW) //把用户定义的数据复制到当前绑定缓冲

	var ebo uint32
	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	// position attribute
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)
	// color attribute
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(3*4))
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
	texture, err := graphical.NewTexture("container.jpeg")
	if err != nil {
		log.Fatalln(err)
	}

	//线框模式(Wireframe Mode)
	//gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)

	var fps uint = 0
	fpsTracker := glfw.GetTime()
	for !w.ShouldClose() {
		// fps
		currentTime := glfw.GetTime()
		if currentTime-fpsTracker >= 1.0 {
			log.Printf("fps:%d/s\n", fps)
			fpsTracker = currentTime
			fps = 0
		}
		fps++

		// glfw background
		gl.ClearColor(0.2, 0.3, 0.4, 1)                     //状态设置
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT) //状态使用

		// event process
		events.Keyboard(w)

		// render window
		gl.UseProgram(program)

		gl.BindTexture(gl.TEXTURE_2D, texture)
		gl.BindVertexArray(vao)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.PtrOffset(0))
		//gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices)/3))
		// gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.PtrOffset(0))
		// gl.BindVertexArray(0)

		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, texture)

		//检查调用事件，交换缓冲
		w.SwapBuffers()
		glfw.PollEvents()
	}
	glfw.Terminate()
}

func framebuffer_size_callback(window *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}
