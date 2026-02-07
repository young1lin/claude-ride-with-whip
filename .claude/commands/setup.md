# Claude Ride With Whip - Setup Guide

## Installation Steps

### 1. Detect Platform

| OS | Detection | File Extension |
|----|-----------|----------------|
| Windows | `os.platform() === 'win32'` | .zip |
| macOS | `os.platform() === 'darwin'` | .tar.gz |
| Linux | `os.platform() === 'linux'` | .tar.gz |

| Arch | Detection |
|------|-----------|
| AMD64/x64 | `process.arch === 'x64'` or `os.arch() === 'amd64'` |
| ARM64 | `process.arch === 'arm64'` or `os.arch() === 'arm64'` |

### 2. Get Latest Version

```
GET https://api.github.com/repos/young1lin/claude-ride-with-whip/releases/latest
```

Extract `tag_name` (e.g., `v1.0.0`), remove `v` prefix.

### 3. Download Binary

```
https://github.com/young1lin/claude-ride-with-whip/releases/download/v{VERSION}/{BINARY_NAME}
```

**Binary naming:**
- Windows: `claude-ride-with-whip_windows_amd64.zip`
- macOS Intel: `claude-ride-with-whip_darwin_amd64.tar.gz`
- macOS ARM: `claude-ride-with-whip_darwin_arm64.tar.gz`
- Linux: `claude-ride-with-whip_linux_amd64.tar.gz`

### 4. Extract and Install

**Windows:**
```bash
unzip claude-ride-with-whip_windows_amd64.zip -d /tmp/claude-ride-with-whip
mv /tmp/claude-ride-with-whip/statusline.exe ~/.claude/claude-ride-with-whip.exe
```

**macOS/Linux:**
```bash
tar -xzf claude-ride-with-whip_darwin_amd64.tar.gz -C /tmp/claude-ride-with-whip
mv /tmp/claude-ride-with-whip/statusline ~/.claude/claude-ride-with-whip
chmod +x ~/.claude/claude-ride-with-whip
```

### 5. Configure settings.json

Add to `~/.claude/settings.json`:
```json
{
  "statusLine": {
    "type": "command",
    "command": "~/.claude/claude-ride-with-whip{.exe}",
    "env": {}
  }
}
```

**⚠️ Windows Path Critical Notes:**

In JSON files, Windows paths require **four backslashes** (`\\\\`) to represent a single backslash:

```json
// ❌ WRONG
"command": "C:\Users\username\.claude\claude-ride-with-whip.exe"

// ✅ CORRECT
"command": "C:\\\\Users\\\\username\\\\.claude\\\\claude-ride-with-whip.exe"
```

**Recommended: Use `~` expansion:**
```json
"command": "~/.claude/claude-ride-with-whip{.exe}"
```

### 6. Verify

```bash
# Windows
~\.claude\claude-ride-with-whip.exe --help

# macOS/Linux
~/.claude/claude-ride-with-whip --help
```

## Features

- Chinese Red Horse ASCII Art Display
- Token usage monitoring
- Git status display
- Project folder name
