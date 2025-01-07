package imgorganizer;

import (
	"os"
	"github.com/xor-gate/goexif2/exif"
)


const LOG_FILE string = "img-organizer.log";


func get_year(file *os.File) (int, error) {
    img, err := exif.Decode(file);
    if err != nil {
        return 0, err
    }

    timeshot, err := img.DateTime();

    return timeshot.Year(), nil;
}
