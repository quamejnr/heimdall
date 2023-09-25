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
	strict := flag.Bool("s", true, "Strict name matching. If `true` filename/folder provided should be the exact match.")
	flag.Parse()

	// Get into the user documents directory
	root, _ := os.UserHomeDir()
	os.Chdir(root + "/Documents")

	files := util.FindFiles(*fileName, *strict)
	if len(files) == 0 {
		fmt.Printf("'%s' not Found\n", *fileName)
		return 
	}
	f := util.PickFile(files)
	util.RunCommand(*cmd, f)

}
