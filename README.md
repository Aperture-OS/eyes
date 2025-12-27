# Eyes

**Eyes** is a lightweight, easy-to-use logging library for Go with colored output powered by [`fatih/color`](https://github.com/fatih/color). It provides easy-to-use functions like Info, Warn, Success, and Fatal with configurable prefixes and timestamps, ideal for CLI apps and system tools

## Features

- Easy-to-use logging functions: `Info`, `Warn`, `Success`, `Fatal`
- Configurable prefix templates with placeholders:
  - `{display_name}` – your logger name
  - `{timestamp}` – current time
  - `{log_level}` – log level (INFO, WARN, SUCCESS, FATAL)
- Configurable timestamp format
- Per-level color schemes

## Installation

```bash
go get github.com/Aperture-OS/eyes
```

## Usage

```go
package main

import (
    "github.com/Aperture-OS/eyes"
    "github.com/fatih/color"
    "time"
)

func main() {
    // Configure the logger
    eyes.SetLoggerConfiguration(eyes.LoggerConfiguration{
        DisplayName:     "DEBUG",
        PrefixTemplate:  "[{display_name}] {timestamp} {log_level} ",
        TimestampFormat: "15:04:05",
        InfoTextColor:    color.New(color.FgCyan),
        WarnTextColor:    color.New(color.FgYellow),
        SuccessTextColor: color.New(color.FgGreen),
        FatalTextColor:   color.New(color.FgRed, color.Bold),
    })

    eyes.Info("This is an info message")
    eyes.Warnf("This is a warning about %s", "disk space")
    eyes.Successln("Operation completed successfully")
    // eyes.Fatal("Fatal error occurred") // will exit the program
}
```

## License

This project is licensed under the **GNU General Public License v3.0**. See [LICENSE](LICENSE) for details.
Copyright © ApertureOS 2025-2026
