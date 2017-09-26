package main

import (
	"fmt"
	"strings"
)

var oldWindows []string

func main() {
	CheckWindows(GetWindows(), "Firefox")
}

// CheckWindows checks if windows needs to be hidden
func CheckWindows(windows []*Window, name string) {
	for _, item := range windows {
		if strings.Contains(item.Title, name) {
			fmt.Println("Hiding", item.Title)
			err := HideTitleBar(item.ID)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
