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

// Info prints an informational message using default formatting.
func Info(arguments ...any) {
	loggerConfig.InfoTextColor.Fprint(
		os.Stderr,
		buildPrefixString("INFO"),
	)
	fmt.Fprint(os.Stderr, arguments...)
	fmt.Fprintln(os.Stderr)
}

// Infof prints a formatted informational message.
func Infof(format string, arguments ...any) {
	loggerConfig.InfoTextColor.Fprint(
		os.Stderr,
		buildPrefixString("INFO"),
	)
	fmt.Fprintf(os.Stderr, format, arguments...)
	fmt.Fprintln(os.Stderr)
}

// Infoln prints an informational message followed by a newline.
func Infoln(arguments ...any) {
	loggerConfig.InfoTextColor.Fprint(
		os.Stderr,
		buildPrefixString("INFO"),
	)
	fmt.Fprintln(os.Stderr, arguments...)
}
