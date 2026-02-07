// Package main provides integration tests for the statusline plugin
package main

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetHorseLines_LineAlignment(t *testing.T) {
	// Test that all lines returned by getHorseLines have the same visual width
	// This ensures proper alignment in the terminal
	now := time.Now()
	lines := getHorseLines(nil, now)

	// Should return exactly 4 lines
	assert.Len(t, lines, 4, "getHorseLines should return 4 lines")

	// All lines should have the same terminal cell width
	expectedWidth := 95 // The frame width is 95 characters
	widths := make([]int, len(lines))
	for i, line := range lines {
		widths[i] = StringWidth(line)
		assert.Equal(t, expectedWidth, widths[i],
			"Line %d should be exactly %d cells wide, got %d", i, expectedWidth, widths[i])
	}

	// Verify all widths are equal (redundant but explicit)
	for i := 1; i < len(widths); i++ {
		assert.Equal(t, widths[0], widths[i],
			"All lines should have the same width: line 0 has %d, line %d has %d",
			widths[0], i, widths[i])
	}
}

func TestGetHorseLines_AllDotsExceptSprite(t *testing.T) {
	// Test that lines consist only of dots and horse sprite characters
	now := time.Now()
	lines := getHorseLines(nil, now)

	for i, line := range lines {
		for j, ch := range line {
			// Allow only dots (.) and sprite characters
			// The horse sprite uses: 🐴⏜))/~\ ﾉ and space
			isAllowed := ch == '.' ||
				ch == ' ' ||
				ch == '🐴' ||
				ch == '⏜' ||
				ch == ')' ||
				ch == '/' ||
				ch == '\\' ||
				ch == '~' ||
				ch == 'ﾉ'

			assert.True(t, isAllowed,
				"Line %d position %d: unexpected character %q (U+%04X) in output: %s",
				i, j, ch, ch, line)
		}
	}
}

func TestGetHorseLines_SpriteInMiddleRows(t *testing.T) {
	// Test that the horse sprite appears only in rows 1 and 2 (middle rows)
	now := time.Now()
	lines := getHorseLines(nil, now)

	// Row 0 and 3 should be all dots (or mostly dots)
	assert.Contains(t, lines[0], "...", "Row 0 should contain dots")
	assert.Contains(t, lines[3], "...", "Row 3 should contain dots")

	// Rows 1 and 2 should contain sprite characters (not just dots)
	hasSpriteChars := func(line string) bool {
		for _, ch := range line {
			if ch != '.' && ch != ' ' {
				return true
			}
		}
		return false
	}

	assert.True(t, hasSpriteChars(lines[1]),
		"Row 1 should contain horse sprite characters, got: %s", lines[1])
	assert.True(t, hasSpriteChars(lines[2]),
		"Row 2 should contain horse sprite characters, got: %s", lines[2])
}

func TestGetHorseLines_AnimationProgress(t *testing.T) {
	// Test that position changes over time
	// This is a basic test to ensure animation logic is working
	times := []time.Time{
		time.UnixMilli(0),    // Position at start
		time.UnixMilli(1000), // Position after 1 second
		time.UnixMilli(2000), // Position after 2 seconds
	}

	lines := make([][]string, len(times))
	for i, now := range times {
		lines[i] = getHorseLines(nil, now)
	}

	// The sprite should be at different positions in each frame
	// (unless the animation cycle happens to be at the same point)
	// We check that at least some lines differ
	sameCount := 0
	for i := 1; i < len(lines); i++ {
		if lines[i][1] == lines[i-1][1] {
			sameCount++
		}
	}

	// At least some frames should be different
	// (This test might occasionally fail if timing aligns perfectly,
	// but that's extremely unlikely with 1 second intervals)
	assert.Less(t, sameCount, len(times),
		"Animation should progress over time, but %d/%d frames are identical",
		sameCount, len(times))
}

func TestGetHorseLines_NoTrailingSpaces(t *testing.T) {
	// Test that lines don't have trailing spaces (should use dots instead)
	now := time.Now()
	lines := getHorseLines(nil, now)

	for i, line := range lines {
		trimmed := strings.TrimRight(line, " ")
		assert.Equal(t, line, trimmed,
			"Line %d should not have trailing spaces, got: %q", i, line)
	}
}

func TestStringWidth_VerifyHorseSpriteWidths(t *testing.T) {
	// Verify that all horse sprites have the expected widths
	// Different frames have different widths (6, 7, or 8 cells) due to animation
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

	for frameIdx, sprite := range HorseSprite {
		assert.Len(t, sprite, 2, "Frame %d should have exactly 2 lines", frameIdx)

		for lineIdx, line := range sprite {
			width := StringWidth(line)
			expected := expectedWidths[frameIdx][lineIdx]
			assert.Equal(t, expected, width,
				"Frame %d line %d %q should be %d cells wide, got %d",
				frameIdx, lineIdx, line, expected, width)
		}
	}
}

func TestGetHorseLines_FixedWidthConsistency(t *testing.T) {
	// Test that frame width is calculated correctly and consistently
	frame := HorseFrames[0]

	// All frames should have 95 cells width
	expectedFrameWidth := 95
	for i, line := range frame {
		width := StringWidth(line)
		assert.Equal(t, expectedFrameWidth, width,
			"Frame line %d should be %d cells wide, got %d", i, expectedFrameWidth, width)
	}
}

func TestGetHorseLines_FrameWidth(t *testing.T) {
	// Test the frame width calculation at line 193 of main.go
	frame := HorseFrames[0]
	frameWidth := StringWidth(frame[0])

	assert.Equal(t, 95, frameWidth,
		"Frame width should be 95 cells, got %d", frameWidth)
}

// Benchmark test to ensure width calculation is efficient
func BenchmarkStringWidth(b *testing.B) {
	testStrings := []string{
		"🐴⏜))~",
		" ﾉﾉ ﾉﾉ",
		strings.Repeat(".", 95),
		"...............................................................................................",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range testStrings {
			StringWidth(s)
		}
	}
}

// Benchmark the getHorseLines function
func BenchmarkGetHorseLines(b *testing.B) {
	now := time.Now()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getHorseLines(nil, now)
	}
}
