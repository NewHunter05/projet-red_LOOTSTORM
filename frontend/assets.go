package main

import (
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func loadSprites() map[string]*ebiten.Image {
	sprites := make(map[string]*ebiten.Image)
	for key, filename := range spriteFiles {
		for _, root := range assetRoots("sprites") {
			path := filepath.Join(root, filename)
			if _, err := os.Stat(path); err != nil {
				continue
			}
			image, _, err := ebitenutil.NewImageFromFile(path)
			if err == nil {
				sprites[key] = image
			}
			break
		}
	}
	return sprites
}

func loadBackground() *ebiten.Image {
	return loadBackgroundFile("dungeon.png")
}

func loadBackgroundFile(filename string) *ebiten.Image {
	for _, root := range assetRoots("backgrounds") {
		path := filepath.Join(root, filename)
		if _, err := os.Stat(path); err != nil {
			continue
		}
		background, _, err := ebitenutil.NewImageFromFile(path)
		if err == nil {
			return background
		}
	}
	return nil
}

func assetRoots(kind string) []string {
	return []string{filepath.Join("frontend", "assets", kind), filepath.Join("assets", kind)}
}
