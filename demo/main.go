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
	gocontroller.StartServer()
	fmt.Println("Server started.")
	sdl.Init(sdl.INIT_EVERYTHING)
	screen = sdl.SetVideoMode(sWidth, sHeight, sBpp, flags)
	sdl.WM_SetCaption("Controller Demo", "none")
	defer sdl.Quit()

	ballImg := sdl.DisplayFormatAlpha(sdl.Load("ball.png"))

	var players = make([]*Player, 0)

	inAgg := gocontroller.NewInputAggregator()

	players = append(players, waitForPlayer(inAgg))

	var speed int = 3

	for {
		inAgg.Collect()
		for _, in := range inAgg.Inputs {
			switch in.Button {
			case gocontroller.UP:
				for i := 0; i < len(players); i++ {
					if in.UserIP == players[i].IP {
						players[i].ball.y -= speed
						break
					}
				}
			case gocontroller.DOWN:
				for i := 0; i < len(players); i++ {
					if in.UserIP == players[i].IP {
						players[i].ball.y += speed
						break
					}
				}
			case gocontroller.LEFT:
				for i := 0; i < len(players); i++ {
					if in.UserIP == players[i].IP {
						players[i].ball.x -= speed
						break
					}
				}
			case gocontroller.RIGHT:
				for i := 0; i < len(players); i++ {
					if in.UserIP == players[i].IP {
						players[i].ball.x += speed
						break
					}
				}
			case gocontroller.START:
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
