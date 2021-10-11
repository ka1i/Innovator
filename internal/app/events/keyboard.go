package events

import (
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func Keyboard(window *glfw.Window) {
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		log.Println("bye bye!")
		window.SetShouldClose(true)
	}
}
