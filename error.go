/*
  Eyes is a lightweight Go logger with colored output using fatih/color.
  Copyright (C) 2025-2026 ApertureOS Team
  Licensed under the MIT license. See LICENSE file in the project root for details.
*/

package eyes

import (
	"fmt"
	"os"
)

// Error prints an error message and doesnt exit.
func Error(arguments ...any) {
	loggerConfig.FatalTextColor.Fprint(
		os.Stderr,
		buildPrefixString("FATAL"),
	)
	fmt.Fprint(os.Stderr, arguments...)
	fmt.Fprintln(os.Stderr)
}

// Errorf prints a formatted error message and doesnt exit
func Errorf(format string, arguments ...any) {
	loggerConfig.FatalTextColor.Fprint(
		os.Stderr,
		buildPrefixString("FATAL"),
	)
	fmt.Fprintf(os.Stderr, format, arguments...)
	fmt.Fprintln(os.Stderr)
}

// Errorln prints an error message followed by a newline
// and doesnt exit
func Errorln(arguments ...any) {
	loggerConfig.FatalTextColor.Fprint(
		os.Stderr,
		buildPrefixString("FATAL"),
	)
	fmt.Fprintln(os.Stderr, arguments...)

}
