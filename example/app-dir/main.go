package main

import (
	"github.com/fastscripts/toolkit"
)

func main() {
	var tools toolkit.Tools

	tools.CreateDirIfNotExists("./test-dir")
}
