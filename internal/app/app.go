package app

import (
	"fmt"
	"log"
	"os"

	"github.com/ka1i/innovator/internal/app/win"
	"github.com/ka1i/innovator/internal/pkg/usage"
	"github.com/ka1i/innovator/pkg/version"
)

type app struct {
	success int
	failure int
}

func (app *app) barry(argc int, argv []string) int {
	var err error

	switch argv[0] {
	case "-f", "file":
		win.MainLoop(argv[1])
	case "-h", "--help", "help":
		usage.Usage()
	case "-v", "--version", "version":
		version.Version.Print()
	default:
		err = fmt.Errorf("innovator usage: innovator -h")
	}
	if err != nil {
		log.Printf("%v\n", err)
		return app.failure
	}
	return app.success
}

var App = GetApp()

func GetApp() *app {
	return &app{
		success: 0,
		failure: 137,
	}
}

func (app *app) Innovator() int {
	var exitcode int = app.success
	if len(os.Args) > 1 {
		var argc = len(os.Args)
		var argv = os.Args[1:]
		exitcode = app.barry(argc, argv)
	} else {
		fmt.Println("innovator usage: innovator -h")
		exitcode = app.failure
	}
	return exitcode
}
