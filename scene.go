package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type scene struct {
	bg   *sdl.Texture
	bird []*sdl.Texture
}

func newScene(r *sdl.Renderer) (*scene, error) {
	bg, err := img.LoadTexture(r, "res/imgs/full-bg.png")
	if err != nil {
		return nil, fmt.Errorf("could not load background image: %v", err)
	}

	for i := 1; i <= 4; i++ {
		bird, err := img.LoadTexture(r, "res/imgs/Frame-1.png")
		if err != nil {
			return nil, fmt.Errorf("could not load bird image: %v", err)
		}
	}

	return &scene{bg: bg, bird: bird}, nil
}

func (s *scene) paint(r *sdl.Renderer) error {
	r.Clear()

	if err := r.Copy(s.bg, nil, nil); err != nil {
		return fmt.Errorf("could not copy texture to renderer: %v", err)
	}

	rect := &sdl.Rect{X: 10, Y: 300 - 43/2, H: 50, W: 43}
	if err := r.Copy(s.bird, nil, rect); err != nil {
		return fmt.Errorf("could not copy texture to renderer: %v", err)
	}

	r.Present()
	return nil
}

func (s *scene) destroy() {
	s.bg.Destroy()
}
