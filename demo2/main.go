package main

import (
	"fmt"
	"io"

	"os/exec"
	"runtime"

	"github.com/thefryscorer/gocontroller"
)

var cmdPipe io.WriteCloser

func keypress(key string) {
	// We need to sleep a little here because some applications (anything that uses
	// SDL for example) ignore any keypresses with a duration too low.
	cmdPipe.Write([]byte(fmt.Sprintf("keydown %v\n", key)))
	cmdPipe.Write([]byte("usleep 50000\n"))
	cmdPipe.Write([]byte(fmt.Sprintf("keyup %v\n", key)))
}

func main() {
	xte := exec.Command("xte")
	var err error
	cmdPipe, err = xte.StdinPipe()
	if err != nil {
		panic(err)
	}
	defer cmdPipe.Close()

	if err := xte.Start(); err != nil {
		panic(err)
	}

	runtime.GOMAXPROCS(8)
	server := gocontroller.NewServer(gocontroller.DEFAULTPAGE, gocontroller.DEFAULTPORT)
	server.Start()
	fmt.Println("Server started.")
	inAgg := server.NewInputAggregator()
	for {
		inAgg.Collect()
		for _, in := range inAgg.Inputs {
			switch in.Button {
			case gocontroller.UP:
				keypress("Up")

			case gocontroller.DOWN:
				keypress("Down")

			case gocontroller.LEFT:
				keypress("Left")

			case gocontroller.RIGHT:
				keypress("Right")

			case gocontroller.START:
				keypress("Return")
			case gocontroller.A:
				keypress("a")

			case gocontroller.B:
				keypress("b")
			}
		}

		//Clear inputs
		inAgg.Clear()

	}
}
