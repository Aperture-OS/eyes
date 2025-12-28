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

// Fatal prints a fatal error message and exits with status code 1.
func Fatal(arguments ...any) {
	loggerConfig.FatalTextColor.Fprint(
		os.Stderr,
		buildPrefixString("FATAL"),
	)
	fmt.Fprint(os.Stderr, arguments...)
	fmt.Fprintln(os.Stderr)
	os.Exit(1)
}

// Fatalf prints a formatted fatal error message and exits with status code 1.
func Fatalf(format string, arguments ...any) {
	loggerConfig.FatalTextColor.Fprint(
		os.Stderr,
		buildPrefixString("FATAL"),
	)
	fmt.Fprintf(os.Stderr, format, arguments...)
	fmt.Fprintln(os.Stderr)
	os.Exit(1)
}

// Fatalln prints a fatal error message followed by a newline
// and exits with status code 1.
func Fatalln(arguments ...any) {
	loggerConfig.FatalTextColor.Fprint(
		os.Stderr,
		buildPrefixString("FATAL"),
	)
	fmt.Fprintln(os.Stderr, arguments...)
	os.Exit(1)
}
