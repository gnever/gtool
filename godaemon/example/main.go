package main

import (
	"time"

	_ "github.com/gnever/gtool/godaemon"
	"github.com/gogf/gf/os/glog"
)

func main() {

	l := glog.New().Line(true)
	l.SetPath("/tmp/godaemon-example-log")
	l.SetStdoutPrint(true)

	l.Debugf("start  %d", time.Now().Unix())
	for i := 0; i < 100; i++ {
		l.Debugf("do  %d", time.Now().Unix())
		time.Sleep(time.Second)
	}
	l.Debugf("end  %d", time.Now().Unix())

}
