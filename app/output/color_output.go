package output

import (
	"github.com/fatih/color"
	"os"
)

func GetOutSuccess(errName string) {
	green := color.New(color.BgGreen)
	green.Println(errName)
}

func GetOutError(errName string) {
	red := color.New(color.FgRed)
	red.Println(errName)

	os.Exit(1)
}
