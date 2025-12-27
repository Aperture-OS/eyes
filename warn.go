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
