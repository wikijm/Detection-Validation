# Changelog

All notable changes to this project will be documented in this file.

## [v2.0.0] - 2025-11-04

### Added
- **Process Spoofing for Network Operations**
  - Added `--binpath` flag to `connect` command for connection attribution
  - Added `--binpath` flag to `download` command for download attribution
  - Added `--binpath` flag to `dnsquery` command for DNS query attribution

- **Enhanced Encryption Features**
  - Added `--password` flag to specify custom encryption password
  - Added `--maxsize` flag to control maximum file size for encryption (in MB)
  - Added encryption summary with statistics (files scanned, encrypted, skipped)
  - Improved error handling to continue encryption even if individual files fail

- **Better Error Handling**
  - All functions now return proper error types
  - Meaningful error messages throughout the application
  - Input validation for all required parameters
  - Graceful error handling with descriptive messages

- **Improved Logging**
  - Added detailed operation logging
  - Success/failure messages for all operations
  - Progress indicators for long-running operations
  - Better visibility into what the tool is doing

- **HTTP Improvements**
  - Added 30-second timeout for download operations
  - HTTP status code validation
  - Better error messages for network failures

### Changed
- **Function Signatures**: All command functions now return `error` type for proper error handling
- **Encryption Behavior**: Changed from failing on first error to continuing through all files
- **DNS Resolution**: Improved output format with better IP address display
- **Registry Operations**: Enhanced parameter handling and error messages

### Fixed
- Fixed typo: "spcific" → "specific" in createfile command usage
- Fixed missing error handling in connect.go
- Fixed improper error handling in encrypt.go
- Fixed file handle management in filecreate.go
- Fixed registry operation return values
- Updated deprecated `ioutil.TempFile` to `os.CreateTemp`

### Security
- Added warning when using default encryption password
- Improved input validation to prevent empty/invalid parameters
- Better error messages that don't leak sensitive information

## [v1.0.0] - Initial Release

### Features
- Process execution with custom parent processes
- File creation with process attribution
- DNS query execution
- Network connections
- Registry key operations
- File encryption capabilities
- Command-line argument handling
