/*
Img-organizer organizes images into directories of the year they were taken.

The year is taken from the file's EXIF data. If this is not possible, an error
will be printed to stderr.

Usage:
    img-organizer <file> ...

Users can pass in one or more files at a time. Processing happens concurrently.
If no files are provided, a default help message is printed.

Options:
    --prefix=<some-prefix>          If EXIF fails, try to discern the four
                                    four digits after the prefix as the year.
*/
package main

import (
	"flag"
	"fmt"
	"os"
    "strconv"
    "path"
)

// NOTE: This can be a global variable. It is set in the beginning of the program
//       and from then on it is read-only. Thus, no need for mutexes either.
var ProgramConfig Config;


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
            continue;
        }

        yearstr := strconv.Itoa(result.Year);

        err := os.MkdirAll(yearstr, 0777); 
        if err != nil {
            fmt.Fprintf(
                os.Stderr, 
                "Couldn't open directory %s for %s\n", 
                yearstr, 
                result.Arg);

            continue;
        }


        err = os.Rename(result.Arg, path.Join(".", yearstr, result.Arg));
        if err != nil {
            fmt.Fprintf(os.Stderr, "Couldn't move %s\n", result.Arg);
            continue;
        }
    }
}
