package main

import (
	"flag"
	"fmt"
	"github.com/quamejnr/heimdall/util"
	"os"
)

func main() {
	// Get command line args
	cmd := flag.String("cmd", "", "Shell command to be run")
	fileName := flag.String("f", "", "Name of file to be searched for")
	strict := flag.Bool("s", false, "Strict name matching. If `true` filename/folder provided should be the exact match.")
	flag.Parse()

	// Get into the user documents directory
	root, _ := os.UserHomeDir()
	os.Chdir(root + "/Documents")

	files := util.FindFiles(*fileName, *strict)
	var f string

	switch numFiles := len(files); {
	case numFiles == 0:
		fmt.Println("File Not Found")
		return
	case numFiles == 1:
		f = files[0]
	default:
		f = util.PickFile(files)
	}
	util.RunCommand(*cmd, f)

}
