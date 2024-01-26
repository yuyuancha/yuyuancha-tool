package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ShowStreesSlotUI() *fyne.Container {
	content = container.NewHBox(
		widget.NewLabel("平台"),
		widget.NewLabel("代理"),
	)

	c := container.New(layout.NewVBoxLayout(), content)

	return c
}
