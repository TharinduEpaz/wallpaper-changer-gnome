package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type BingResponse struct {
	URL string `json:"url"`
}

func main() {

	// Download the image
	imgResp, err := http.Get("https://unsplash.it/1920/1080/?random")
	if err != nil {
		fmt.Println("Error downloading image:", err)
		return
	}
	defer imgResp.Body.Close()

	// Create directory if it doesn't exist
	homeDir, _ := os.UserHomeDir()
	wallpaperDir := filepath.Join(homeDir, ".wallpapers")
	os.MkdirAll(wallpaperDir, 0755)

	// Save the image
	wallpaperPath := filepath.Join(wallpaperDir, "unsplash-wallpaper.jpg")
	imgFile, err := os.Create(wallpaperPath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer imgFile.Close()

	if _, err := io.Copy(imgFile, imgResp.Body); err != nil {
		fmt.Println("Error saving image:", err)
		return
	}

	// Set wallpaper using gsettings
	cmd := exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri-dark", 
		fmt.Sprintf("file://%s", wallpaperPath))
	if err := cmd.Run(); err != nil {
		fmt.Println("Error setting wallpaper:", err)
		return
	}

	fmt.Println("Wallpaper updated successfully!")
}