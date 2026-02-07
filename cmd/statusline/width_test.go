// Package main provides tests for terminal width calculation
package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringWidth_Emoji(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "horse emoji only",
			input:    "🐴",
			expected: 2,
		},
		{
			name:     "horse with tilde and parentheses",
			input:    "🐴⏜))~",
			expected: 6, // 🐴(2) + ⏜(1) + )(1) + )(1) + ~(1) = 5 runes, 6 cells
		},
		{
			name:     "sprite line 1 (full width characters)",
			input:    " ﾉﾉ ﾉﾉ",
			expected: 6, // space(1) + ﾉ(1) + ﾉ(1) + space(1) + ﾉ(1) + ﾉ(1) + space(1) = 6
		},
		{
			name:     "ASCII only dots",
			input:    "...",
			expected: 3,
		},
		{
			name:     "mixed ASCII and emoji",
			input:    "..🐴..",
			expected: 6, // ..(2) + 🐴(2) + ..(2) = 6
		},
		{
			name:     "empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "space only",
			input:    " ",
			expected: 1,
		},
		{
			name:     "slash and backslash",
			input:    " / \\ ",
			expected: 5, // space(1) + /(1) + space(1) + \\(1) + space(1) = 5
		},
		{
			name:     "frame 1 sprite line 2",
			input:    " / \\ ﾉﾉ",
			expected: 7, // space(1) + /(1) + space(1) + \\(1) + space(1) + ﾉ(1) + ﾉ(1) = 7
		},
		{
			name:     "frame 2 sprite line 2",
			input:    "  \\\\ //",
			expected: 7, // space(1) + space(1) + \\(1) + \\(1) + space(1) + /(1) + /(1) = 7
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, StringWidth(tt.input),
				"StringWidth(%q) should return %d", tt.input, tt.expected)
		})
	}
}

func TestStringWidth_CompareWithRuneCount(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		runeCount     int
		expectedWidth int
		shouldDiffer  bool
	}{
		{
			name:          "emoji shows difference",
			input:         "🐴",
			runeCount:     1,
			expectedWidth: 2,
			shouldDiffer:  true,
		},
		{
			name:          "ASCII only same",
			input:         "abc",
			runeCount:     3,
			expectedWidth: 3,
			shouldDiffer:  false,
		},
		{
			name:          "horse sprite line 1",
			input:         "🐴⏜))~",
			runeCount:     5,
			expectedWidth: 6,
			shouldDiffer:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualRuneCount := len([]rune(tt.input))
			actualWidth := StringWidth(tt.input)

			assert.Equal(t, tt.runeCount, actualRuneCount,
				"Rune count verification for %q", tt.input)
			assert.Equal(t, tt.expectedWidth, actualWidth,
				"StringWidth(%q) should return %d", tt.input, tt.expectedWidth)

			if tt.shouldDiffer {
				assert.NotEqual(t, actualRuneCount, actualWidth,
					"StringWidth should differ from rune count for %q", tt.input)
			} else {
				assert.Equal(t, actualRuneCount, actualWidth,
					"StringWidth should equal rune count for %q", tt.input)
			}
		})
	}
}

func TestStringWidth_AllHorseSprites(t *testing.T) {
	// Test all horse sprite frames to ensure correct width calculation
	// This helps catch any regressions if sprite data changes
	// Note: Different frames have different widths (6, 7, or 8 cells)
	// Frame 0,4: 6,6 | Frame 1,5: 7,7 | Frame 2,6: 8,7 | Frame 3,7: 7,7
	expectedWidths := map[int][2]int{
		0: {6, 6},
		1: {7, 7},
		2: {8, 7},
		3: {7, 7},
		4: {6, 6},
		5: {7, 7},
		6: {8, 7},
		7: {7, 7},
	}

	for i, sprite := range HorseSprite {
		t.Run(fmt.Sprintf("frame_%d_line_0", i), func(t *testing.T) {
			line0 := sprite[0]
			width := StringWidth(line0)
			expected := expectedWidths[i][0]
			assert.Equal(t, expected, width,
				"Frame %d line 0 %q should be %d cells wide", i, line0, expected)
		})

		t.Run(fmt.Sprintf("frame_%d_line_1", i), func(t *testing.T) {
			line1 := sprite[1]
			width := StringWidth(line1)
			expected := expectedWidths[i][1]
			assert.Equal(t, expected, width,
				"Frame %d line 1 %q should be %d cells wide", i, line1, expected)
		})
	}
}
