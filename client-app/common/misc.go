package common

import (
	"fmt"
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/fatih/color"
)

func makeSthBoldColorfulSth(attr color.Attribute, format string, x ...interface{}) string {
	ba := color.New(color.Bold, attr)
	return ba.Sprintf(format, x)
}

func EvalError(err error) {
	if err == terminal.InterruptErr {
		os.Exit(0)
	}
	fmt.Println(makeSthBoldColorfulSth(color.FgRed, "error: %v\n", err))
}

func Info(msg interface{}) string {
	return makeSthBoldColorfulSth(color.FgYellow, "%s\n", msg)
}

func PressAnyToContinue() {
	fmt.Println(Info("Press any key to continue..."))
	_, _ = fmt.Scanln()
}

func PrintResult(msg interface{}) {
	log.Println(makeSthBoldColorfulSth(color.FgBlue, "Result: %v\n", msg))
	PressAnyToContinue()
}
