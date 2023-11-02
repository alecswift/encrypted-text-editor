package main

import (
	// "image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/example/component/icon"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	// "gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	
)

func main() {
	go func() {
		w := app.NewWindow(
			app.Title("Text Editor"),
			app.Size(unit.Dp(1024), unit.Dp(768)),
		)
		err := run(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(w *app.Window) error {
	th := material.NewTheme()
	var ops op.Ops
	var textInput widget.Editor

	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err

		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			
			modal := component.NewModal()
			appBar := component.NewAppBar(modal)
			appBar.NavigationIcon = icon.MenuIcon
			appBar.Layout(gtx, th, "Actions", "")

			inset := layout.UniformInset(unit.Dp(25))
			inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				textBox := material.Editor(th, &textInput, "")
				return textBox.Layout(gtx)
			})

			e.Frame(gtx.Ops)
		}
	}// indirect
}