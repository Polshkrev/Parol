package themes

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// Custom card theme.
type Card struct{}

// Construct a new card theme.
func NewCard() *Card {
	return new(Card)
}

// Change the colour of the theme.
// Returns a colour based on the given name and variant.
func (Card) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if variant == theme.VariantDark && name == theme.ColorNameShadow {
		return color.White
	}
	return theme.DefaultTheme().Color(name, variant)
}

// Font method override.
// Returns a [fyne.Resource] of the given style.
func (Card) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

// Icon method override.
// Returns a [fyne.Resource] of the given name.
func (Card) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

// Size method override.
// Returns the size of the current theme.
func (Card) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)

}
