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

Usage: img-organizer <file> ...

`;

func setup_flags() {
    flag.Usage = func() {
        fmt.Print(help_message);
        flag.PrintDefaults();
        os.Exit(0);
    }
}
