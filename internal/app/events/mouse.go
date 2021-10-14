package events

import (
	"fmt"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func MouseButtonCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	fmt.Println(button, action, mods)
}

func CursorPosCallback(window *glfw.Window, xpos float64, ypos float64) {
	fmt.Println(xpos, ypos)
}
