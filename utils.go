package imgorganizer;

import (
    "log"
	"os"
	"github.com/xor-gate/goexif2/exif"
)


const LOG_FILE string = "img-organizer.log";


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


func get_year(file *os.File) (int, error) {
    img, err := exif.Decode(file);
    if err != nil {
        return 0, err
    }

    timeshot, err := img.DateTime();

    return timeshot.Year(), nil;
}
