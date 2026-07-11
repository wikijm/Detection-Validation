package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func filewrite(fpath string, binPath string) error {
	if fpath == "" {
		return fmt.Errorf("file path cannot be empty")
	}

	if binPath != "" {
		directory := filepath.Dir(binPath)
		fileName := filepath.Base(binPath)

		copyBinaryTo(directory, fileName)

		args := []string{"createfile", "--path", fpath}
		output := execute(binPath, args)
		fmt.Println(output)
		return nil
	}

	_, err := os.Stat(fpath)
	if os.IsNotExist(err) {
		dir, file := filepath.Split(fpath)
		log.Printf("Creating file: %s in directory: %s\n", file, dir)

		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}

		fileHandle, err := os.Create(fpath)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %v", fpath, err)
		}
		defer fileHandle.Close()

		content := "Test file\ncreated by https://github.com/alwashali/Detection-Validation\n"
		_, err = fileHandle.WriteString(content)
		if err != nil {
			return fmt.Errorf("failed to write to file %s: %v", fpath, err)
		}

		fmt.Printf("Successfully created file: %s\n", fpath)
		return nil
	} else if err != nil {
		return fmt.Errorf("error checking file %s: %v", fpath, err)
	}

	log.Printf("File %s already exists, skipping creation\n", fpath)
	return nil
}
