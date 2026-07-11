package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/crypto/openpgp"
)

// EncryptFiles traverses the file system starting from the given root folder
// and encrypts all files that match the given filename pattern.
// Each encrypted file will be prepended with "enc_".
// Files that are larger than maxSize MB will be skipped.
func EncryptFiles(rootFolder, filenamePattern, password string, maxSize int64) error {
	if rootFolder == "" {
		return fmt.Errorf("root folder cannot be empty")
	}

	if filenamePattern == "" {
		filenamePattern = "*"
	}

	if password == "" {
		password = "default_password_change_me"
		log.Println("Warning: Using default password. Consider setting a custom password.")
	}

	if maxSize <= 0 {
		maxSize = 2 * 1024 * 1024 // 2 MB default
	} else {
		maxSize = maxSize * 1024 * 1024 // Convert MB to bytes
	}

	fileCount := 0
	encryptedCount := 0
	skippedCount := 0

	log.Printf("Starting encryption in folder: %s\n", rootFolder)
	log.Printf("Pattern: %s, Max file size: %d MB\n", filenamePattern, maxSize/(1024*1024))

	err := filepath.Walk(rootFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %s: %v\n", path, err)
			return nil // Continue walking despite errors
		}

		if info.IsDir() {
			return nil
		}

		fileCount++

		if info.Size() > maxSize {
			log.Printf("Skipping large file: %s (%d bytes)\n", path, info.Size())
			skippedCount++
			return nil
		}

		matched, err := filepath.Match(filenamePattern, info.Name())
		if err != nil {
			log.Printf("Error matching pattern for %s: %v\n", path, err)
			return nil
		}

		if matched {
			inFile, err := os.Open(path)
			if err != nil {
				log.Printf("Error opening file %s: %v\n", path, err)
				return nil
			}
			defer inFile.Close()

			outPath := filepath.Join(filepath.Dir(path), "enc_"+info.Name())
			outFile, err := os.Create(outPath)
			if err != nil {
				log.Printf("Error creating encrypted file %s: %v\n", outPath, err)
				return nil
			}
			defer outFile.Close()

			w, err := openpgp.SymmetricallyEncrypt(outFile, []byte(password), nil, nil)
			if err != nil {
				log.Printf("Error encrypting file %s: %v\n", path, err)
				return nil
			}

			if _, err = io.Copy(w, inFile); err != nil {
				log.Printf("Error copying data for %s: %v\n", path, err)
				w.Close()
				return nil
			}

			w.Close()
			encryptedCount++
			log.Printf("Successfully encrypted: %s -> %s\n", path, outPath)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to traverse directory: %v", err)
	}

	fmt.Printf("\nEncryption Summary:\n")
	fmt.Printf("  Files scanned: %d\n", fileCount)
	fmt.Printf("  Files encrypted: %d\n", encryptedCount)
	fmt.Printf("  Files skipped: %d\n", skippedCount)

	return nil
}
