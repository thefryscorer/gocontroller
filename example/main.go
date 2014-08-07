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
	layout := gocontroller.Layout{Style: gocontroller.DefaultCSS, Buttons: []gocontroller.Button{
		{Left: 20, Top: 20, Key: "Up"},
		{Left: 20, Top: 60, Key: "Down"},
		{Left: 10, Top: 40, Key: "Left"},
		{Left: 30, Top: 40, Key: "Right"},
		{Left: 60, Top: 40, Key: "a", Color:"#204387"},
		{Left: 80, Top: 40, Key: "b", Color:"#208743"},
		{Left: 45, Top: 10, Key: "Return"},
	}}
	server := gocontroller.NewServer(layout, gocontroller.DefaultPort)
	server.Start()
	fmt.Println("Server started.")
	inAgg := server.NewInputAggregator()
	for {
		inAgg.Collect()
		for _, in := range inAgg.Inputs {
			keypress(in.Key)
		}

		//Clear inputs
		inAgg.Clear()

	}
}
