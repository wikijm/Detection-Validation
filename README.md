# Detection-Validation

## Purpose

The tool automates the process of simulating malicious process events without need to go through setup of real processes. 

Suppose you want to test w3wp.exe spawning Powershell, you will need to go through iis setup to simulate w3wp.exe events, which is a consuming task if you have many rules to validate. Since detection engines work based simple string matching from telemetry collection tools such as Sysmon or EDR, any binary with the same parent process name, child process name, commandline and path can be used to test the logic, hence no need to setup iis to simulate the behavior. 

![w3wp_powershell.png](img/w3wp_powershell.png)

The tool allows you to create a child process with a custom parent, child, commandline and path. In addition to couple of other events such as file create from specific process and path, DNS query, registry, and process connections. 

## Features

- ✅ **Process Execution** - Execute commands with custom parent process names
- ✅ **File Creation** - Create files appearing to originate from specific processes
- ✅ **DNS Queries** - Perform DNS lookups from custom binary paths
- ✅ **Network Connections** - Establish connections from spoofed process names
- ✅ **Registry Operations** - Add/delete registry keys with custom process attribution
- ✅ **File Encryption** - Simulate ransomware-like file encryption
- ✅ **File Download** - Download files with custom process attribution
- ✅ **Enhanced Error Handling** - Better error messages and debugging
- ✅ **Flexible Configuration** - Customizable parameters for all operations

```
NAME:
   Malware Cli - A new cli application

USAGE:
   mcli.exe [global options] command [command options] [arguments...]

DESCRIPTION:
   Detection validation tool.
   The objective is to generate event with specific conditions to validate detection rule.
   You can execute commands such as w3wp.exe spawning shell or winword creating file or making DNS queries.

COMMANDS:
   argsfree    Accept any commandline
   connect     Connect to host (with optional --binpath for process spoofing)
   download    Download file (with optional --binpath for process spoofing)
   dnsquery    Resolve DNS (with optional --binpath for process spoofing)
   execute     Execute command with custom commandline and parent process
   encrypt     Encrypt all files in a folder that match a pattern
   createfile  Create file at a specific path (with optional --binpath)
   reg         Add or delete registry key (with optional --binpath)
   simulate    Simulate tool usage by executing specific command strings
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

## Examples

### Process Execution
**winword.exe spawning cscript.exe**  
```powershell
mcli.exe execute --parent winword.exe --command cscript.exe
```

**w3wp.exe spawning powershell with arguments**
```powershell
mcli.exe execute --parent w3wp.exe --command powershell.exe --arg "-enc base64command"
```

### Simulation
**Simulate Mimikatz execution from cutecat.exe**
```powershell
mcli.exe simulate --tool mimikatz --parent cutecat.exe --delay 1
```

**Simulate custom commands from a file**
```powershell
mcli.exe simulate --file commands.txt --parent custom.exe
```

### DNS Queries
**rundll32.exe making DNS request** 
```powershell
mcli.exe dnsquery --host malicious.com --binpath c:\temp\rundll32.exe
```

**Normal DNS query (without process spoofing)**
```powershell
mcli.exe dnsquery --host example.com
```

### File Operations
**w.exe creating file from path C:\temp**  
```powershell
mcli.exe createfile --path c:\users\public\test.dat --binpath c:\temp\w.exe
```

**Create file without process spoofing**
```powershell
mcli.exe createfile --path c:\temp\document.txt
```

### File Encryption
**Encrypt all .txt files in a folder**
```powershell
mcli.exe encrypt --path c:\test --pattern "*.txt" --password "MySecurePass123" --maxsize 5
```

**Encrypt all files (using defaults)**
```powershell
mcli.exe encrypt --path c:\test
```

### Network Operations
**Excel.exe connecting to suspicious host**
```powershell
mcli.exe connect --host 192.168.1.100 --port 4444 --binpath c:\temp\excel.exe
```

**Download file appearing from winword.exe**
```powershell
mcli.exe download --url https://example.com/payload.exe --binpath c:\temp\winword.exe
```

### Registry Operations
**Add registry key with process attribution**
```powershell
mcli.exe reg --keypath "Software\Test" --keyname "TestValue" --value "Data" --binpath c:\temp\regsvr32.exe
```

**Delete registry key**
```powershell
mcli.exe reg --keypath "Software\Test" --keyname "TestValue" --delete
```

## Command Reference

### Global Flags
All commands support `--help` to display usage information.

### Specific Command Options

#### `connect`
- `--host` (required): Target hostname or IP address
- `--port` (required): Target port number
- `--binpath` (optional): Full path of the binary making the connection

#### `download`
- `--url` (required): URL of the file to download
- `--binpath` (optional): Full path of the binary downloading the file

#### `dnsquery`
- `--host` (required): Hostname to resolve
- `--binpath` (optional): Full path of the binary making the DNS query

#### `execute`
- `--command` (required): Command to execute
- `--parent` (optional): Parent process name
- `--arg` (optional): Command arguments
- `--copy` (optional): Path to copy binary before execution (default: C:/Users/Public)

#### `encrypt`
- `--path` (required): Folder path to encrypt files in
- `--pattern` (optional): File name pattern to match (default: *)
- `--password` (optional): Encryption password
- `--maxsize` (optional): Maximum file size in MB to encrypt (default: 2)

#### `createfile`
- `--path` (required): Full path and file name
- `--binpath` (optional): Full path of the binary creating the file

#### `reg`
- `--keypath` (required): Registry key path
- `--keyname` (required): Registry key name
- `--value` (optional): Key value
- `--binpath` (optional): Full path of the binary modifying the registry
- `--delete` (optional): Delete the registry key instead of creating it

#### `simulate`
- `--tool` (optional): The tool to simulate (e.g., mimikatz)
- `--file` (optional): File containing command strings to execute (one per line)
- `--parent` (optional): Parent process name (default: cutecat.exe)
- `--delay` (optional): Delay in seconds between executions (default: 1)
- `--copy` (optional): Path to copy binary before execution (default: C:/Users/Public)

## Recent Enhancements

### v2.0 Improvements
- ✅ **Enhanced Error Handling**: All commands now return meaningful error messages
- ✅ **Process Spoofing for All Network Operations**: Added `--binpath` support to `connect`, `download`, and `dnsquery` commands
- ✅ **Flexible Encryption**: Added customizable password and file size limits for the encrypt command
- ✅ **Better Logging**: Improved console output with operation status and results
- ✅ **Input Validation**: Added validation for required parameters
- ✅ **HTTP Timeouts**: Added timeout handling for download operations
- ✅ **Encryption Summary**: Detailed statistics after encryption operations
- ✅ **Continuous Operation**: Encryption now continues even if individual files fail
- ✅ **Fixed Typos**: Corrected "spcific" to "specific" in help text

## Installation

### Download Pre-built Binary
Download the latest release from the [Releases](https://github.com/wikijm/detection-validation/releases) page. The binary is signed with Sigstore for verification.

### Build from Source

**Windows (Powershell)**

**Run app to download prerequisites and check execution**
```powershell
go run .
```

**Compile app**
```powershell
go build -o mcli.exe .
```
