# Claude Ride With Whip - Statusline Plugin

A fun Claude Code statusline plugin that displays an animated red galloping horse in the status bar. The horse moves along a dotted path and cycles through 8 animation frames to create a galloping effect.

## Project Overview

- **Language**: Go (pure Go, no external dependencies)
- **Purpose**: Claude Code statusline plugin for visual entertainment
- **Features**:
  - Animated horse ASCII art (8 frames, 250ms cycle)
  - Horse gallops from right to left along dotted path
  - Red color output for horse sprite
  - Standalone animation mode for testing
  - Debug mode for tracking call timing
  - Cross-platform support (Windows, macOS, Linux)

## Project Structure

```
claude-ride-with-whip/
├── cmd/
│   └── statusline/
│       ├── main.go              # Entry point, JSON parsing, status line rendering
│       ├── frames.go            # Horse frame definitions (8 frames)
│       ├── console.go           # Console initialization (shared)
│       ├── console_windows.go   # Windows-specific console setup
│       └── console_unix.go      # Unix-specific console setup (if added)
├── Makefile                     # Build automation
├── VERSION                      # Current version
└── .claude/
    └── CLAUDE.md                # This file
```

## Building

### Prerequisites

- Go 1.21 or later
- Make (Windows: use WSL or MinGW, or run go commands directly)

### Build Commands

```bash
# Build for current platform
make build

# Output: statusline.exe (Windows) or statusline (Unix)
```

### All Available Make Targets

```bash
make build        # Build the binary for the current platform
make clean        # Remove build artifacts
make test         # Run tests with coverage report
make lint         # Run golangci-lint
make fmt          # Format code with go fmt and gofmt
make install      # Build and install to ~/.claude
make build-all    # Build for all platforms (Windows, macOS, Linux)
make release      # Create a new release (requires gh CLI)
make help         # Show all available commands
```

### Cross-Platform Build

```bash
make build-all

# Output in dist/:
#   statusline_windows_amd64.exe
#   statusline_darwin_amd64
#   statusline_darwin_arm64
#   statusline_linux_amd64
```

## Testing

### Run Tests

```bash
make test
```

This runs:
- All tests with race detection
- Generates coverage.out (raw data)
- Generates coverage.html (HTML report)

### Test Coverage

View the HTML coverage report:
```bash
# Open coverage.html in a browser
start coverage.html  # Windows
open coverage.html   # macOS
xdg-open coverage.html  # Linux
```

### Standalone Animation Test

Test the animation without Claude Code:

```bash
# After building
.\statusline.exe --animate

# Or on Unix
./statusline --animate
```

Press Ctrl+C to exit the animation demo.

### Debug Mode

Enable debug logging to track call timing:

```bash
# With Claude Code: add debug flag to settings.json
# Or standalone test:
echo '{"cwd":"C:\\Project","model":{"display_name":"Test"}}' | .\statusline.exe --debug
```

Debug log location:
- Windows: `C:\Users\<Username>\AppData\Local\Temp\claude_statusline_debug.log`
- Unix: `/tmp/claude_statusline_debug.log`

## Installation

### Install to Claude Code (Manual)

```bash
# Build and copy to Claude config directory
make install

# This copies the binary to:
#   Windows: %USERPROFILE%\.claude\statusline.exe
#   Unix: ~/.claude/statusline
```

### Configure Claude Code

Add to your `~/.claude/settings.json`:

**Windows (note the four backslashes for path escaping):**
```json
{
  "statusLine": {
    "type": "command",
    "command": "C:\\\\Users\\\\<YourUsername>\\\\.claude\\\\statusline.exe"
  }
}
```

**Unix:**
```json
{
  "statusLine": {
    "type": "command",
    "command": "/home/<yourusername>/.claude/statusline"
  }
}
```

Or use the install path from the Makefile which handles OS detection.

## Command-Line Options

```
statusline [flags]

Flags:
  -h, --help     Show help message
  -v, --version  Show version information
  -a, --animate  Run continuous animation in terminal (for testing)
  -d, --debug    Enable debug logging
```

## Animation Details

- **Frame cycle**: 250ms per frame (8 frames = 2 second loop)
- **Position**: 500ms per step (horse moves right to left)
- **Colors**: Horse sprite rendered in red (ANSI color 160)
- **Width**: Path width adapts to terminal width (typically 60-80 chars)

## Code Formatting

```bash
# Format all Go code
make fmt

# This runs:
#   go fmt ./...
#   gofmt -w -s .
```

## Linting

```bash
# Requires golangci-lint to be installed
make lint
```

Install golangci-lint:
```bash
# Windows (PowerShell)
 scoop install golangci-lint

# macOS
 brew install golangci-lint

# Linux
 curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```

## Version

Current version is stored in the `VERSION` file:
```
0.1.0
```

Update version for release:
```bash
make release
```

## Development Notes

- Windows console mode is set to support ANSI escape sequences
- Input JSON is optional - the horse displays even without valid input
- The animation timing is based on `time.Now()` for continuous motion
- Debug mode saves state to temp directory for troubleshooting

## Fun Facts

The horse animation consists of 8 frames that cycle to create a galloping effect. The red color represents energy and passion in Chinese culture ("中国红" - China red).
