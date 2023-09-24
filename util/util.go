package util

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// FindFiles returns list of files and folders whose name matches the string fname.
// When the parameter strict is true, it returns files/folders with the exact match as fname.
// If strict is false, it matches files/folders irrespective of case and extension.
func FindFiles(fname string, strict bool) []string {
	var result []string
	err := filepath.WalkDir(".", func(path string, dir fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		file := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
		found := strings.EqualFold(file, fname)
		if strict {
			found = filepath.Base(path) == fname
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

// PickFile returns name of a file/folder among a list of files/folders.
// Returns the only name of the file/folder if only one file is returned
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

	for input > len(files) || input < 0 {
		fmt.Printf("Invalid option: '%d'. Choose between range 1 - %d\n", input, len(files))
	}
	return files[input-1]
}

// RunCommand runs shell commands provided by the parameter `command` on a file.
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
