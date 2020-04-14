package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	git "github.com/go-git/go-git/v5"
)

func clone(path string) error {
	color.Blue("Cloning maps repository to %v...", path)
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      "https://github.com/infinity-tts-community/maps.git",
		Progress: os.Stdout,
	})
	return err
}

func getpath() (string, error) {
	var path string
	switch os := runtime.GOOS; os {
	case "windows":
		path = "~/Documents/My Games/Tabletop Simulator/Saves/Saved Objects/Community Maps"
	case "darwin":
		path = "~/Library/Tabletop Simulator/Saves/Saved Objects/Community Maps"
	case "openbsd":
		path = "~/maps"
	default:
		return "", fmt.Errorf("Operating system %v not supported", os)
	}

	return homedir.Expand(path)
}

func main() {
	path, err := getpath()
	if err != nil {
		log.Fatal(err)
	}

	r, err := git.PlainOpen(path)
	if err != nil {
		if err == git.ErrRepositoryNotExists {
			color.Yellow("Maps directory not found")
			if err := clone(path); err != nil {
				log.Fatal("Error cloning repository:", err)
			} else {
				color.Green("Done")
				os.Exit(0)
			}
		} else {
			log.Fatal(err)
		}
	}

	w, err := r.Worktree()
	if err != nil {
		log.Fatal(err)
	}

	color.Blue("Updating maps...")
	if err := w.Pull(&git.PullOptions{Progress: os.Stdout}); err != nil {
		if err == git.NoErrAlreadyUpToDate {
			color.Green("Already up to date")
		} else {
			log.Fatal(err)
		}
	}
}
