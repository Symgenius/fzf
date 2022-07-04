package main

import (
	fzf "github.com/Symgenius/fzf/src"
	"github.com/Symgenius/fzf/src/protector"
)

var version string = "0.30"
var revision string = "devel"

func main() {
	protector.Protect()
	fzf.Run(fzf.ParseOptions(), version, revision)
}
