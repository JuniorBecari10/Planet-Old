package main

import (
	"context"
	"os"
	
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) OpenFileDialog() (string, error) {
	result, err := runtime.OpenFileDialog(
		a.ctx,
		runtime.OpenDialogOptions {},
	)

	if err != nil {
		return "", err
	}

	return result, nil
}

func (a *App) OpenDirDialog() (string, error) {
	result, err := runtime.OpenDirectoryDialog(
		a.ctx,
		runtime.OpenDialogOptions {
			ShowHiddenFiles: true,
			CanCreateDirectories: true,
		},
	)

	if err != nil {
		return "", err
	}

	return result, nil
}

func (a *App) ReadFile(path string) (string, error) {
	content, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(content), nil
}

func (a *App) SaveFile(path string, content string) error {
	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()
	
	_, err = file.WriteString(content)

	if err != nil {
		return err
	}

	err = file.Sync()
	if err != nil {
		return err
	}

	return nil
}
