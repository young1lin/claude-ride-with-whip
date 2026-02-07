//go:build windows

package main

import (
	"syscall"
	"unsafe"
)

// Windows console initialization for UTF-8 and ANSI support
func initConsole() {
	// Set UTF-8 code page
	cpMode := uint16(65001) // CP_UTF8
	syscall.Syscall(
		syscall.MustLoadDLL("kernel32.dll").MustFindProc("SetConsoleOutputCP").Addr(),
		1,
		uintptr(cpMode),
		0,
		0,
	)

	// Enable virtual terminal processing for ANSI escape sequences
	// This is required for ANSI colors to work in Windows console
	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	getStdHandle := kernel32.MustFindProc("GetStdHandle")
	getConsoleMode := kernel32.MustFindProc("GetConsoleMode")
	setConsoleMode := kernel32.MustFindProc("SetConsoleMode")

	const STD_OUTPUT_HANDLE = ^uint32(0) - 11 // -11
	const ENABLE_VIRTUAL_TERMINAL_PROCESSING = 0x0004

	stdout, _, _ := getStdHandle.Call(uintptr(STD_OUTPUT_HANDLE))
	var consoleMode uint32
	getConsoleMode.Call(stdout, uintptr(unsafe.Pointer(&consoleMode)))
	setConsoleMode.Call(stdout, uintptr(consoleMode|ENABLE_VIRTUAL_TERMINAL_PROCESSING))
}
