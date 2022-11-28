/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"Go-cup/cmd"
	"github.com/common-nighthawk/go-figure"
)

func main() {
	myFigure := figure.NewColorFigure("Go cup!!!", "", "red", false)
	myFigure.Print()
	cmd.Execute()
}
