package graphical

import (
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()

	// init glfw
	if err := glfw.Init(); err != nil {
		panic(err)
	}
}

type hints struct {
	title         string
	width, height int
	resizable     bool
	borderless    bool
	maximized     bool
}

func WindowHint() *hints {
	return &hints{
		title:  "glfw3",
		width:  1024,
		height: 600,
	}
}

// Title option sets the title (caption) of the window.
func (h *hints) Title(title string) {
	h.title = title
}

// Size option sets the width and height of the window.
func (h *hints) Size(width, height int) {
	h.width = width
	h.height = height
}

// Resizable option makes the window resizable by the user.
func (h *hints) Resizable() {
	h.resizable = true
}

// Borderless option makes the window borderless.
func (h *hints) Borderless() {
	h.borderless = true
}

// Maximized option makes the window start maximized.
func (h *hints) Maximized() {
	h.maximized = true
}

type Window func() (w *glfw.Window, err error)

func CreateWindow(hint *hints) Window {
	return func() (*glfw.Window, error) {
		// init glfw
		if hint.resizable {
			glfw.WindowHint(glfw.Resizable, glfw.True)
		} else {
			glfw.WindowHint(glfw.Resizable, glfw.False)
		}
		if hint.borderless {
			glfw.WindowHint(glfw.Decorated, glfw.False)
		}
		if hint.maximized {
			glfw.WindowHint(glfw.Maximized, glfw.True)
		}

		// create window
		window, err := glfw.CreateWindow(hint.width, hint.height, hint.title, nil, nil)
		if err != nil {
			return nil, err
		}
		if hint.maximized {
			hint.width, hint.height = window.GetFramebufferSize() // set o.width and o.height to the window size due to the window being maximized
		}
		return window, nil
	}
}
