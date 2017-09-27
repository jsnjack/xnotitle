package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

var debugFlag *bool
var processedWindows []string

func main() {
	debugFlag = flag.Bool("debug", false, "Display debug information")
	nameFlag := flag.String("name", "Mozilla Firefox", "Hide title bar for windows which title contains <name>")
	periodFlag := flag.Int("period", 1000, "Check for new windows every <period> ms")
	flag.Parse()
	if *debugFlag {
		fmt.Printf("Name flag: %s\n", *nameFlag)
	}
	c := time.Tick(time.Duration(*periodFlag) * time.Millisecond)
	for range c {
		CheckWindows(GetWindows(), nameFlag)
	}
}

// CheckWindows checks if windows needs to be hidden
func CheckWindows(windows []*Window, name *string) {
	if *debugFlag {
		fmt.Printf("Inspecting %d windows\n", len(windows))
	}
	var newList []string
	for _, item := range windows {
		newList = append(newList, item.ID)
		if !in(&item.ID, &processedWindows) && strings.Contains(item.Title, *name) {
			if *debugFlag {
				fmt.Println("Hiding window titlebar:", item.ID, item.Title)
			}
			err := HideTitleBar(item.ID)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
	processedWindows = newList
}

// Checks if windowID is in windowIDList
func in(item *string, list *[]string) bool {
	for _, value := range *list {
		if *item == value {
			return true
		}
	}
	return false
}
