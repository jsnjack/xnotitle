package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// Window represents window structure
type Window struct {
	ID    string
	Title string
}

// GetWindows returns list of the windows
func GetWindows() []*Window {
	var windowList []*Window
	cmd := exec.Command("wmctrl", "-l")
	output, err := cmd.Output()
	if err != nil {
		return windowList
	}
	for _, line := range strings.Split(string(output), "\n") {
		splitted := strings.SplitN(line, " ", 2)
		switch len(splitted) {
		case 2:
			splittedTitle := strings.SplitN(splitted[1], " ", 4)
			windowList = append(windowList, &Window{ID: splitted[0], Title: splittedTitle[3]})
		default:
			fmt.Println("Ignored")
		}
	}
	return windowList
}

// HideTitleBar hides titlebar of the window
func HideTitleBar(id string) error {
	cmd := exec.Command("xprop", "-id", id, "-f", "_GTK_HIDE_TITLEBAR_WHEN_MAXIMIZED", "32c", "-set", "_GTK_HIDE_TITLEBAR_WHEN_MAXIMIZED", "0x1")
	return cmd.Run()
}
