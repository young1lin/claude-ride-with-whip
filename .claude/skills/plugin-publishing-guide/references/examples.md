# Plugin Examples

Collection of real-world plugin examples with analysis of what makes them effective.

## Example 1: Statusline Plugin

**Repository:** `young1lin/claude-token-monitor`

### SKILL.md Analysis

```yaml
---
name: claude-token-monitor
description: Real-time token usage statusline for Claude Code. Use when user asks to "install statusline", "setup token monitor", "enable status bar", or display session statistics.
disable-model-invocation: true
---
```

**Why this works:**
- Clear specific triggers: "install statusline", "setup token monitor"
- Describes functionality: "Real-time token usage"
- Sets `disable-model-invocation: true` - this is a setup tool, not a context reference

### Installation Commands

```bash
/plugin marketplace add young1lin/claude-token-monitor
/plugin install claude-token-monitor
```

### What Users Get

- Automatic token monitoring in status bar
- No additional commands needed after install
- "Install and forget" pattern

## Example 2: PDF Manipulation Plugin

Hypothetical plugin for PDF operations.

### SKILL.md Structure

```yaml
---
name: pdf-toolkit
description: PDF manipulation toolkit for extracting text, filling forms, merging/splitting documents, and converting formats. Use when user asks to "extract PDF text", "fill PDF form", "merge PDFs", "split PDF", or "convert PDF to images".
---
```

### Body Content

```markdown
# PDF Toolkit

## Quick Start

### Extract Text from PDF

```bash
pdf-toolkit:extract-text document.pdf
```

### Fill PDF Form

```bash
pdf-toolkit:fill-form form.pdf --data field_values.json
```

### Merge PDFs

```bash
pdf-toolkit:merge output.pdf file1.pdf file2.pdf file3.pdf
```

## Resources

- [Form Filling Guide](references/forms.md)
- [API Reference](references/api.md)
```

### Why This Structure Works

- **Task-based organization**: Each operation is clearly separated
- **Command examples**: Users can copy/paste commands
- **Progressive disclosure**: Advanced docs in references/

## Example 3: Frontend Builder Plugin

Hypothetical plugin for web development.

### SKILL.md

```yaml
---
name: frontend-builder
description: Create distinctive, production-grade frontend interfaces with high design quality. Use when user asks to "build website", "create landing page", "design dashboard", "build React app", or "create web component".
---
```

### Directory Structure

```
frontend-builder/
├── SKILL.md
├── assets/
│   └── starter-template/    # React/HTML boilerplate
│       ├── package.json
│       ├── src/
│       └── public/
└── references/
    ├── components.md        # Component patterns
    └── styling.md           # Design system guidelines
```

### Key Features

- **Assets with boilerplate**: Provides starter code users can build on
- **Reference docs**: Design patterns loaded only when needed
- **High-freedom description**: Many ways to request frontend work

## Example 4: API Integration Plugin

Hypothetical plugin for working with specific APIs.

### SKILL.md

```yaml
---
name: slack-helper
description: Slack API integration for sending messages, creating channels, managing users, and webhooks. Use when user asks to "send Slack message", "create Slack channel", "invite user to Slack", or "Slack webhook".
---
```

### With Scripts

```
slack-helper/
├── SKILL.md
├── scripts/
│   ├── send_message.py     # Direct API calls
│   ├── create_channel.py
│   └── webhook.py
└── references/
    └── slack_api.md        # API documentation
```

### Why Scripts Make Sense Here

- **Deterministic operations**: API calls require exact syntax
- **Reusable code**: Same API patterns used repeatedly
- **Error handling**: Scripts handle rate limits, retries

## Example 5: Domain-Specific Plugin

Plugin for company-specific knowledge.

### SKILL.md

```yaml
---
name: company-docs
description: Company documentation and standards. Use when referencing company policies, brand guidelines, API schemas, or internal workflows.
---
```

### Directory Structure

```
company-docs/
├── SKILL.md
└── references/
    ├── brand.md            # Brand guidelines
    ├── apis.md             # Internal API docs
    ├── policies.md         # Company policies
    └── workflows.md        # Team workflows
```

### Why References-Only Structure

- **Large documentation**: Company docs are extensive
- **Context-dependent**: Only relevant sections load based on query
- **Frequently updated**: Docs change more often than code

## Example 6: Multi-Plugin Repository

**Repository:** `anthropics/skills` (official example)

### Structure

```
skills/
├── docx/
│   └── SKILL.md
├── pdf/
│   ├── SKILL.md
│   └── scripts/
├── frontend-design/
│   ├── SKILL.md
│   └── assets/
└── project-planner/
    ├── SKILL.md
    └── references/
```

### Why This Works

- **Logical grouping**: Related skills in one place
- **Independent installation**: Users install only what they need
- **Shared maintenance**: Single repo for all official skills

## Description Writing Patterns

### Pattern 1: Verb + Object

```yaml
description: Extract text from PDF documents. Use when user asks to "extract PDF text" or "get PDF content".
```

### Pattern 2: Capability List

```yaml
description: Image processing toolkit supporting resize, crop, rotate, filter, and format conversion. Use when working with images.
```

### Pattern 3: Tool + Triggers

```yaml
description: Database query helper. Use when user asks to "query database", "run SQL", "fetch data", or "database lookup".
```

### Pattern 4: Domain + Operations

```yaml
description: AWS deployment automation for EC2, S3, and Lambda. Use when user asks to "deploy to AWS", "create EC2 instance", "upload to S3", or "update Lambda function".
```

## Common Mistakes

### Mistake 1: Vague Description

```yaml
# Bad - unclear when to use
description: A helpful tool for various tasks.

# Good - specific triggers
description: File compression toolkit. Use when user asks to "compress files", "reduce file size", or "create archive".
```

### Mistake 2: Missing Installation Instructions

```yaml
# SKILL.md should include installation

---

# My Plugin

## Installation

```bash
/plugin marketplace add YOUR_USERNAME/my-plugin
/plugin install my-plugin
```

## Usage
...
```

### Mistake 3: No Progressive Disclosure

```markdown
# Bad - everything in one file
# SKILL.md with 500+ lines of detailed docs

# Good - lean SKILL.md with references
# SKILL.md with overview and links to references/
```

### Mistake 4: Name Mismatch

```yaml
# Bad - repository name doesn't match
# SKILL.md: name: my-cool-plugin
# Repository: github.com/user/stuff

# Good - consistent naming
# SKILL.md: name: my-cool-plugin
# Repository: github.com/user/my-cool-plugin
```

## Plugin Categories

### Category 1: Developer Tools
- Code generation
- API helpers
- Testing utilities
- Deployment automation

### Category 2: Document Processing
- PDF manipulation
- DOCX editing
- Spreadsheet operations
- Format conversion

### Category 3: Design & Creative
- Frontend building
- Image processing
- Template generation
- Brand guidelines

### Category 4: Domain-Specific
- Company knowledge
- Industry-specific tools
- Specialized workflows
- API integrations

### Category 5: Setup & Configuration
- Environment setup
- Plugin installation helpers
- Configuration management
- Status monitoring
