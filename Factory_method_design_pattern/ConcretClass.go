package main

// concrete class
func UIrendering(os string) UI {
	switch os {
	case "windows":
		return &WindowsFactory{}
	case "linux":
		return &LinuxFactory{}
	}
	return nil
}
