package main

import (
	"github.com/alekseinovikov/ebitengine-hello-world/internal/assets"
	"github.com/alekseinovikov/ebitengine-hello-world/internal/controls"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	input "github.com/quasilyte/ebitengine-input"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gmath"
)

// resource - "github.com/quasilyte/ebitengine-resource"
// audio - "github.com/hajimehoshi/ebiten/v2/audio"
// assets - "github.com/quasilyte/ebitengine-hello-world/internal/assets"
func createLoader() *resource.Loader {
	sampleRate := 44100
	audioContext := audio.NewContext(sampleRate)
	loader := resource.NewLoader(audioContext)
	loader.OpenAssetFunc = assets.OpenAsset
	return loader
}

func main() {
	g := &myGame{
		windowWidth:  320,
		windowHeight: 240,
		loader:       createLoader(),
	}

	assets.RegisterResources(g.loader)
	ebiten.SetWindowSize(g.windowWidth, g.windowHeight)
	ebiten.SetWindowTitle("Ebitengine Quest")

	g.init()
	// RunGame expects a Game interface, which has three methods:
	// Update, Draw and Layout
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

type myGame struct {
	windowWidth  int
	windowHeight int
	loader       *resource.Loader
	input        *input.Handler
	inputSystem  input.System
	player       *Player
}

type Player struct {
	pos gmath.Vec // {X, Y}
	img *ebiten.Image
}

func (g *myGame) init() {
	gopher := g.loader.LoadImage(assets.ImageGopher).Data
	g.player = &Player{img: gopher}

	g.inputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})
	g.input = g.inputSystem.NewHandler(0, controls.DefaultKeymap)
}

// Update By default TPS is 60
func (g *myGame) Update() error {
	g.inputSystem.Update()
	speed := 64.0 * (1.0 / 60)
	var v gmath.Vec
	if g.input.ActionIsPressed(controls.ActionMoveRight) {
		v.X += speed
	}
	if g.input.ActionIsPressed(controls.ActionMoveDown) {
		v.Y += speed
	}
	if g.input.ActionIsPressed(controls.ActionMoveLeft) {
		v.X -= speed
	}
	if g.input.ActionIsPressed(controls.ActionMoveUp) {
		v.Y -= speed
	}
	g.player.pos = g.player.pos.Add(v)
	return nil
}

func (g *myGame) Draw(screen *ebiten.Image) {
	var options ebiten.DrawImageOptions
	options.GeoM.Translate(g.player.pos.X, g.player.pos.Y)
	screen.DrawImage(g.player.img, &options)
}

func (g *myGame) Layout(w, h int) (int, int) {
	return g.windowWidth, g.windowHeight
}
