package main

import (
	// "image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	// "gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	
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
	var toolBar widget.Clickable

	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err

		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			layout.Flex{
				// Vertical alignment, from top to bottom
				Axis: layout.Vertical,
				
			}.Layout(gtx,
				
				// The toolbar
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						toolBarBtn := material.Button(th, &toolBar, "Actions")
						return toolBarBtn.Layout(gtx)
					},
				),
				
				// The main textbox
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						padding := layout.UniformInset(unit.Dp(10))
						
						return padding.Layout(gtx, 
							func(gtx layout.Context) layout.Dimensions {
								textBox := material.Editor(th, &textInput, "")
								return textBox.Layout(gtx)
							})
					},
				),
			)

			

			e.Frame(gtx.Ops)
		}
	}// indirect
}