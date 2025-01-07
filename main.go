package imgorganizer

import (
	"fmt"
	"log"
	"os"
    "path"
    "strconv"
)

func run(arg string, errorlog chan error) {
    file, err := os.Open(arg);
    if err != nil {
        errorlog <- fmt.Errorf("Couldn't open file: %s\n", arg);
        return;
    }

    year, err := get_year(file);
    if err != nil {
        errorlog <- fmt.Errorf("Couldn't parse EXIF of file: %s\n", arg);
        return;
    }

    yearstr := strconv.Itoa(year);

    err = os.MkdirAll(yearstr, 777); 
    if err != nil {
        errorlog <- fmt.Errorf("Couldn't open directory %s for %s\n", yearstr, arg);
        return;
    }


    err = os.Rename(arg, path.Join(".", yearstr, arg));
    if err != nil {
        errorlog <- fmt.Errorf("Couldn't move %s\n", arg);
        return;
    }

    errorlog <- nil;
}


func main() {
    logfile, err := os.OpenFile(LOG_FILE, os.O_APPEND | os.O_CREATE, 700);
    if err != nil {
        log.Panic("Couldn't open log file!\n");
    }

    errorlogs := make(chan error, len(os.Args));

    for i := 1; i < len(os.Args); i++ {
        arg := os.Args[i];

        if arg == "help" || arg == "-h" || arg == "--help" {
            fmt.Println(help());
            os.Exit(0);
        }

        go run(arg, errorlogs);
    }

    for i := 0; i < len(os.Args); i++ {
        err := <- errorlogs;
        if err != nil {
            fmt.Fprintf(os.Stderr, err.Error());
            logfile.WriteString(arg)

        }
    }
}
