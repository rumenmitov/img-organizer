package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strconv"
)

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


func main() {
    setup_flags();
    flag.Parse();

    if (len(os.Args) == 1) {
        flag.Usage();
    }

    resultlogs := make(chan Result, len(os.Args));

    for i := 1; i < len(flag.Args()); i++ {
        arg := os.Args[i];

        go run(arg, resultlogs);
    }

    for i := 1; i < len(flag.Args()); i++ {
        result := <- resultlogs;
        if result.Error != nil {
            fmt.Fprintf(os.Stderr, "%s\n", result.Error.Error());
        }
    }
}
