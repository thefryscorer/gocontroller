package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"runtime/pprof"

	"github.com/thefryscorer/gocontroller"
)

func keypress(key string) {
	fmt.Println(key)
	go func() {
		cmd := exec.Command("xdotool", "key", key)
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Warning: %v\n", err)
		}
	}()
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
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
				pprof.StopCPUProfile()
				os.Exit(0)

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
