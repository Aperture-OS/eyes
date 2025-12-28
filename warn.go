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

// Warn prints a warning message.
func Warn(arguments ...any) {
	loggerConfig.WarnTextColor.Fprint(
		os.Stderr,
		buildPrefixString("WARN"),
	)
	fmt.Fprint(os.Stderr, arguments...)
	fmt.Fprintln(os.Stderr)
}

// Warnf prints a formatted warning message.
func Warnf(format string, arguments ...any) {
	loggerConfig.WarnTextColor.Fprint(
		os.Stderr,
		buildPrefixString("WARN"),
	)
	fmt.Fprintf(os.Stderr, format, arguments...)
	fmt.Fprintln(os.Stderr)
}

// Warnln prints a warning message followed by a newline.
func Warnln(arguments ...any) {
	loggerConfig.WarnTextColor.Fprint(
		os.Stderr,
		buildPrefixString("WARN"),
	)
	fmt.Fprintln(os.Stderr, arguments...)
}
