package myTheme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"

	resource "github.com/yuyuancha/gui/asset/resource"
)

type MyTheme struct{}

var _ fyne.Theme = (*MyTheme)(nil)

func (m MyTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(n, v)
}
func (m MyTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m MyTheme) Font(style fyne.TextStyle) fyne.Resource {
	return &fyne.StaticResource{
		StaticName:    "TaipeiSans.ttf",
		StaticContent: resource.Ttf,
	}
}

func (m MyTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
