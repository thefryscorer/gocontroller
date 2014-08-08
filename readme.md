# GoController

Gocontroller is a library for using a web browser as a controller for games and applications. It uses Go to host a gamepad webpage locally, which can be used by multiple devices simultaneously to play games and/or send keys. Gocontroller can be used for:

- Writing Go-SDL games with mobile controller support.
- As a substitute for physical controllers for local multiplayer games and emulators
- Controlling applications on the RaspberryPi
- With custom gamepads it could be adapted to multiplayer simulation games

[Video of an example program using Gocontroller](http://dbyron.id.au/static/files/videos/goctrl.mp4)

##Installation

    #go get github.com/thefryscorer/gocontroller

## Features

- Multiple players with IP address checking
- Support for multiple instances on different ports
- User defined buttons and styles

## Examples

This repository includes two examples of different uses of Gocontroller. 

The first uses Gocontroller to trigger system calls to xte on a linux machine to send key inputs to applications on the desktop. Requires a linux system and the xte binary to be installed to run.

The second is a very rudimentary multiplayer SDL application using Gocontroller to move circles around a screen. This example relies on the Go-SDL package and SDL must be installed.

### More examples

#### Simple custom gamepad

	package main
	
	import (
		"fmt"

		"github.com/thefryscorer/gocontroller"
	)

	func main() {
		layout := gocontroller.Layout{Style: gocontroller.DefaultCSS, Buttons: []gocontroller.Button{
			{Left: 30, Top: 50, Key: "On"},
			{Left: 70, Top: 50, Key: "Off"},
		}}
		server := gocontroller.NewServer(layout, gocontroller.DefaultPort)
		server.Start()
		fmt.Println("Server started.")
		inAgg := server.NewInputAggregator()
		for {
			inAgg.Collect()
			for _, in := range inAgg.Inputs {
				if in.Event == gocontroller.PRESS {
					switch in.Key {
					case "On":
						fmt.Println("On was pressed")
					case "Off":
						fmt.Println("Off was pressed")
					}
				}
			}
			inAgg.Clear()
		}
	}	

