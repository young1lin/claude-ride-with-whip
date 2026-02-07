---
name: plugin-publishing-guide
description: Guide for publishing Claude Code binary plugins (Go/Rust/C++) to GitHub. Use when a developer asks "how to publish my plugin", "make my plugin installable", "create setup command", or "share my compiled tool" with others.
---

# Binary Plugin Publishing Guide

Teach developers how to publish compiled Claude Code plugins (Go, Rust, C++) so users can install them with three commands.

## Quick Start

**Three commands for users to install your plugin:**

```bash
/plugin marketplace add YOUR_USERNAME/your-plugin
/YOUR_USERNAME/your-plugin:setup
```

That's it! Re-running `:setup` updates to the latest version.

## What You Need

| Component | Purpose |
|-----------|---------|
| `.claude/commands/setup.md` | Installation & update instructions |
| GitHub Releases | Platform-specific binaries |
| GoReleaser (optional) | Automated releases |

## Step 1: Create Setup Command

Create `.claude/commands/setup.md` in your repository:

```markdown
# Your Plugin - Setup Guide

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
GET https://api.github.com/repos/YOUR_USERNAME/your-plugin/releases/latest
```

Extract `tag_name` (e.g., `v1.0.0`), remove `v` prefix.

### 3. Download Binary

```
https://github.com/YOUR_USERNAME/your-plugin/releases/download/v{VERSION}/{BINARY_NAME}
```

**Binary naming:**
- Windows: `your-plugin_windows_amd64.zip`
- macOS Intel: `your-plugin_darwin_amd64.tar.gz`
- macOS ARM: `your-plugin_darwin_arm64.tar.gz`
- Linux: `your-plugin_linux_amd64.tar.gz`

### 4. Extract and Install

**Windows:**
```bash
unzip your-plugin_windows_amd64.zip -d /tmp/your-plugin
mv /tmp/your-plugin/your-plugin.exe ~/.claude/your-plugin.exe
```

**macOS/Linux:**
```bash
tar -xzf your-plugin_darwin_amd64.tar.gz -C /tmp/your-plugin
mv /tmp/your-plugin/your-plugin ~/.claude/your-plugin
chmod +x ~/.claude/your-plugin
```

### 5. Configure settings.json

Add to `~/.claude/settings.json`:
```json
{
  "statusLine": {
    "type": "command",
    "command": "~/.claude/your-plugin{.exe}",
    "env": {}
  }
}
```

**⚠️ Windows Path Critical Notes:**

In JSON files, Windows paths require **four backslashes** (`\\\\`) to represent a single backslash:

```json
// ❌ WRONG - Single backslash (invalid JSON escape)
"command": "C:\Users\username\.claude\your-plugin.exe"

// ❌ WRONG - Double backslash (still invalid in JSON)
"command": "C:\\Users\\username\\.claude\\your-plugin.exe"

// ✅ CORRECT - Four backslashes
"command": "C:\\\\Users\\\\username\\\\.claude\\\\your-plugin.exe"
```

**Path Separator Summary:**

| Platform | Path Separator | JSON Escape | Example |
|----------|---------------|-------------|---------|
| macOS/Linux | `/` | `/` | `~/.claude/your-plugin` |
| Windows | `\` | `\\\\` | `C:\\\\Users\\\\username\\\\.claude\\\\your-plugin.exe` |

**Recommended: Use `~` expansion** (works on all platforms):
```json
"command": "~/.claude/your-plugin{.exe}"
```

Claude Code automatically expands `~` to the user's home directory.

### 6. Verify

```bash
# Windows
~\.claude\your-plugin.exe --help

# macOS/Linux
~/.claude/your-plugin --help
```

## Troubleshooting

- **Download fails**: Check https://github.com/YOUR_USERNAME/your-plugin/releases
- **Permission denied (macOS/Linux)**: `chmod +x ~/.claude/your-plugin`
- **Not showing**: Verify `settings.json` path, restart Claude Code
```

## Step 2: Configure GoReleaser (Go Projects)

Create `.goreleaser.yaml`:

```yaml
builds:
  - id: your-plugin
    main: ./cmd/your-plugin
    binary: your-plugin
    env:
      - CGO_ENABLED=0
    goos:
      - linux
      - darwin
      - windows
    goarch:
      - amd64
      - arm64
    ignore:
      - goos: windows
        goarch: arm64
    ldflags:
      - -s -w

archives:
  - builds:
      - your-plugin
    format: tar.gz
    format_overrides:
      - goos: windows
        format: zip
    name_template: "{{ .ProjectName }}_{{ .Os }}_{{ .Arch }}"

release:
  github:
    owner: YOUR_USERNAME
    name: your-plugin
```

**For other languages (Rust, C++):** Build binaries manually or use equivalent release automation.

## Step 3: Publish to GitHub

### Option A: GitHub Actions (Recommended)

Create `.github/workflows/release.yml`:

```yaml
name: Release

on:
  push:
    tags:
      - 'v*'

permissions:
  contents: write

jobs:
  release:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.23'  # Adjust to your Go version
          cache: true

      - name: Run GoReleaser
        uses: goreleaser/goreleaser-action@v6
        with:
          version: latest
          args: release --clean
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

**Then publish:**
```bash
git tag v1.0.0
git push origin v1.0.0
```

GitHub Actions automatically builds and releases all platform binaries.

### Option B: Manual Release

```bash
# Initialize git (if new)
git init
git add .
git commit -m "Initial release"

# Create GitHub repository (public)
git remote add origin https://github.com/YOUR_USERNAME/your-plugin.git
git push -u origin main

# Create first release
git tag v1.0.0
git push origin v1.0.0
goreleaser release --clean
```

## Step 4: User Installation

Users install with three commands:

```bash
/plugin marketplace add YOUR_USERNAME/your-plugin
/YOUR_USERNAME/your-plugin:setup
```

Re-running `:setup` updates to the latest version.

## Verification Checklist

Before publishing:

- [ ] `.claude/commands/setup.md` exists
- [ ] `setup.md` has platform detection, download URL, install steps
- [ ] Repository is public
- [ ] GoReleaser configured (if using Go)
- [ ] GitHub release has platform-specific binaries

Test installation:
```bash
/YOUR_USERNAME/your-plugin:setup
```

## Example: Complete Plugin

**Reference:** `young1lin/claude-token-monitor`

**User installation:**
```bash
/plugin marketplace add young1lin/claude-token-monitor
/young1lin/claude-token-monitor:setup
```

**Repository structure:**
```
claude-token-monitor/
├── .claude/
│   └── commands/
│       └── setup.md       # Installation instructions
├── cmd/statusline/
│   └── main.go
├── .goreleaser.yaml
└── README.md
```

**Binary naming (from Goreleaser):**
- `statusline_windows_amd64.zip`
- `statusline_darwin_amd64.tar.gz`
- `statusline_darwin_arm64.tar.gz`
- `statusline_linux_amd64.tar.gz`

## Manual Release (Any Language)

If not using GoReleaser:

```bash
# Build for each platform
go build -o your-plugin_windows_amd64.exe ./cmd/your-plugin
go build -o your-plugin_darwin_amd64 ./cmd/your-plugin
go build -o your-plugin_darwin_arm64 ./cmd/your-plugin
go build -o your-plugin_linux_amd64 ./cmd/your-plugin

# Create archives
zip your-plugin_windows_amd64.zip your-plugin_windows_amd64.exe
tar -czf your-plugin_darwin_amd64.tar.gz your-plugin_darwin_amd64
tar -czf your-plugin_darwin_arm64.tar.gz your-plugin_darwin_arm64
tar -czf your-plugin_linux_amd64.tar.gz your-plugin_linux_amd64

# Create GitHub release and upload
gh release create v1.0.0 \
  your-plugin_*_amd64.zip \
  your-plugin_*_amd64.tar.gz \
  your-plugin_*_arm64.tar.gz
```

## Update Process

**Developer:**
1. Make changes
2. Commit and push: `git add . && git commit -m "feat: new feature" && git push`
3. Tag new version: `git tag v1.1.0 && git push origin v1.1.0`
4. **GitHub Actions**: Automatically builds and releases
5. **Or manually**: `goreleaser release --clean`

**User:**
```bash
/YOUR_USERNAME/your-plugin:setup
```

Re-run setup to get the latest version.

## Resources

See [references/workflows.md](references/workflows.md) for detailed workflows.

See [references/examples.md](references/examples.md) for more examples.
