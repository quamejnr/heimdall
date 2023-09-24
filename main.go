package main

import (
	"flag"
	"github.com/quamejnr/heimdall/util"
	"os"
)

func main() {
	// Get into the user documents directory
	root, _ := os.UserHomeDir()
	os.Chdir(root + "/Documents")

	// Get command line args
	cmd := flag.String("cmd", "", "Shell command to be run")
	file := flag.String("f", "", "Name of file to be searched for")
	strict := flag.Bool("s", false, "Strict name matching. If `true` filename/folder provided should be the exact match.")
	flag.Parse()

	files := util.FindFiles(*file, *strict)
	util.RunCommand(*cmd, files)
}
