package main

import (
	"fmt"
	"os"

	"github.com/AllenDang/giu"
)

// Color struct to hold RGBA values
type Color struct {
	R, G, B, A uint8
}

func NewColor(r, g, b, a uint8) Color {
	return Color{R: r, G: g, B: b, A: a}
}

// ColorName type to act as an "enum" for color names
type ColorName string

// ColorNames struct to group color constants by palette
var ColorNames = struct {
	Basic struct {
		Red, Green, Blue, Black, White, Yellow, Cyan, Magenta, Gray ColorName
	}
	Pastel struct {
		Pink, Peach, Lavender, Mint, Yellow, Blue ColorName
	}
	Vivid struct {
		Orange, Purple, Turquoise, Lime, Pink, Gold ColorName
	}
	Neutral struct {
		Charcoal, Cream, Beige, Taupe, Olive, Slate ColorName
	}
	Additional struct {
		Orange, Purple, Brown, Teal, Indigo, Maroon, Navy, Gold ColorName
	}
}{
	Basic: struct {
		Red, Green, Blue, Black, White, Yellow, Cyan, Magenta, Gray ColorName
	}{
		Red: "basic_red", Green: "basic_green", Blue: "basic_blue", Black: "basic_black",
		White: "basic_white", Yellow: "basic_yellow", Cyan: "basic_cyan",
		Magenta: "basic_magenta", Gray: "basic_gray",
	},
	Pastel: struct {
		Pink, Peach, Lavender, Mint, Yellow, Blue ColorName
	}{
		Pink: "pastel_pink", Peach: "pastel_peach", Lavender: "pastel_lavender",
		Mint: "pastel_mint", Yellow: "pastel_yellow", Blue: "pastel_blue",
	},
	Vivid: struct {
		Orange, Purple, Turquoise, Lime, Pink, Gold ColorName
	}{
		Orange: "vivid_orange", Purple: "vivid_purple", Turquoise: "vivid_turquoise",
		Lime: "vivid_lime", Pink: "vivid_pink", Gold: "vivid_gold",
	},
	Neutral: struct {
		Charcoal, Cream, Beige, Taupe, Olive, Slate ColorName
	}{
		Charcoal: "neutral_charcoal", Cream: "neutral_cream", Beige: "neutral_beige",
		Taupe: "neutral_taupe", Olive: "neutral_olive", Slate: "neutral_slate",
	},
	Additional: struct {
		Orange, Purple, Brown, Teal, Indigo, Maroon, Navy, Gold ColorName
	}{
		Orange: "orange", Purple: "purple", Brown: "brown", Teal: "teal",
		Indigo: "indigo", Maroon: "maroon", Navy: "navy", Gold: "gold",
	},
}

// Color map using ColorName as keys and Color as values
var colorMap = map[ColorName]Color{
	// Basic Colors
	ColorNames.Basic.Red:     NewColor(255, 0, 0, 255),
	ColorNames.Basic.Green:   NewColor(0, 255, 0, 255),
	ColorNames.Basic.Blue:    NewColor(0, 0, 255, 255),
	ColorNames.Basic.Black:   NewColor(0, 0, 0, 255),
	ColorNames.Basic.White:   NewColor(255, 255, 255, 255),
	ColorNames.Basic.Yellow:  NewColor(255, 255, 0, 255),
	ColorNames.Basic.Cyan:    NewColor(0, 255, 255, 255),
	ColorNames.Basic.Magenta: NewColor(255, 0, 255, 255),
	ColorNames.Basic.Gray:    NewColor(128, 128, 128, 255),

	// Pastel Colors
	ColorNames.Pastel.Pink:     NewColor(255, 182, 193, 255),
	ColorNames.Pastel.Peach:    NewColor(255, 218, 185, 255),
	ColorNames.Pastel.Lavender: NewColor(230, 230, 250, 255),
	ColorNames.Pastel.Mint:     NewColor(189, 252, 201, 255),
	ColorNames.Pastel.Yellow:   NewColor(253, 253, 150, 255),
	ColorNames.Pastel.Blue:     NewColor(173, 216, 230, 255),

	// Vivid Colors
	ColorNames.Vivid.Orange:    NewColor(255, 69, 0, 255),
	ColorNames.Vivid.Purple:    NewColor(148, 0, 211, 255),
	ColorNames.Vivid.Turquoise: NewColor(64, 224, 208, 255),
	ColorNames.Vivid.Lime:      NewColor(191, 255, 0, 255),
	ColorNames.Vivid.Pink:      NewColor(255, 20, 147, 255),
	ColorNames.Vivid.Gold:      NewColor(255, 215, 0, 255),

	// Neutral Colors
	ColorNames.Neutral.Charcoal: NewColor(54, 69, 79, 255),
	ColorNames.Neutral.Cream:    NewColor(245, 245, 220, 255),
	ColorNames.Neutral.Beige:    NewColor(245, 245, 220, 255),
	ColorNames.Neutral.Taupe:    NewColor(72, 60, 50, 255),
	ColorNames.Neutral.Olive:    NewColor(128, 128, 0, 255),
	ColorNames.Neutral.Slate:    NewColor(112, 128, 144, 255),

	// Additional Colors
	ColorNames.Additional.Orange: NewColor(255, 165, 0, 255),
	ColorNames.Additional.Purple: NewColor(128, 0, 128, 255),
	ColorNames.Additional.Brown:  NewColor(165, 42, 42, 255),
	ColorNames.Additional.Teal:   NewColor(0, 128, 128, 255),
	ColorNames.Additional.Indigo: NewColor(75, 0, 130, 255),
	ColorNames.Additional.Maroon: NewColor(128, 0, 0, 255),
	ColorNames.Additional.Navy:   NewColor(0, 0, 128, 255),
	ColorNames.Additional.Gold:   NewColor(255, 215, 0, 255),
}

// NewColorByName function to retrieve colors by name
func NewColorByName(name ColorName) Color {
	if color, exists := colorMap[name]; exists {
		return color
	}
	fmt.Printf("Error: color '%s' not found\n", name)
	os.Exit(1)
	return NewColor(255, 0, 0, 255) // Default color (not expected to be reached)
}

// Foreground color (text color)
func (widget *Widget) ForegroundColor(color Color) *Widget {
	return NewWidget(giu.Style().
		SetColor(giu.StyleColorText, color.ToRGBA()).
		To(widget.Widget),
	)
}

// Background
func (widget *Widget) BackgroundColor(color Color) *Widget {
	return NewWidget(giu.Style().
		SetColor(giu.StyleColorChildBg, color.ToRGBA()). // Set background color
		To(
			giu.Child().Border(false).Size(-1, 20).Layout(giu.Layout{widget.Widget}), // Fixed height of 30
		),
	)
}
