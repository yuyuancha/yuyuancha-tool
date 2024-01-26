package view

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/yuyuancha/gov-work/tool"
	"os"
	"strings"
)

// folderNameForLocation 資料夾路名結構
type folderNameForLocation struct {
	folderLabel   *widget.Label
	locationEntry *widget.Entry
}

var selectedFoldersForLocation []string

// ShowFolderNameForLocation Tab 頁面：取得資料夾路名
func ShowFolderNameForLocation() *fyne.Container {
	pwd, _ := os.Getwd()
	folderNames := tool.GetFirstLevelFolderNames(pwd)

	selectedFolderNameLabel := widget.NewLabel("")

	checkGroups := widget.NewCheckGroup(folderNames, func(strs []string) {
		selectedFoldersForLocation = strs
	})

	folderNameScroll := container.NewHScroll(checkGroups)
	folderNameScroll.Resize(fyne.NewSize(100, 50))

	resultContainer := container.NewVBox()

	startBtn := widget.NewButton("開始抓取", func() {
		resultContainer.RemoveAll()
		selectedFolderNameLabel.SetText(fmt.Sprintf("%s", strings.Join(selectedFoldersForLocation, "、")))
		results := startCatchFolderLocationName()
		for _, result := range results {
			resultContainer.Add(result.folderLabel)
			resultContainer.Add(result.locationEntry)
		}
	})
	startBtn.Resize(fyne.NewSize(100, 50))

	content = container.NewVBox(
		widget.NewLabel("抓取資料夾"),
		folderNameScroll,
		startBtn,
		widget.NewLabel("已選擇資料夾："),
		selectedFolderNameLabel,
		widget.NewLabel("抓取結果："),
		resultContainer,
	)

	c := container.New(layout.NewVBoxLayout(), content)

	return c
}

// startCatchFolderLocationName 開始抓取資料夾路名
func startCatchFolderLocationName() []folderNameForLocation {
	var results []folderNameForLocation
	pwd, _ := os.Getwd()

	for _, folderName := range selectedFoldersForLocation {
		path := fmt.Sprintf("%s/%s", pwd, folderName)
		names := tool.GetFirstLevelFolderNames(path)

		for index, name := range names {
			names[index] = tool.SplitStrByRegex(name, `\d{8}\s\s`)
		}

		result := folderNameForLocation{
			folderLabel:   widget.NewLabel(folderName),
			locationEntry: widget.NewEntry(),
		}

		result.locationEntry.Text = strings.Join(names, "、")
		result.locationEntry.Size()

		results = append(results, result)
	}

	return results
}
