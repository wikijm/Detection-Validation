package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var defaultTools = map[string][]string{
	"mimikatz": {
		"a2vyymvyb3m6omxpc3qgl2v4cg9yda==",
		"bhnhzhvtcdo6c2ft",
		"bhnhzhvtcdo6c2vjcmv0cw==",
		"bhnhzhvtcdo6y2fjagu=",
		"bwlzyzo6c2njbq==",
		"c2vrdxjsc2e6omxvz29ucgfzc3dvcmrz",
		"chjpdmlszwdlojpkzwj1zw==",
		"chjpdmlszwdlojpkzwj1zyxzzwt1cmxzyto6bg9nb25wyxnzd29yzhm=",
		"crypto::",
		"dg9rzw46omvszxzhdgu=",
		"dg9rzw46onjldmvyda==",
		"dhm6omxvz29ucgfzc3dvcmrz",
		"dmf1bhq6omnyzwq=",
		"dmf1bhq6omxpc3q=",
		"kcq2pa06qpxv86au",
		"kerberos::",
		"lsadump::",
		"misc::sccm",
		"privilege::",
		"privilege::debug,sekurlsa::logonpasswords",
		"sekurlsa::",
		"token::",
		"ts::logonpasswords",
		"ts::mstsc",
		"vault::",
		"y0hkcgrtbhnav2rst2pwa1pxsjfaexh6wld0mwntehpzvg82ykc5bmiynxdzwe56zdi5evpitt0",
		"y0hkcgrtbhnav2rst2pwa1pxsjfaexh6wld0mwntehpzvg82ykc5bmiynxdzwe56zdi5evpitt0=",
		"y3j5chrvojpjbmc=",
		"y3j5chrvojpjyxbp",
		"y3j5chrvojpjzxj0awzpy2f0zxmgl2v4cg9yda==",
		"y3j5chrvojprzxlzic9lehbvcnq=",
		"y3j5chrvojprzxlzic9tywnoaw5lic9lehbvcnq=",
		"yldsell6bzzjmk5qyle9pq==",
		"zehnnk9tehzamjl1y0dgemmzzhzjbvj6",
	},
}

func SimulateTool(toolName string, filePath string, parent string, copyPath string, delay int) error {
	var commands []string

	if filePath != "" {
		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("failed to open file %s: %w", filePath, err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line != "" {
				commands = append(commands, line)
			}
		}
		if err := scanner.Err(); err != nil {
			return fmt.Errorf("error reading file %s: %w", filePath, err)
		}
	} else if toolName != "" {
		cmds, ok := defaultTools[strings.ToLower(toolName)]
		if !ok {
			return fmt.Errorf("unknown tool: %s", toolName)
		}
		commands = cmds
	} else {
		return fmt.Errorf("either --tool or --file must be specified")
	}

	if len(commands) == 0 {
		return fmt.Errorf("no commands to execute")
	}

	for i, cmd := range commands {
		fmt.Printf("[%d/%d] Simulating execution of string: %s\n", i+1, len(commands), cmd)
		_, err := ExecuteCommand(cmd, "", parent, copyPath)
		if err != nil {
			log.Printf("Error executing command '%s': %v\n", cmd, err)
		}

		if i < len(commands)-1 && delay > 0 {
			time.Sleep(time.Duration(delay) * time.Second)
		}
	}

	fmt.Println("Simulation completed.")
	return nil
}
