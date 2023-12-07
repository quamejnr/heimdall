/*
Copyright © 2023 KWAME AKUAMOAH-BOATENG <kaboateng14@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/quamejnr/heimdall/cmd/utils"
	"github.com/spf13/cobra"
)

var (
	fileName, command string
	strict            bool
)

const defaultDir = "Documents"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "heimdall",
	Short: "Heimdall is a CLI tool that allows you to run your basic shell commands on a file/folder from anywhere.",
	Long: `
Heimdall is a CLI tool that allows you to run your basic shell commands like 'cat', 'code', 'ls', 'vim' etc. on a file 
The beauty about Heimdall is it allows you to run these commands on your file/folder without having to navigate to the directory your file/folder exists in. 
For example: 'heimdall vim example.go' or 'heimdall -c=vim -f=example.go' will open your file 'example.go' in vim. 

NB: File/Folder needs to be in '$USER/Documents' directory.
`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			command, fileName = args[0], args[1]
		}
		heimdall(command, fileName, strict)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&command, "cmd", "c", "", "Shell command")
	rootCmd.PersistentFlags().StringVarP(&fileName, "file", "f", "", "File/Folder to run shell command on")
	rootCmd.Flags().BoolVarP(&strict, "strict", "s", true, "Strict name matching")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func heimdall(command string, fileName string, strict bool) {
	// Get into the user documents directory
	root, _ := os.UserHomeDir()
	lookUpDir := utils.GetLookUpDir("HEIMDALL_DIR", defaultDir)
	dir := filepath.Join(root, lookUpDir)
	fmt.Println(dir)
	os.Chdir(dir)

	files := utils.FindFiles(fileName, strict)
	if len(files) == 0 {
		fmt.Printf("'%s' not Found\n", fileName)
		return
	}
	f := utils.PickFile(files)
	utils.RunCommand(command, f)
}
