/*
  Eyes is a lightweight Go logger with colored output using fatih/color.
  Copyright (C) 2025-2026 ApertureOS Team

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
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
