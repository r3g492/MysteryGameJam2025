package embedWrapper

import (
	"embed"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
	"os"
	"path/filepath"
)

//go:embed resources/music/*
var musicFS embed.FS

//go:embed resources/texture/*
var textureFS embed.FS

func LoadSoundFromEmbedded(filename string) rl.Sound {
	data, err := musicFS.ReadFile("resources/music/" + filename)
	if err != nil {
		log.Fatalf("failed to read embedded sound %s: %v", filename, err)
	}
	tmpFile, err := os.CreateTemp("", "*.mp3")
	if err != nil {
		log.Fatalf("failed to create temporary file for %s: %v", filename, err)
	}
	_, err = tmpFile.Write(data)
	if err != nil {
		log.Fatalf("failed to write to temporary file for %s: %v", filename, err)
	}
	tmpFile.Close()
	if err != nil {
		log.Fatalf("failed to rename temporary file: %v", err)
	}
	snd := rl.LoadSound(tmpFile.Name())
	return snd
}

func LoadTextureFromEmbedded(filename string, resizeWidth int32, resizeHeight int32) (rl.Texture2D, *rl.Image) {
	data, err := textureFS.ReadFile("resources/texture/" + filename)
	if err != nil {
		log.Fatalf("failed to read embedded file %s: %v", filename, err)
	}
	ext := filepath.Ext(filename)
	img := rl.LoadImageFromMemory(ext, data, int32(len(data)))
	if img.Width == 0 {
		log.Fatalf("failed to load image %s from embedded data", filename)
	}
	if resizeWidth != -1 && resizeHeight != -1 {
		rl.ImageResize(img, resizeWidth, resizeHeight)
	}
	tex := rl.LoadTextureFromImage(img)
	rl.UnloadImage(img)
	return tex, img
}
