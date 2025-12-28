/*
  Eyes is a lightweight Go logger with colored output using fatih/color.
  Copyright (C) 2025-2026 ApertureOS Team
  Licensed under the MIT license. See LICENSE file in the project root for details.
*/

package eyes

import (
	"github.com/fatih/color"
	"strings"
	"time"
)

// LoggerConfiguration
// Holds all user-configurable settings for the logger
type LoggerConfiguration struct {
	// DisplayName is substituted into {display_name}
	// Example: "DEBUG", "Blink", "Aperture"
	DisplayName string

	// PrefixTemplate defines the exact text printed
	// before every log message.
	//
	// Supported placeholders:
	//   {display_name} -> DisplayName
	//   {timestamp}    -> current time
	//   {log_level}    -> INFO, WARN, SUCCESS, FATAL
	PrefixTemplate string

	// TimestampFormat controls how {timestamp} is rendered
	// Uses time.Format layout strings
	TimestampFormat string

	// Colors used for each log level
	InfoTextColor    *color.Color
	WarnTextColor    *color.Color
	SuccessTextColor *color.Color
	FatalTextColor   *color.Color
}

// DefaultLoggerConfiguration
// Used unless overridden via SetLoggerConfiguration
var loggerConfig = LoggerConfiguration{
	DisplayName:     "DEBUG",
	PrefixTemplate:  "[{display_name}] {timestamp} {log_level}: ",
	TimestampFormat: "15:04:05",

	InfoTextColor:    color.New(color.FgBlue),
	WarnTextColor:    color.New(color.FgYellow),
	SuccessTextColor: color.New(color.FgGreen),
	FatalTextColor:   color.New(color.FgRed),
}

// SetLoggerConfiguration
// Replaces the global logger configuration
// example:
/*
	eyes.SetLoggerConfiguration(eyes.LoggerConfiguration{
		DisplayName:     "DEBUG",  // goes in [DISPLAYNAME] 15:05:04 FATAL: Error message
		PrefixTemplate:  "[{display_name}] {timestamp} {log_level} ", // note the space at the end
		TimestampFormat: "15:04:05", // which is: time.Now().Format("TimestampFormat HERE")

		InfoTextColor:    color.New(color.FgCyan),   // use any color from github.com/fatih/color
		WarnTextColor:    color.New(color.FgYellow), // usually you should use ansi-16 colors
		SuccessTextColor: color.New(color.FgGreen),
		FatalTextColor:   color.New(color.FgRed, color.Bold),
	})
*/
func SetLoggerConfiguration(config LoggerConfiguration) {
	loggerConfig = config
}

// buildPrefixString
// Expands PrefixTemplate into a concrete prefix string
func buildPrefixString(logLevel string) string {
	prefix := loggerConfig.PrefixTemplate

	prefix = strings.ReplaceAll(
		prefix,
		"{display_name}",
		loggerConfig.DisplayName,
	)

	prefix = strings.ReplaceAll(
		prefix,
		"{timestamp}",
		time.Now().Format(loggerConfig.TimestampFormat),
	)

	prefix = strings.ReplaceAll(
		prefix,
		"{log_level}",
		logLevel,
	)

	return prefix
}
