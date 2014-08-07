# GoController

Gocontroller is a library for using a web browser as a controller for games and applications. It uses Go to host a gamepad webpage locally, which can be used by multiple devices simultaneously to play games and/or send keys. Gocontroller can be used for:

- Writing Go-SDL games with mobile controller support.
- As a substitute for physical controllers for local multiplayer games and emulators
- Controlling applications on the RaspberryPi
- With custom gamepads it could be adapted to multiplayer simulation games

##Installation

    #go get github.com/thefryscorer/gocontroller

## Examples

### Simple custom gamepad

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
    			switch in.Key {
    			case "On":
    				fmt.Println("On was pressed")
    			case "Off":
    				fmt.Println("Off was pressed")
    			}
    		}
    	        //Clear inputs
    	        inAgg.Clear()
    	}       
    }





## Features

- Multiple players with IP address checking
- Support for multiple instances on different ports
- User defined buttons and styles

## To Do

- Differentiate between button presses and releases (Javascript frightens me)
- More examples
