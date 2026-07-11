package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func init() {
	app.Name = "Malware Cli"

	app.Description = `Detection validation tool. 
	 The objective is to generate events with specific conditions to validate detection rules.
	 You can execute commands such as w3wp.exe spawning a shell, or winword.exe creating a file or making DNS queries.`

}

func main() {
	app.Commands = []cli.Command{
		{
			Name: "argsfree",

			Usage: "Accept any commandline",

			Action: func(c *cli.Context) {

				fmt.Println(os.Args)

			},
		},
		{
			Name: "connect",

			Usage: "Connect to host",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "host",
					Usage:    "hostname or IP Address",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "port",
					Usage:    "port number",
					Required: true,
				},
				&cli.StringFlag{
					Name:  "binpath",
					Usage: "Full path of the binary making the connection. Example: C:/temp/binary.exe",
				},
			},
			Action: func(c *cli.Context) error {
				return connectToHost(c.String("host"), c.String("port"), c.String("binpath"))
			},
		},
		{
			Name: "download",

			Usage: "Download file",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "url",
					Usage:    "File URL",
					Required: true,
				},
				&cli.StringFlag{
					Name:  "binpath",
					Usage: "Full path of the binary downloading the file. Example: C:/temp/binary.exe",
				},
			},
			Action: func(c *cli.Context) error {
				return downloadFile(c.String("url"), c.String("binpath"))
			},
		},
		{
			Name: "dnsquery",

			Usage: "Resolve DNS",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "host",
					Usage:    "hostname to resolve",
					Required: true,
				},
				&cli.StringFlag{
					Name:  "binpath",
					Usage: "Full path of the binary making the DNS query. Example: C:/temp/binary.exe",
				},
			},
			Action: func(c *cli.Context) error {
				return resolve(c.String("host"), c.String("binpath"))
			},
		},
		{
			Name:  "execute",
			Usage: "Execute command with custom commandline and parent process",
			Flags: []cli.Flag{

				&cli.StringFlag{
					Name:     "command",
					Usage:    "Hostname or IP Address",
					Required: true,
				},
				&cli.StringFlag{
					Name:  "parent",
					Usage: "Optional parent command to execute",
				},
				&cli.StringFlag{
					Name:  "arg",
					Usage: "Command arguments",
				},
				&cli.StringFlag{
					Name:  "copy",
					Usage: "Copy to path before execution",
					Value: "C:/Users/Public",
				},
			},
			Action: func(c *cli.Context) {

				ExecuteCommand(c.String("command"), c.String("arg"), c.String("parent"), c.String("copy"))

			},
		}, {
			Name:  "encrypt",
			Usage: "encrypt all files in a folder that match a pattern",
			Flags: []cli.Flag{

				&cli.StringFlag{
					Name:     "path",
					Usage:    "Folder path",
					Required: true,
				},
				&cli.StringFlag{
					Name:  "pattern",
					Usage: "File name pattern (default: *)",
					Value: "*",
				},
				&cli.StringFlag{
					Name:  "password",
					Usage: "Encryption password (default: auto-generated)",
				},
				&cli.Int64Flag{
					Name:  "maxsize",
					Usage: "Maximum file size in MB to encrypt (default: 2)",
					Value: 2,
				},
			},
			Action: func(c *cli.Context) error {
				return EncryptFiles(c.String("path"), c.String("pattern"), c.String("password"), c.Int64("maxsize"))
			},
		},
		{
			Name: "createfile",

			Usage: "Create file at a specific path",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "path",
					Usage:    "full path and file name",
					Required: true,
				},
				&cli.StringFlag{
					Name:  "binpath",
					Usage: "Full path of the binary creating the file. Example: C:/temp/binary.exe",
				},
			},
			Action: func(c *cli.Context) error {
				return filewrite(c.String("path"), c.String("binpath"))
			},
		},
		{
			Name: "reg",

			Usage: "Add or delete registry key",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "keyname",
					Usage:    "Key name",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "keypath",
					Usage:    "Key path",
					Required: true,
				},
				&cli.StringFlag{
					Name:  "value",
					Usage: "Key value",
				},
				&cli.StringFlag{
					Name:  "binpath",
					Usage: "Full path of the binary modifying the registry. Example: C:/temp/binary.exe",
				},
				&cli.BoolFlag{
					Name:  "delete",
					Usage: "Delete key",
				},
			},
			Action: func(c *cli.Context) error {
				return AddRegistryKey(c.String("keypath"), c.String("keyname"), c.String("value"), c.String("binpath"), c.Bool("delete"))
			},
		},
		{
			Name: "simulate",

			Usage: "Simulate tool usage by executing specific command strings",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "tool",
					Usage: "The tool to simulate (e.g., mimikatz)",
				},
				&cli.StringFlag{
					Name:  "file",
					Usage: "Optional file containing command strings (one per line)",
				},
				&cli.StringFlag{
					Name:  "parent",
					Usage: "Optional parent process name",
					Value: "cutecat.exe",
				},
				&cli.IntFlag{
					Name:  "delay",
					Usage: "Delay in seconds between command executions",
					Value: 1,
				},
				&cli.StringFlag{
					Name:  "copy",
					Usage: "Copy to path before execution",
					Value: "C:/Users/Public",
				},
			},
			Action: func(c *cli.Context) error {
				return SimulateTool(c.String("tool"), c.String("file"), c.String("parent"), c.String("copy"), c.Int("delay"))
			},
		},
	}

	//log.Println("Received arguments: ", os.Args)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
