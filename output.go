package main

import (
	"github.com/fatih/color"
	"fmt"
	"github.com/urfave/cli/v2"
)

var (
	red  = color.New(color.FgRed).SprintFunc()
	cyan = color.New(color.FgCyan).SprintFunc()
)

func printErr(err error) {
	fmt.Println(cyan(commandName), red(err.Error()))
}

func exitErr(err error) cli.ExitCoder {
	return cli.Exit(red(err.Error()), 1)
}
