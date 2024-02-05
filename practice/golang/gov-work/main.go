package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/yuyuancha/gov-work/asset/myTheme"
	"github.com/yuyuancha/gov-work/view"
)

func main() {
	myApp := app.New()

	myApp.Settings().SetTheme(myTheme.MyTheme{})
	theme.AccountIcon()

	myWindow := myApp.NewWindow("行政工作自動化程序")

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("抓取資料夾路名", theme.ContentPasteIcon(), view.ShowFolderNameForLocation()),
		//container.NewTabItemWithIcon("SLOT測試", theme.ComputerIcon(), view.ShowStreesSlotUI()),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	myWindow.SetContent(tabs)

	myWindow.Resize(fyne.NewSize(800, 800))
	myWindow.ShowAndRun()
}
