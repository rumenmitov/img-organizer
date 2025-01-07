package imgorganizer;

import (
    "log"
    "os"
)

const help_message string =
`
img-organizer

Organizes images into directories of the year they were taken.
Image names that failed are saved in img-organizer.log

Usage: img-organizer <file1.png> <file2.jpg> ...
`;


func help() []byte {
    helpfile, err := os.Open("help.txt");
    if err != nil {
        log.Panic("Couldn't open help.txt!");
    }

    helpstat, err := helpfile.Stat();
    if err != nil {
        log.Panic("Couldn't get help.txt stats!");
    }

    var help []byte = make([]byte, helpstat.Size());
    var bytes_read int64 = 0;

    for ;bytes_read < helpstat.Size(); {
        bytes, err := helpfile.Read(help);
        if err != nil {
            log.Panic("Couldn't read help.txt!");
        }

        bytes_read += int64(bytes);
    }

    return help;
}


