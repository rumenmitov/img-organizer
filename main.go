/*
Img-organizer organizes images into directories of the year they were taken.

The year is taken from the file's EXIF data. If this is not possible, an error
will be printed to stderr.

Usage:
    img-organizer <file> ...

Users can pass in one or more files at a time. Processing happens concurrently.
If no files are provided, a default help message is printed.
*/
package main

import (
	"flag"
	"fmt"
	"os"
)


func main() {
    setup_flags();
    flag.Parse();

    if (len(os.Args) == 1) {
        flag.Usage();
    }

    resultlogs := make(chan Result, len(os.Args));

    for i := 0; i < len(flag.Args()); i++ {
        arg := flag.Args()[i];

        go run(arg, resultlogs);
    }

    for i := 0; i < len(flag.Args()); i++ {
        result := <- resultlogs;
        if result.Error != nil {
            fmt.Fprintf(os.Stderr, "%s\n", result.Error.Error());
        }
    }
}
