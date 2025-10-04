package main

import (
	"fmt"
)

// Client
func main() {
	WindowsUI := UIrendering("windows")
	MacUI := UIrendering("Mac")
	fmt.Println(WindowsUI.RenderImage().Render())
	fmt.Println(WindowsUI.RenderBatten().Render())
	fmt.Println(MacUI.RenderImage().Render())
	fmt.Println(MacUI.RenderBatten().Render())
}
