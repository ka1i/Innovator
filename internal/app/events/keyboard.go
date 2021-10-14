package events

import (
	"fmt"
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func KeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	fmt.Println(key, scancode, action, mods)
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		log.Println("bye bye!")
		window.SetShouldClose(true)
	}
}
