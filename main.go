package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
)

var out = colorable.NewColorableStdout()

func main() {
	fmt.Println("hoge normal color")

	iColor := color.New(color.FgGreen)
	fI := iColor.FprintFunc()
	fI(out, "hoge green color\n")

	wColor := color.New(color.FgYellow)
	fW := wColor.FprintFunc()
	fW(out, "hoge yellow color\n")

	eColor := color.New(color.FgRed)
	fE := eColor.FprintFunc()
	fE(out, "hoge red color\n")
}
