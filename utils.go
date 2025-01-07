package main;

import (
	"os"
    "fmt"
	"path"
	"strconv"
	"github.com/xor-gate/goexif2/exif"
)


// Used to pass information from goroutine back to sender.
type Result struct {
    Arg string;
    Error error;
};

func run(arg string, resultlog chan Result) {
    var result Result;
    result.Arg = arg;
    result.Error = nil;

    file, err := os.Open(arg);
    if err != nil {
        result.Error = fmt.Errorf("Couldn't open file: %s\n", arg);
        resultlog <- result;
        return;
    }

    year, err := get_year(file);
    if err != nil {
        result.Error = fmt.Errorf("Couldn't parse EXIF of file: %s\n", arg);
        resultlog <- result;
        return;
    }

    yearstr := strconv.Itoa(year);

    err = os.MkdirAll(yearstr, 777); 
    if err != nil {
        result.Error = fmt.Errorf("Couldn't open directory %s for %s\n", yearstr, arg);
        resultlog <- result;
        return;
    }


    err = os.Rename(arg, path.Join(".", yearstr, arg));
    if err != nil {
        result.Error = fmt.Errorf("Couldn't move %s\n", arg);
        resultlog <- result;
        return;
    }

    resultlog <- result;
}


// Given a file, return the year that it was taken (based on EXIF data).
func get_year(file *os.File) (int, error) {
    img, err := exif.Decode(file);
    if err != nil {
        return 0, err
    }

    timeshot, err := img.DateTime();

    return timeshot.Year(), nil;
}
