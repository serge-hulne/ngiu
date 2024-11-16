package ngiu

import (
	"image/color"

	"github.com/AllenDang/giu"
)

func (c Color) ToRGBA() color.RGBA {
	return color.RGBA{R: c.R, G: c.G, B: c.B, A: c.A}
}

func Spacer(w, h float32) giu.Widget {
	return giu.Dummy(w, h)
}

func VSpacer(h float32) giu.Widget {
	return giu.Dummy(0, h)
}

func HStack(widgets ...giu.Widget) giu.Widget {
	return giu.Row(widgets...)
}

func VStack(widgets ...giu.Widget) giu.Widget {
	return giu.Column(widgets...)
}

// Define a custom struct that holds a giu.Widget
type Widget struct {
	giu.Widget
}

// Constructor for creating a CustomWidget
func NewWidget(w giu.Widget) *Widget {
	return &Widget{w}
}

func (widget *Widget) Padding(top, right, bottom, left float32) *Widget {
	return NewWidget(giu.Layout{
		giu.Dummy(0, top), // Top padding
		giu.Row(
			giu.Dummy(left, 0), // Left padding
			giu.Column(
				widget,
			),
			giu.Dummy(right, 0), // Right padding
		),
		giu.Dummy(0, bottom), // Bottom padding
	})
}

func (widget *Widget) PaddingTop(top float32) *Widget {
	return NewWidget(giu.Layout{
		giu.Dummy(0, top), // Top padding
		giu.Row(
			giu.Column(
				widget,
			),
		),
	})
}

func (widget *Widget) PaddingLeft(left float32) *Widget {
	return NewWidget(giu.Layout{
		giu.Row(
			giu.Dummy(left, 0), // Left padding
			giu.Column(
				widget,
			),
		),
	})
}

// Font weight (by choosing a larger or smaller font size)
func (widget *Widget) FontWeight(weight string) *Widget {
	var fontSize float32
	switch weight {
	case "bold":
		fontSize = 18
	case "medium":
		fontSize = 14
	case "light":
		fontSize = 10
	default:
		fontSize = 14 // default to medium
	}
	return NewWidget(giu.Style().
		SetFontSize(fontSize).
		To(widget.Widget),
	)
}

// Status bar
// TODO : put in lib (ngiu)
func StatusBar(text string, height float32) giu.Widget {
	// Approximate height of the status label

	// Spacer to dynamically push the status label to the bottom
	spacer := giu.Custom(func() {
		_, availableHeight := giu.GetAvailableRegion()
		giu.Dummy(0, availableHeight-height).Build()
	})

	// Status label at the bottom
	statusLabel := n.NewWidget(giu.Label(text)).
		BackgroundColor(n.NewColorByName(n.ColorNames.Basic.Blue)).
		ForegroundColor(n.NewColorByName(n.ColorNames.Vivid.Lime)).
		FontWeight("bold").
		PaddingTop(10)

	return giu.Layout{spacer, statusLabel}
}

