# Detailed Publishing Workflows

This reference provides step-by-step workflows for common plugin publishing scenarios.

## Workflow: Publishing from Existing Project

You have an existing project and want to make it installable as a plugin.

### Step 1: Add SKILL.md to Project Root

Create `SKILL.md` in your project root:

```yaml
---
name: your-project-name
description: What your plugin does and when to use it. Include specific triggers.
---
```

### Step 2: Organize Plugin Resources

Move or create resource directories as needed:

```
your-project/
├── SKILL.md          # Add this
├── scripts/          # Existing scripts, or create new
├── README.md         # Existing user docs
└── src/              # Your existing source code
```

### Step 3: Publish to GitHub

```bash
# If not already a git repo
git init
git add .
git commit -m "Add plugin support"

# Create GitHub repo (same name as project)
# Add remote and push
git remote add origin https://github.com/YOUR_USERNAME/your-project.git
git push -u origin main
```

### Step 4: Test Installation

In Claude Code:

```bash
/plugin marketplace add YOUR_USERNAME/your-project
/plugin install your-project-name
```

## Workflow: Creating Plugin from Scratch

Starting fresh with a new plugin.

### Step 1: Plan Your Plugin

Answer these questions:
- What problem does it solve?
- What should trigger it? (keywords, file types, commands)
- Does it need scripts, references, or assets?

### Step 2: Create Directory Structure

```bash
mkdir my-new-plugin
cd my-new-plugin

# Create SKILL.md
cat > SKILL.md << 'EOF'
---
name: my-new-plugin
description: Brief description with triggers.
---

# My Plugin

## Overview

What this plugin does.

## Usage

How users interact with it.
EOF

# Create optional directories
mkdir -p scripts references assets
```

### Step 3: Add Functionality

**Add scripts** if you need executable code:
```bash
cat > scripts/helper.py << 'EOF'
#!/usr/bin/env python3
"""Helper script for my plugin."""
print("Hello from my plugin!")
EOF
chmod +x scripts/helper.py
```

**Add references** for documentation Claude should read:
```bash
cat > references/api.md << 'EOF'
# API Reference

Detailed API docs here...
EOF
```

**Add assets** for templates or output files:
```bash
# Copy templates, images, boilerplate code
cp -r my-template/* assets/
```

### Step 4: Test Locally

Before publishing, test your plugin:
1. Verify SKILL.md has valid YAML
2. Test scripts execute without errors
3. Check references are well-formatted
4. Verify assets paths are correct

### Step 5: Publish

```bash
git init
git add .
git commit -m "Initial plugin"
git remote add origin https://github.com/YOUR_USERNAME/my-new-plugin.git
git push -u origin main
```

## Workflow: Publishing Multi-Plugin Repository

One repository containing multiple plugins.

### Structure

```
my-plugins/
├── plugins/
│   ├── plugin-a/
│   │   ├── SKILL.md
│   │   └── scripts/
│   ├── plugin-b/
│   │   ├── SKILL.md
│   │   └── references/
│   └── plugin-c/
│       ├── SKILL.md
│       └── assets/
├── README.md
└── LICENSE
```

### Installation

Users install individual plugins:

```bash
/plugin marketplace add YOUR_USERNAME/my-plugins
/plugin install plugin-a
/plugin install plugin-b
```

### Publishing

```bash
# The repository name can differ from plugin names
git remote add origin https://github.com/YOUR_USERNAME/my-plugins.git
git push -u origin main
```

## Workflow: Updating Published Plugin

Make changes to an existing plugin.

### Minor Updates (Documentation, Small Fixes)

1. Edit files locally
2. Commit and push
3. Users get updates automatically on next install

```bash
git add SKILL.md
git commit -m "docs: improve description"
git push
```

### Major Updates (New Features, Breaking Changes)

1. Update SKILL.md version or description
2. Add/update functionality
3. Create GitHub release

```bash
git add .
git commit -m "feat: add new feature"

# Tag release
git tag v1.1.0
git push origin main --tags

# Or use GitHub CLI
gh release create v1.1.0 --notes "Added new feature"
```

### User Update Process

Users update by reinstalling:

```bash
/plugin update my-plugin
# or
/plugin install my-plugin
```

## Workflow: Debugging Installation Issues

### Plugin Not Found

**Symptoms:**
```
Error: Plugin 'my-plugin' not found
```

**Causes & Solutions:**

1. **Repository is private**
   - Go to GitHub repo → Settings → Make public

2. **Wrong repository name**
   - Verify: `SKILL.md name` matches repo name
   - Example: `name: my-plugin` → `github.com/user/my-plugin`

3. **Plugin in subdirectory without proper structure**
   - Ensure SKILL.md is at root or in `plugins/plugin-name/`

### Description Not Triggering

**Symptoms:**
Plugin installs but doesn't activate when expected.

**Solution:**
Improve description in SKILL.md with specific triggers:

```yaml
# Too vague
description: A tool for working with files.

# Better
description: File manipulation toolkit. Use when user asks to "compress files", "merge files", "split archive", or "batch rename files".
```

### Scripts Not Executing

**Symptoms:**
Scripts referenced in SKILL.md don't run.

**Checks:**

1. Script exists in `scripts/` directory
2. File has execute permissions (`chmod +x`)
3. Shebang line present (`#!/usr/bin/env python3`)

### Assets Not Found

**Symptoms:**
Error copying or using asset files.

**Checks:**

1. Assets are in `assets/` directory
2. Paths in SKILL.md are relative to assets/
3. Asset files are committed to git (not in .gitignore)
