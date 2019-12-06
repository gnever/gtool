package godaemon

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func init() {
	var godaemon = flag.Bool("d", false, "run as a daemon with -d=true")
	if !flag.Parsed() {
		flag.Parse()
	}

	if *godaemon {
		args := os.Args[1:]

		for i := 0; i < len(args); i++ {
			if args[i] == "-d=true" ||  args[i] == "-d" {
				args[i] = "-d=false"
				break
			}
		}
		cmd := exec.Command(os.Args[0], args...)

		cmd.Stdin = nil
		cmd.Stdout = nil
		cmd.Stderr = nil
		cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}

		err := cmd.Start()
		if err == nil {
			fmt.Println("[PID]", cmd.Process.Pid)
			cmd.Process.Release()
		} else {
			fmt.Println(err)
		}

		os.Exit(0)
	}
}
