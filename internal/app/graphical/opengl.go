package graphical

import (
	_ "image/png"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()

	// init opengl
	if err := gl.Init(); err != nil {
		panic(err)
	}
}
