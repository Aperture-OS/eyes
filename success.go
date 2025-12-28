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

// Success prints a success message.
func Success(arguments ...any) {
	loggerConfig.SuccessTextColor.Fprint(
		os.Stderr,
		buildPrefixString("SUCCESS"),
	)
	fmt.Fprint(os.Stderr, arguments...)
	fmt.Fprintln(os.Stderr)
}

// Successf prints a formatted success message.
func Successf(format string, arguments ...any) {
	loggerConfig.SuccessTextColor.Fprint(
		os.Stderr,
		buildPrefixString("SUCCESS"),
	)
	fmt.Fprintf(os.Stderr, format, arguments...)
	fmt.Fprintln(os.Stderr)
}

// Successln prints a success message followed by a newline.
func Successln(arguments ...any) {
	loggerConfig.SuccessTextColor.Fprint(
		os.Stderr,
		buildPrefixString("SUCCESS"),
	)
	fmt.Fprintln(os.Stderr, arguments...)
}
