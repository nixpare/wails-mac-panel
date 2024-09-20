package main

import (
	"embed"
	"log"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
	"golang.design/x/hotkey"
)

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

//go:embed frontend/dist
var assets embed.FS

var (
	app *application.App
	window *application.WebviewWindow
	panel *application.WebviewPanel
)

// main function serves as the application's entry point. It initializes the application, creates a window,
// and starts a goroutine that emits a time-based event every second. It subsequently runs the application and
// logs any error that might occur.
func main() {
	// Create a new Wails application by providing the necessary options.
	// Variables 'Name' and 'Description' are for application metadata.
	// 'Assets' configures the asset server with the 'FS' variable pointing to the frontend files.
	// 'Bind' is a list of Go struct instances. The frontend has access to the methods of these instances.
	// 'Mac' options tailor the application when running an macOS.
	app = application.New(application.Options{
		Name:        "gui",
		Description: "A demo of using raw HTML & CSS",
		Services: []application.Service{
			application.NewService(&GreetService{}),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: false,
		},
	})

	// Register the hotkey (CTRL + CMD + Space) that will create a new panel or hide/focus
	// the already existing one.
	// When it creates the new panel, it calls the function [createPanel].
	// The panel created will float in front of any application, even those in fullscreen without
	// switching back to the desktop, until it loses focus.
	go setupHotkey()

	// Create a goroutine that emits an event containing the current time every second.
	// The frontend can listen to this event and update the UI accordingly.
	go func() {
		for {
			now := time.Now().Format(time.RFC1123)
			app.EmitEvent("time", now)
			time.Sleep(time.Second)
		}
	}()

	// Run the application. This blocks until the application has been exited.
	err := app.Run()

	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Fatal(err)
	}
}

var (
	defaultWindowOptions = application.WebviewWindowOptions{
		Title: "GUI",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(27, 38, 54),
		URL:              "/",
		KeyBindings: map[string]func(_ *application.WebviewWindow){
			"ctrl+m": func(_ *application.WebviewWindow) {
				log.Println("CTRL+M")
			},
			"ctrl+enter": func(_ *application.WebviewWindow) {
				log.Println("CTRL+ENTER")
			},
			"ctrl+a": func(_ *application.WebviewWindow) {
				log.Println("CTRL+A")
			},
			"ctrl+l": func(_ *application.WebviewWindow) {
				log.Println("CTRL+L")
			},
		},
	}

	defaultPanelOptions = application.WebviewPanelOptions{
		WebviewWindowOptions: defaultWindowOptions,
		Floating: true,
	}
)

func createPanel() {
	panel = app.NewWebviewPanelWithOptions(defaultPanelOptions)

	panel.OnWindowEvent(events.Common.WindowClosing, func(event *application.WindowEvent) {
		panel = nil
	})
	panel.OnWindowEvent(events.Common.WindowLostFocus, func(event *application.WindowEvent) {
		panel.Hide()
	})
}

func setupHotkey() {
	showHideHotkey := hotkey.New([]hotkey.Modifier{ hotkey.ModCtrl, hotkey.ModCmd }, hotkey.KeySpace)
	if err := showHideHotkey.Register(); err != nil {
		log.Println(err)
		return
	}
	
	go func() {
		for range showHideHotkey.Keydown() {
			if panel == nil {
				createPanel()
			}

			if panel.IsVisible() {
				panel.Hide()
			} else {
				panel.Focus()
			}
		}
	}()

	app.OnShutdown(func() {
		if err := showHideHotkey.Unregister(); err != nil {
			log.Println(err)
		}
	})
}
