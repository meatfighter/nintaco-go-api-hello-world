package main

import (
	"fmt"

	"github.com/meatfighter/nintaco-go-api/nintaco"
)

const (
	str = "Hello, World!"

	spriteID   = 123
	spriteSize = 32
)

type helloWorld struct {
	api nintaco.API

	spriteX  int
	spriteY  int
	spriteVx int
	spriteVy int

	strWidth int
	strX     int
	strY     int
}

func newHelloWorld() *helloWorld {
	return &helloWorld{
		api:      nintaco.GetAPI(),
		spriteX:  0,
		spriteY:  8,
		spriteVx: 1,
		spriteVy: 1,
	}
}

func (h *helloWorld) launch() {
	h.api.AddFrameListener(h)
	h.api.AddStatusListener(h)
	h.api.AddActivateListener(h)
	h.api.AddDeactivateListener(h)
	h.api.AddStopListener(h)
	h.api.Run()
}

func (h *helloWorld) APIEnabled() {
	fmt.Println("API enabled")

	sprite := make([]int, spriteSize*spriteSize)
	for y := 0; y < spriteSize; y++ {
		Y := float64(y) - spriteSize/2 + 0.5
		for x := 0; x < spriteSize; x++ {
			X := float64(x) - spriteSize/2 + 0.5
			var color int
			if X*X+Y*Y <= spriteSize*spriteSize/4 {
				color = nintaco.ColorOrange
			} else {
				color = -1
			}
			sprite[spriteSize*y+x] = color
		}
	}
	h.api.CreateSprite(spriteID, spriteSize, spriteSize, sprite)

	h.strWidth = h.api.GetStringWidth(str, false)
	h.strX = (256 - h.strWidth) / 2
	h.strY = (240 - 8) / 2
}

func (h *helloWorld) APIDisabled() {
	fmt.Println("API disabled")
}

func (h *helloWorld) Dispose() {
	fmt.Println("API stopped")
}

func (h *helloWorld) StatusChanged(message string) {
	fmt.Printf("Status message: %s\n", message)
}

func (h *helloWorld) FrameRendered() {
	h.api.DrawSprite(spriteID, h.spriteX, h.spriteY)
	if h.spriteX+spriteSize == 255 {
		h.spriteVx = -1
	} else if h.spriteX == 0 {
		h.spriteVx = 1
	}
	if h.spriteY+spriteSize == 231 {
		h.spriteVy = -1
	} else if h.spriteY == 8 {
		h.spriteVy = 1
	}
	h.spriteX += h.spriteVx
	h.spriteY += h.spriteVy

	h.api.SetColor(nintaco.ColorDarkBlue)
	h.api.FillRect(h.strX-1, h.strY-1, h.strWidth+2, 9)
	h.api.SetColor(nintaco.ColorBlue)
	h.api.DrawRect(h.strX-2, h.strY-2, h.strWidth+3, 10)
	h.api.SetColor(nintaco.ColorWhite)
	h.api.DrawString(str, h.strX, h.strY, false)
}

func main() {
	nintaco.InitRemoteAPI("localhost", 9999)
	newHelloWorld().launch()
}
