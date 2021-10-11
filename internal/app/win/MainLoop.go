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

func MainLoop() {
	w := initWindow()
	// render loop
	for !w.ShouldClose() {
		renderLoop(w)
	}

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

	return w
}

func renderLoop(w *glfw.Window) {
	// event process
	events.Keyboard(w)

	// glfw background
	gl.ClearColor(0.98, 0.98, 0.98, 0.7)                //状态设置
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT) //状态使用

	//检查调用事件，交换缓冲
	w.SwapBuffers()
	glfw.PollEvents()
}
