package main

import (
	"github.com/alekseinovikov/ebitengine-hello-world/internal/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	resource "github.com/quasilyte/ebitengine-resource"
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
}

func (g *myGame) Update() error {
	return nil
}

func (g *myGame) Draw(screen *ebiten.Image) {
	gopher := g.loader.LoadImage(assets.ImageGopher).Data
	var options ebiten.DrawImageOptions
	screen.DrawImage(gopher, &options)
}

func (g *myGame) Layout(w, h int) (int, int) {
	return g.windowWidth, g.windowHeight
}
