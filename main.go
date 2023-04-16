package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App {
		Title:  "Planet",
		Width:  1240,
		Height: 740,
		AssetServer: &assetserver.Options {
			Assets: assets,
		},
		WindowStartState:   options.Maximised,
		BackgroundColour: &options.RGBA{R: 18, G: 22, B: 29, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
