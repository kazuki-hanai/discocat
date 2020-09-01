package main

import (
	"github.com/fatih/color"
	"fmt"
	"os"
)

func printErr(err error) {
	red  := color.New(color.FgRed).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Println(cyan(commandName), red(err.Error()))
}

func exitErr(err error) {
	printErr(err)
	os.Exit(1)
}
