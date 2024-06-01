package main

import (
	"fmt"
	"runtime/debug"
	"strings"
)

var (
	version   = "dev"
	commitSHA = "none"
	buildDate = "unknown"
)

func main() {
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("example-go-application")
	fmt.Println(strings.Repeat("-", 30))
	fmt.Printf("* version:\t%s\n", version)
	fmt.Printf("* commit:\t%s\n", commitSHA)
	fmt.Printf("* build date:\t%s\n", buildDate)
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}
	fmt.Println(strings.Repeat("-", 30))
	for _, setting := range buildInfo.Settings {
		fmt.Printf("%s\t= %s\n", setting.Key, setting.Value)
	}
	fmt.Println(strings.Repeat("-", 30))
}
