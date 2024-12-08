package renderer

import "github.com/gdamore/tcell/v2"

// ColorMode defines different color rendering modes
type ColorMode byte

const (
	ColorModeMonochrome ColorMode = 0x00
	ColorMode16         ColorMode = 0x01
	ColorMode256        ColorMode = 0x02
)

// MapColor converts color index to tcell color based on color mode
func MapColor(mode ColorMode, colorIndex byte) tcell.Color {
	switch mode {
	case ColorModeMonochrome:
		return tcell.ColorWhite
	case ColorMode16:
		// Basic 16-color palette
		colors := []tcell.Color{
			tcell.ColorBlack, tcell.ColorMaroon, tcell.ColorGreen,
			tcell.ColorOlive, tcell.ColorNavy, tcell.ColorPurple,
			tcell.ColorTeal, tcell.ColorSilver,
			tcell.ColorGray, tcell.ColorRed, tcell.ColorLime,
			tcell.ColorYellow, tcell.ColorBlue, tcell.ColorFuchsia,
			tcell.ColorAqua, tcell.ColorWhite,
		}
		if int(colorIndex) < len(colors) {
			return colors[colorIndex]
		}
		return tcell.ColorWhite
	case ColorMode256:
		// Simplistic 256-color mapping
		return tcell.PaletteColor(colorIndex)
	default:
		return tcell.ColorWhite
	}
}
