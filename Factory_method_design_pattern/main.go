package main

import (
	"fmt"
)

// Client
func main() {
	// Create a sedan car using the sedan factory
	WindowsUI := UIrendering("windows")
	LinuxUI := UIrendering("linux")
	fmt.Println(WindowsUI.RenderImage().Render())
	fmt.Println(WindowsUI.RenderButten().Render())
	fmt.Println(LinuxUI.RenderImage().Render())
	fmt.Println(LinuxUI.RenderButten().Render())
}
