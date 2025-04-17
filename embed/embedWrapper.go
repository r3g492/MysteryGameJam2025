package embedWrapper

import (
	"embed"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
	"os"
)

//go:embed resources/music/*
var musicFS embed.FS

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
