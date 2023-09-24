package main

import (
	"os"
	"flag"
	"github.com/quamejnr/heimdall/util"

)

func main() {
	// Get into the user documents directory
	root, _ := os.UserHomeDir()
	os.Chdir(root + "/Documents")

	// Get command line args
	cmd := flag.String("cmd", "", "Command to be run")
	file := flag.String("f", "", "File to be searched for")
	flag.Parse()

	files := util.FindFiles(*file)
	util.RunCommand(*cmd, files)
}


