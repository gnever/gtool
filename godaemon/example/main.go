package main

import (
	"time"
	"fmt"

	_ "github.com/gnever/gtool/godaemon"
    "github.com/gnever/gtool/gfile"
)

func main() {

    file := "/tmp/godaemon-example-log/show.log"
    gfile.PutContentsAppend(file, fmt.Sprintf("start  %d\n", time.Now().Unix()))
	for i := 0; i < 100; i++ {
        gfile.PutContentsAppend(file, fmt.Sprintf("do  %d\n", time.Now().Unix()))
		time.Sleep(time.Second)
	}
    gfile.PutContentsAppend(file, fmt.Sprintf("end  %d\n", time.Now().Unix()))

}
