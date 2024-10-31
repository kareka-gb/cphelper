package main

import (
	"os"

	"github.com/kareka-gb/cphelper/apps"
)

func main() {
	app := apps.GetterApp{}
	app.Run(os.Args[1])
}
