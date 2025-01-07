package main

import (
	"flag"
	"fmt"
    "os"
)

const help_message string =
`
img-organizer

Organizes images into directories of the year they were taken.

Usage:
    img-organizer <file> ...

Users can pass in one or more files at a time. Processing happens concurrently.
If no files are provided, a default help message is printed.

Options:
`;


// Initializes program config with the flags,
// overrides flag.Usage to print help message, then exits program.
func setup_flags() {
    flag.StringVar(&ProgramConfig.Prefix, "prefix", "", "Try parsing with prefix.");

    flag.Usage = func() {
        fmt.Print(help_message);
        flag.PrintDefaults();
        os.Exit(0);
    }
}
