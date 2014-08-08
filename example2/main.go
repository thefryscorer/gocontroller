package main

import (
	"image"
	"image/color"
	"math/rand"
	"os"
	"time"

	"github.com/banthar/Go-SDL/sdl"
	"github.com/thefryscorer/gocontroller"
)

const (
	screenWidth  = 800
	screenHeight = 500
	screenBPP    = 32
	screenFlags  = sdl.SWSURFACE

	playerSpeed = 5
	gameFPS     = 30
)

var (
	playerColors = []color.NRGBA{
		{255, 255, 255, 255},
		{128, 255, 255, 255},
		{255, 128, 255, 255},
		{255, 255, 128, 255},
	}
)

// Helper functions.

// sdlBlit draws the first argument on top of the second argument, with the first argument's
// top-left corner located at (x, y).
func sdlBlit(surf *sdl.Surface, screen *sdl.Surface, x int, y int) {
	rect := sdl.Rect{
		X: int16(x),
		Y: int16(y),
		W: uint16(surf.W),
		H: uint16(surf.H),
	}
	screen.Blit(&rect, surf, nil)
}

// makeBall returns an SDL surface containing an image of a ball with the given radius and color.
func makeBall(radius int, col color.Color) *sdl.Surface {
	img := image.NewNRGBA(image.Rect(0, 0, radius*2, radius*2))
	for x := -radius; x < radius; x++ {
		for y := -radius; y < radius; y++ {
			if x*x+y*y < radius*radius {
				img.Set(radius+x, radius+y, col)
			}
		}
	}
	return sdl.CreateSurfaceFromImage(img)
}

// Player Management.

type Player struct {
	x, y   int
	dX, dY int
	ip     string
	surf   *sdl.Surface
}

func (p *Player) processInput(in gocontroller.Input) {
	switch in.Key {
	case "Up":
		if in.Event == gocontroller.PRESS {
			p.dY = -playerSpeed
		} else if in.Event == gocontroller.RELEASE {
			p.dY = 0
		}
	case "Down":
		if in.Event == gocontroller.PRESS {
			p.dY = playerSpeed
		} else if in.Event == gocontroller.RELEASE {
			p.dY = 0
		}
	case "Left":
		if in.Event == gocontroller.PRESS {
			p.dX = -playerSpeed
		} else if in.Event == gocontroller.RELEASE {
			p.dX = 0
		}
	case "Right":
		if in.Event == gocontroller.PRESS {
			p.dX = playerSpeed
		} else if in.Event == gocontroller.RELEASE {
			p.dX = 0
		}
	case "A":
		if in.Event == gocontroller.PRESS {
			p.surf = makeBall(16, color.NRGBA{R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)), A: 255})
		}
	}
}

func (p *Player) update() {
	p.x += p.dX
	p.y += p.dY
}

func newPlayer(ip string) Player {
	col := playerColors[rand.Intn(len(playerColors))]
	return Player{
		x:    screenWidth/2 - 8,
		y:    screenHeight/2 - 8,
		ip:   ip,
		surf: makeBall(16, col),
	}
}

func main() {
	//runtime.GOMAXPROCS(4)

	layout := gocontroller.Layout{Style: gocontroller.DefaultCSS, Buttons: []gocontroller.Button{
		{Left: 20, Top: 20, Key: "Up"},
		{Left: 20, Top: 60, Key: "Down"},
		{Left: 10, Top: 40, Key: "Left"},
		{Left: 30, Top: 40, Key: "Right"},
		{Left: 45, Top: 10, Key: "Start"},
		{Left: 75, Top: 40, Key: "A", Color: "#872828"},
	}}
	server := gocontroller.NewServer(layout, gocontroller.DefaultPort)
	server.Start()

	sdl.Init(sdl.INIT_VIDEO)
	screenSurf := sdl.SetVideoMode(screenWidth, screenHeight, screenBPP, screenFlags)
	sdl.WM_SetCaption("Controller Demo", "none")
	defer sdl.Quit()

	players := make([]Player, 0)
	inAgg := server.NewInputAggregator()

	frameTime := 1.0 / gameFPS * float64(time.Second)
	ticker := time.NewTicker(time.Duration(frameTime))

	for {
		inAgg.Collect()
		for _, in := range inAgg.Inputs {
			// Find the player associated with the input (if any).
			found := false
			for i := 0; i < len(players); i++ {
				if players[i].ip == in.UserIP {
					players[i].processInput(in)
					found = true
					break
				}
			}

			// If no player was found, create a new one.
			if !found && in.Key == "Start" {
				players = append(players, newPlayer(in.UserIP))
			}
		}

		// Update the player positions.
		for i := 0; i < len(players); i++ {
			players[i].update()
		}

		// Clear the screen.
		screenSurf.FillRect(&screenSurf.Clip_rect, 0)

		// Redraw players.
		for _, p := range players {
			sdlBlit(p.surf, screenSurf, p.x, p.y)
		}

		// Update screen.
		screenSurf.Flip()

		//Clear inputs
		inAgg.Clear()

		//Check for SDL quit
		for ev := sdl.PollEvent(); ev != nil; ev = sdl.PollEvent() {
			if _, ok := ev.(*sdl.QuitEvent); ok {
				os.Exit(0)
			}
		}

		// Keep the framerate reasonable.
		<-ticker.C
	}
}
