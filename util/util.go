package util

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)


func RunCommand(command string, files []string) {
	switch numFiles := len(files); {
	case numFiles == 0:
		fmt.Println("File not found")
		return

	case numFiles == 1:
		cmd := exec.Command(command, files[0])
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			fmt.Println("ERROR: Command couldn't run on file.\nCommand may be running on wrong file or program does not support.", command)
			return
		}

	default:
		fmt.Println("Choose option: ")
		for i, f := range files {
			fmt.Printf("%d.\t%s\n", i+1, f)
		}

		var input int
		fmt.Scanln(&input)

		if input > len(files) || input < 0 {
			fmt.Printf("Invalid option: '%d'. Choose between range 1 - %d\n", input, len(files))
			RunCommand(command, files)
			return
		}
		// Put the chosen file into an array to be passed to RunCommand function
		files := []string{files[input-1]}
		RunCommand(command, files)
		return
	}

}

func FindFiles(fname string) []string {
	var result []string
	err := filepath.WalkDir(".", func(path string, dir fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if filepath.Base(path) == fname {
			result = append(result, path)
		}
		return nil
	})
	if err != nil {
		return result
	}

	return result
}
