package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/yuyuancha/gui/asset/myTheme"
	"image/color"
)

func main() {
	myApp := app.New()

	myApp.Settings().SetTheme(myTheme.MyTheme{})
	theme.AccountIcon()

	myWindow := myApp.NewWindow("打個招呼～")

	label := widget.NewLabel("哈囉，你好！")

	//setCircleExample(myCanvas, 255, 0, 0)
	myWindow.SetContent(label)

	myWindow.Resize(fyne.NewSize(300, 100))
	myWindow.ShowAndRun()
}

// setCircleExample 設定圓形範例
func setCircleExample(c fyne.Canvas, r, g, b uint8) {
	inputColor := color.NRGBA{R: r, G: g, B: b, A: 255}

	circle := canvas.NewCircle(color.White)

	circle.StrokeWidth = 4
	circle.StrokeColor = inputColor

	c.SetContent(circle)
}
