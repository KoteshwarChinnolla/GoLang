package main

import (
	"Factory_method_design_pattern/interfaces"
	"Factory_method_design_pattern/linux"
	"Factory_method_design_pattern/windows"
	"fmt"
)

// concrete class
func UIrendering(os string) interfaces.UI {
	switch os {
	case "windows":
		fmt.Println("windows")
		return &windows.WindowsFactory{}
	case "linux":
		return &linux.LinuxFactory{}
	}
	return nil
}
