package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/xor-gate/goexif2/exif"
)


// Program options.
type Config struct {
    Prefix string; // specify a prefix for when EXIF parsing fails
};


// Used to pass information from goroutine back to sender.
type Result struct {
    Arg string;
    Year int;
    Error error;
};


func run(arg string, resultlog chan Result) {
    var result Result;
    result.Arg = arg;
    result.Error = nil;

    file, err := os.Open(arg);
    if err != nil {
        result.Error = fmt.Errorf("Couldn't open file: %s", arg);
        resultlog <- result;
        return;
    }

    year, err := get_year_exif(file);
    if err != nil {
        if ProgramConfig.Prefix == "" {
            result.Error = fmt.Errorf(
                "Couldn't parse EXIF. File: %s", arg);

            resultlog <- result;
            return;
        }

        year, err = get_year_prefix(arg);
        if err != nil {
            result.Error = fmt.Errorf(
                "Couldn't parse EXIF and prefix does not match. File: %s", arg);

            resultlog <- result;
            return;
        }
    }

    result.Year = year;
    resultlog <- result;
}


// Given a file, return the year that it was taken (based on EXIF data).
func get_year_exif(file *os.File) (int, error) {
    img, err := exif.Decode(file);
    if err != nil {
        return 0, err
    }

    timeshot, err := img.DateTime();

    return timeshot.Year(), nil;
}


// Try to parse the name of the file based on a given prefix 
// (specified with the -prefix option). The function will try to parse
// the next four characters as the year.
func get_year_prefix(name string) (int, error) {
    prefix_len := len(ProgramConfig.Prefix);

    if name[0:prefix_len] != ProgramConfig.Prefix {
        return -1, fmt.Errorf("File name does not match 'IMG-' prefix: %s", name);
    }

    year, err := strconv.Atoi(name[prefix_len:prefix_len + 4]);
    if err != nil {
        return -1, err;
    }

    return year, nil;
}
