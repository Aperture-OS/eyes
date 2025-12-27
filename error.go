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
