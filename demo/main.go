package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/banthar/Go-SDL/sdl"
	"github.com/thefryscorer/gocontroller"
)

var (
	sWidth  int = 800
	sHeight int = 500
	sBpp    int = 32

	screen *sdl.Surface
	flags  uint32 = sdl.SWSURFACE
)

func sdlBlit(surf *sdl.Surface, screen *sdl.Surface, x int, y int) {
	rect := sdl.Rect{
		X: int16(x),
		Y: int16(y),
		W: uint16(surf.W),
		H: uint16(surf.H),
	}
	screen.Blit(&rect, surf, nil)
}

type Player struct {
	ball *Ball
	IP   string
}

type Ball struct {
	x int
	y int
}

func newPlayer(ip string) *Player {
	return &Player{
		ball: &Ball{
			x: 0,
			y: 0,
		},
		IP: ip,
	}
}

func waitForPlayer(inAgg gocontroller.InputAggregator) *Player {
	for {
		inAgg.Collect()
		for _, in := range inAgg.Inputs {
			return newPlayer(in.UserIP)
		}
		inAgg.Clear()
	}
}

func main() {
	runtime.GOMAXPROCS(4)
	layout := gocontroller.Layout{Style: gocontroller.DefaultCSS, Buttons: []gocontroller.Button{
		{Left: 20, Top: 20, Key: "Up"},
		{Left: 20, Top: 60, Key: "Down"},
		{Left: 10, Top: 40, Key: "Left"},
		{Left: 30, Top: 40, Key: "Right"},
		{Left: 45, Top: 10, Key: "Start"},
	}}
	server := gocontroller.NewServer(layout, gocontroller.DefaultPort)
	server.Start()
	fmt.Println("Server started.")
	inAgg := server.NewInputAggregator()

	sdl.Init(sdl.INIT_EVERYTHING)
	screen = sdl.SetVideoMode(sWidth, sHeight, sBpp, flags)
	sdl.WM_SetCaption("Controller Demo", "none")
	defer sdl.Quit()

	ballImg := sdl.DisplayFormatAlpha(sdl.Load("ball.png"))

	var players = make([]*Player, 0)

	players = append(players, waitForPlayer(inAgg))

	var speed int = 3

	for {
		inAgg.Collect()
		for _, in := range inAgg.Inputs {
			switch in.Key {
			case "Up":
				for i := 0; i < len(players); i++ {
					if in.UserIP == players[i].IP {
						players[i].ball.y -= speed
						break
					}
				}
			case "Down":
				for i := 0; i < len(players); i++ {
					if in.UserIP == players[i].IP {
						players[i].ball.y += speed
						break
					}
				}
			case "Left":
				for i := 0; i < len(players); i++ {
					if in.UserIP == players[i].IP {
						players[i].ball.x -= speed
						break
					}
				}
			case "Right":
				for i := 0; i < len(players); i++ {
					if in.UserIP == players[i].IP {
						players[i].ball.x += speed
						break
					}
				}
			case "Start":
				for _, p := range players {
					if p.IP != in.UserIP {
						players = append(players, newPlayer(in.UserIP))
					}
				}
			}
		}

		//drawing
		screen.FillRect(&screen.Clip_rect, 0)
		for _, p := range players {
			sdlBlit(ballImg, screen, p.ball.x, p.ball.y)
		}
		screen.Flip()

		//Clear inputs
		inAgg.Clear()

		//Check for sdl quit
		for ev := sdl.PollEvent(); ev != nil; ev = sdl.PollEvent() {
			if _, ok := ev.(*sdl.QuitEvent); ok {
				os.Exit(0)
			}
		}
	}
}
