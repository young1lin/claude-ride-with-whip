# Claude Ride With Whip

A fun Claude Code statusline plugin that displays an animated red galloping horse in the status bar. The horse moves along a dotted path and cycles through 8 animation frames to create a galloping effect.

## Overview

- **Language**: Go (pure Go, no external dependencies)
- **Purpose**: Claude Code statusline plugin for visual entertainment
- **Features**:
  - Animated horse ASCII art (8 frames, 250ms cycle)
  - Horse gallops from right to left along dotted path
  - Red color output for horse sprite
  - Standalone animation mode for testing
  - Debug mode for tracking call timing
  - Cross-platform support (Windows, macOS, Linux)

## Quick Start

```bash
# Build
make build

# Test animation
.\bin\statusline.exe --animate

# Run tests
make test
```

## Installation

### Install to Claude Code

```bash
make install
```

### Configure Claude Code

Add to `~/.claude/settings.json` (Windows - note the four backslashes):

```json
{
  "statusLine": {
    "type": "command",
    "command": "C:\\\\Users\\\\<YourUsername>\\\\.claude\\\\statusline.exe"
  }
}
```

## Command-Line Options

```
statusline [flags]

Flags:
  -h, --help     Show help message
  -v, --version  Show version information
  -a, --animate  Run continuous animation in terminal (for testing)
  -d, --debug    Enable debug logging
```

## Make Targets

```bash
make build        # Build for current platform
make clean        # Remove build artifacts
make test         # Run tests with coverage report
make lint         # Run golangci-lint
make fmt          # Format code
make install      # Install to ~/.claude
make build-all    # Build for all platforms
make release      # Create a new release
make help         # Show all commands
```

## License

[License](LICENSE)
