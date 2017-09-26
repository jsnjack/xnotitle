package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

var debugFlag *bool

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
	for _, item := range windows {
		if strings.Contains(item.Title, *name) {
			if *debugFlag {
				fmt.Println("Hiding window titlebar:", item.ID, item.Title)
			}
			err := HideTitleBar(item.ID)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
