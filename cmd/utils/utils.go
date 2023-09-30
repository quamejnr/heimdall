package utils

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

)
// findFiles return list of files and folders whose name matches the string fname.
// When the parameter strict is true, it returns files/folders with the exact match as fname.
// If strict is false, it matches files/folders irrespective of case and extension.
func FindFiles(fname string, strict bool) []string {
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
func PickFile(files []string) string {
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
		return PickFile(files)
	}
	return files[input-1]
}

// runCommand runs shell commands provided by the parameter `command` on a file.
func RunCommand(command string, file string) {
	ctx := context.Background()
	cmd := exec.CommandContext(ctx, command, file)
	cmd.Stdin = os.Stdout
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR: Command couldn't run on file.\nCommand may be running on wrong file or program does not support.", command)
		return
	}
}
