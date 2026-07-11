package main

import (
	"fmt"
	"log"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

func AddRegistryKey(keyPath string, keyName string, keyValue interface{}, binPath string, delete bool) error {
	if keyPath == "" || keyName == "" {
		return fmt.Errorf("registry key path and name cannot be empty")
	}

	if binPath != "" {
		directory := filepath.Dir(binPath)
		fileName := filepath.Base(binPath)

		copyBinaryTo(directory, fileName)

		args := []string{"reg", "--keyname", keyName, "--keypath", keyPath}
		if keyValue != nil && keyValue.(string) != "" {
			args = append(args, "--value", keyValue.(string))
		}
		if delete {
			args = append(args, "--delete")
		}
		output := execute(binPath, args)
		fmt.Println(output)
		return nil
	}

	key, err := registry.OpenKey(registry.CURRENT_USER, keyPath, registry.ALL_ACCESS)
	if err != nil {
		return fmt.Errorf("failed to open registry key %s: %v", keyPath, err)
	}
	defer key.Close()

	if delete {
		if keyValue != nil && keyValue.(string) != "" {
			return fmt.Errorf("option --value and --delete cannot be used together")
		}
		if err = key.DeleteValue(keyName); err != nil {
			return fmt.Errorf("failed to delete registry value %s\\%s: %v", keyPath, keyName, err)
		}
		log.Printf("Registry key deleted successfully: %s\\%s\n", keyPath, keyName)
		return nil
	}

	valueStr := ""
	if keyValue != nil {
		valueStr = keyValue.(string)
	}

	if err = key.SetStringValue(keyName, valueStr); err != nil {
		return fmt.Errorf("failed to set registry value %s\\%s: %v", keyPath, keyName, err)
	}

	log.Printf("Registry key added successfully: %s\\%s = %s\n", keyPath, keyName, valueStr)
	return nil
}
