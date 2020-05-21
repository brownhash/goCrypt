package logger

import (
	"fmt"
	"os"
)

const (
	ErrorColor  = "\033[1;31m%s\033[0m" // ansi color
	WarningColor = "\033[1;33m%s\033[0m" // ansi color
	InfoColor    = "\033[1;34m%s\033[0m" // ansi color
	HelpText = `Usage: gocrypt [option] [supporting options]

Options:
1) -w (write a file): gocrypt -w filename [optional arguments]
2) -r (read a encrypted file): gocrypt -r filename [optional arguments]
3) -h (help): gocrypt -h

Optional Arguments:
1) -k (encryption key): 
          gocrypt -w filename -k encryption_key
          gocrypt -r filename -k encryption_key
   if key is not provided then the default key is used
`
)

// logs a critical message and exits after printing the error
func Critical(message string) {
	fmt.Printf(ErrorColor, message)
	os.Exit(1)
}

// logs a warning message and allows the further execution
func Warning(message string) {
	fmt.Printf(WarningColor, message)
}

// logs any informative data
func Info(message string) {
	fmt.Printf(InfoColor, message)
}

// logs help context as a warning
func Help() {
	Warning(HelpText)
}
