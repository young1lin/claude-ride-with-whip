// Package main is the entry point for the claude-ride-with-whip statusline plugin
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// StatusLineInput represents the JSON input from Claude Code
type StatusLineInput struct {
	Model struct {
		DisplayName string `json:"display_name"`
		ID          string `json:"id"`
	} `json:"model"`
	ContextWindow struct {
		TotalInputTokens  int `json:"total_input_tokens"`
		TotalOutputTokens int `json:"total_output_tokens"`
		ContextWindowSize int `json:"context_window_size"`
		CurrentUsage      struct {
			InputTokens          int `json:"input_tokens"`
			OutputTokens         int `json:"output_tokens"`
			CacheReadInputTokens int `json:"cache_read_input_tokens"`
			CacheCreationTokens  int `json:"cache_creation_tokens"`
		} `json:"current_usage"`
	} `json:"context_window"`
	TranscriptPath string `json:"transcript_path"`
	Cwd            string `json:"cwd"`
	Workspace      struct {
		CurrentDir string `json:"current_dir"`
		ProjectDir string `json:"project_dir"`
	} `json:"workspace"`
	RateLimit struct {
		Remaining int `json:"remaining"`
		Limit     int `json:"limit"`
	} `json:"rate_limit"`
}

func main() {
	// Parse command line flags
	args := os.Args[1:]
	animateMode := false
	debugMode := false
	for _, arg := range args {
		if arg == "--version" || arg == "-v" {
			fmt.Println("claude-ride-with-whip statusline v0.1.0")
			os.Exit(0)
		}
		if arg == "--help" || arg == "-h" {
			printHelp()
			os.Exit(0)
		}
		if arg == "--animate" || arg == "-a" {
			animateMode = true
		}
		if arg == "--debug" || arg == "-d" {
			debugMode = true
		}
	}

	// Get debug log file path
	var debugFile *os.File
	if debugMode {
		debugFilePath := getDebugFilePath()
		var err error
		debugFile, err = os.OpenFile(debugFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			// Fail silently for debug mode issues
			debugFile = nil
		}
		defer func() {
			if debugFile != nil {
				debugFile.Close()
			}
		}()
	}

	// Initialize console for Windows
	if runtime.GOOS == "windows" {
		initConsole()
	}

	// Animation mode: continuous animation in terminal
	if animateMode {
		runAnimationMode()
		return
	}

	// Read all input from stdin
	inputBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		// Silent fail for invalid input
		os.Exit(0)
	}

	// Trim null bytes
	inputBytes = trimNullBytes(inputBytes)
	if len(inputBytes) == 0 {
		// No input, still render the horse
		renderStatusLineMulti(nil, debugFile)
		return
	}

	// Try to parse JSON (optional for this plugin)
	var input StatusLineInput
	_ = json.Unmarshal(inputBytes, &input)

	// Render status line (multi-line output) - always show the horse
	renderStatusLineMulti(&input, debugFile)
}

// trimNullBytes removes null bytes from input
func trimNullBytes(data []byte) []byte {
	result := make([]byte, 0, len(data))
	for _, b := range data {
		if b != 0 {
			result = append(result, b)
		}
	}
	return result
}

// ANSI color constants
const (
	colorReset  = "\x1b[0m"
	colorRed160 = "\x1b[38;5;160m" // China red
	colorClear  = "\x1b[2J\x1b[H"  // Clear screen and move cursor to home
)

// renderStatusLineMulti renders the status line with multi-line output
// Colors the horse sprite area red, dots remain default
func renderStatusLineMulti(input *StatusLineInput, debugFile *os.File) {
	now := time.Now()
	horse := getHorseLines(debugFile, now)

	// Find maximum sprite width across all frames
	maxSpriteWidth := 0
	for _, frame := range HorseSprite {
		for _, line := range frame {
			// Use terminal cell width (emoji = 2 cells, not 1)
			lineWidth := StringWidth(line)
			if lineWidth > maxSpriteWidth {
				maxSpriteWidth = lineWidth
			}
		}
	}

	for _, line := range horse {
		// Process each line with per-character color logic (like animation mode)
		// Dots and spaces remain default color, other characters are red
		var result strings.Builder
		inColor := false

		for _, ch := range line {
			if ch == '.' || ch == ' ' {
				if inColor {
					result.WriteString(colorReset)
					inColor = false
				}
				result.WriteRune(ch)
			} else {
				if !inColor {
					result.WriteString(colorRed160)
					inColor = true
				}
				result.WriteRune(ch)
			}
		}

		// Reset color at end if needed
		if inColor {
			result.WriteString(colorReset)
		}

		fmt.Println(result.String())
	}
}

// getHorseLines returns the current frame of the horse animation
// The horse moves right to left along a dotted path
func getHorseLines(debugFile *os.File, now time.Time) []string {
	// Frame animation: 250ms per frame
	frameIndex := int(now.UnixMilli()/250) % NumFrames()
	frame := HorseFrames[frameIndex]
	sprite := GetHorseSprite(frameIndex)

	// Use terminal cell width for frame width (all lines have 95 cells)
	frameWidth := StringWidth(frame[0])

	// Position animation: move right to left
	// Starts at right, moves left, then wraps
	maxPos := frameWidth - 20 // Leave space for horse width
	position := (maxPos - (int(now.UnixMilli()/500) % maxPos))

	// Debug logging
	if debugFile != nil {
		lastCallTime, _, _ := loadLastCallState()
		timeSinceLastCall := now.Sub(lastCallTime)

		debugMsg := fmt.Sprintf(
			"[%s] frame=%d/%d position=%d/%d time_since_last=%v\n",
			now.Format("2006-01-02 15:04:05.000"),
			frameIndex,
			NumFrames(),
			position,
			maxPos,
			timeSinceLastCall,
		)
		debugFile.WriteString(debugMsg)

		// Save current state
		saveLastCallState(now, frameIndex, position)
	}

	// Create result with dotted path
	result := make([]string, 4)

	// Build each row: dots + horse sprite + dots
	for i := 0; i < 4; i++ {
		var row strings.Builder

		// Add leading dots
		for j := 0; j < position; j++ {
			row.WriteByte('.')
		}

		// Add horse sprite if this row has sprite content
		spriteRowIndex := i - 1 // Sprite starts at row 1
		if spriteRowIndex >= 0 && spriteRowIndex < len(sprite) {
			// Add sprite content character by character (handles emoji correctly)
			spriteLine := sprite[spriteRowIndex]
			spriteRunes := []rune(spriteLine)
			for _, r := range spriteRunes {
				row.WriteRune(r)
			}

			// Fill remaining space with dots to maintain visual alignment
			// All lines have 95 terminal cells (visual width is consistent)
			spriteLineWidth := StringWidth(spriteLine)
			remaining := frameWidth - position - spriteLineWidth
			if remaining > 0 {
				for j := 0; j < remaining; j++ {
					row.WriteByte('.')
				}
			}
		} else {
			// No sprite content, fill rest with dots
			remaining := frameWidth - position
			for j := 0; j < remaining; j++ {
				row.WriteByte('.')
			}
		}

		result[i] = row.String()
	}

	return result
}

func basename(path string) string {
	// Handle both Unix and Windows paths
	path = strings.ReplaceAll(path, "\\", "/")
	parts := strings.Split(path, "/")
	if len(parts) > 0 {
		last := parts[len(parts)-1]
		if last != "" {
			return last
		}
		if len(parts) > 1 {
			return parts[len(parts)-2]
		}
	}
	return path
}

func printHelp() {
	fmt.Printf(`claude-ride-with-whip statusline v0.1.0

A Claude Code statusline plugin that displays an animated red galloping horse.

Usage:
  statusline [flags]

Flags:
  -h, --help     Show this help message
  -v, --version  Show version information
  -a, --animate  Run continuous animation in terminal (press Ctrl+C to exit)
  -d, --debug    Enable debug logging to track call timing and animation state

This plugin reads JSON input from stdin and outputs a red horse ASCII art.
The horse animation cycles through 8 frames to create a galloping effect.

Animation timing:
  - Frame cycle: 250ms per frame
  - Position: 500ms per step
`)
}

// runAnimationMode runs continuous animation in the terminal
func runAnimationMode() {
	// Clear screen once at start
	fmt.Print(colorClear)

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		// Clear screen and move cursor to home
		fmt.Print(colorClear)

		// Display header
		fmt.Println(colorRed160 + "╔════════════════════════════════════════════════════════════╗" + colorReset)
		fmt.Println(colorRed160 + "║" + colorReset + "         🐴 Claude Ride With Whip - Animation Demo 🐴         " + colorRed160 + "║" + colorReset)
		fmt.Println(colorRed160 + "║" + colorReset + "                  Press Ctrl+C to exit                    " + colorRed160 + "║" + colorReset)
		fmt.Println(colorRed160 + "╚════════════════════════════════════════════════════════════╝" + colorReset)
		fmt.Println()

		// Render horse
		horse := getHorseLines(nil, time.Now())
		for _, line := range horse {
			var result strings.Builder
			inColor := false

			for _, ch := range line {
				if ch == '.' || ch == ' ' {
					if inColor {
						result.WriteString(colorReset)
						inColor = false
					}
					result.WriteRune(ch)
				} else {
					if !inColor {
						result.WriteString(colorRed160)
						inColor = true
					}
					result.WriteRune(ch)
				}
			}

			if inColor {
				result.WriteString(colorReset)
			}

			fmt.Println(result.String())
		}

		fmt.Println()
		fmt.Println(colorRed160 + "✨ 马到成功 · 一马当先 · 龙马精神 ✨" + colorReset)
	}
}

// Debug state file path
const debugStateFile = "claude_statusline_debug_state.json"

// getDebugFilePath returns the path to the debug log file
func getDebugFilePath() string {
	// Use temp directory
	cacheDir := os.TempDir()
	return filepath.Join(cacheDir, "claude_statusline_debug.log")
}

// CallState stores the last call state for debug mode
type CallState struct {
	LastCallTime int64 `json:"last_call_time"`
	FrameIndex   int   `json:"frame_index"`
	Position     int   `json:"position"`
}

// loadLastCallState loads the previous call state from disk
func loadLastCallState() (time.Time, int, int) {
	cacheDir := os.TempDir()
	statePath := filepath.Join(cacheDir, debugStateFile)

	data, err := os.ReadFile(statePath)
	if err != nil {
		// First run or state file doesn't exist
		return time.Time{}, 0, 0
	}

	var state CallState
	if err := json.Unmarshal(data, &state); err != nil {
		return time.Time{}, 0, 0
	}

	return time.UnixMilli(state.LastCallTime), state.FrameIndex, state.Position
}

// saveLastCallState saves the current call state to disk
func saveLastCallState(now time.Time, frameIndex, position int) {
	cacheDir := os.TempDir()
	statePath := filepath.Join(cacheDir, debugStateFile)

	state := CallState{
		LastCallTime: now.UnixMilli(),
		FrameIndex:   frameIndex,
		Position:     position,
	}

	data, _ := json.Marshal(state)
	_ = os.WriteFile(statePath, data, 0644)
}
