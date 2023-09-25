/*
Copyright © 2023 KWAME AKUAMOAH-BOATENG <kaboateng14@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	fileName, command string
	strict            bool
)

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
	rootCmd.PersistentFlags().StringVarP(&fileName, "filename", "f", "", "File/Folder to run shell command on")
	rootCmd.Flags().BoolVarP(&strict, "strict", "s", true, "Strict name matching")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func heimdall(command string, fileName string, strict bool) {
	// Get into the user documents directory
	root, _ := os.UserHomeDir()
	dir := filepath.Join(root, "Documents")
	os.Chdir(dir)

	files := findFiles(fileName, strict)
	if len(files) == 0 {
		fmt.Printf("'%s' not Found\n", fileName)
		return
	}
	f := pickFile(files)
	runCommand(command, f)
}

// findFiles returns list of files and folders whose name matches the string fname.
// When the parameter strict is true, it returns files/folders with the exact match as fname.
// If strict is false, it matches files/folders irrespective of case and extension.
func findFiles(fname string, strict bool) []string {
	var result []string
	err := filepath.WalkDir(".", func(path string, dir fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Matches exact filenames when strict flag is true
		found := filepath.Base(path) == fname

		if !strict {
			// Matches filenames ignoring file extensions and cases when the strict flag is False
			file := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
			fnameNoExt := strings.Split(fname, ".")[0]
			found = strings.EqualFold(file, fnameNoExt)
		}

		if found {
			result = append(result, path)
		}
		return nil
	})

	if err != nil {
		return result
	}
	return result
}

// pickFile returns file among a list of files/folders according to a chosen option.
// If only one file is passed in the files parameter, that file will be returned.
func pickFile(files []string) string {
	if len(files) == 1 {
		return files[0]
	}
	fmt.Println("Choose option: ")
	for i, f := range files {
		fmt.Printf("(%d)\t%s\n", i+1, f)
	}

	var input int
	fmt.Scanln(&input)

	if input > len(files) || input < 0 {
		fmt.Printf("Invalid option: '%d'. Choose between range 1 - %d\n", input, len(files))
		return pickFile(files)
	}
	return files[input-1]
}

// runCommand runs shell commands provided by the parameter `command` on a file.
func runCommand(command string, file string) {
	ctx := context.Background()
	cmd := exec.CommandContext(ctx, command, file)
	cmd.Stdin = os.Stdout
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR: Command couldn't run on file.\nCommand may be running on wrong file or program does not support.", command)
		return
	}
}
