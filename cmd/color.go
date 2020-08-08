package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	out          = colorable.NewColorableStdout()
	successColor = color.New(color.FgHiGreen)
	infoColor    = color.New(color.FgHiBlue)
	warnColor    = color.New(color.FgHiYellow)
	errorColor   = color.New(color.FgHiRed)
)

// coloredOutput outputs the input from fp and args to w, coloring it with the f function.
func coloredOutput(w io.Writer, fp *os.File, args []string, f func(w io.Writer, a ...interface{})) error {
	if !terminal.IsTerminal(int(fp.Fd())) {
		// If there is input from the pipe, it is also colored and output.
		b, err := ioutil.ReadAll(fp)
		if err != nil {
			return fmt.Errorf("failed to read pipe: %v", err)
		}
		f(w, fmt.Sprintln(string(b)))
	}

	for _, v := range args {
		f(w, fmt.Sprintln(v))
	}
	return nil
}
