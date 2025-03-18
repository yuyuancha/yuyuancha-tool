package main

import (
	"fmt"
	"github.com/otiai10/gosseract/v2"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"sort"
	"strings"
)

type GocvLogic struct{}

// NewGocvLogic new gocv logic
func NewGocvLogic() *GocvLogic {
	return &GocvLogic{}
}

// GetUncollectedCards 取得未收集的卡片
func (logic *GocvLogic) GetUncollectedCards() {
	img := gocv.IMRead(capturePath, gocv.IMReadColor)
	if img.Empty() {
		panic("無法讀取影像")
	}

	defer func() { _ = img.Close() }()

	// 灰階
	gray := gocv.NewMat()
	defer func() { _ = gray.Close() }()
	gocv.CvtColor(img, &gray, gocv.ColorBGRToGray)

	// 二值化
	binary := gocv.NewMat()
	defer func() { _ = binary.Close() }()
	gocv.Threshold(gray, &binary, 200, 220, gocv.ThresholdBinaryInv)

	// 尋找輪廓
	contours := gocv.FindContours(binary, gocv.RetrievalExternal, gocv.ChainApproxSimple)

	// 整理組合輪廓
	combinedRects := logic.tidyContours(contours)
	for i, combinedRect := range combinedRects {
		if i > 8 {
			continue
		}

		logic.checkRegionSize(&combinedRect, &img)
		numberROI := binary.Region(combinedRect)

		text, err := logic.recognizeDigits(&numberROI)
		if err != nil {
			fmt.Println("辨識失敗")
			continue
		}
		_ = numberROI.Close()

		// 在原圖上標記辨識結果
		gocv.Rectangle(&img, combinedRect, color.RGBA{255, 0, 0, 0}, 2)
		gocv.PutText(&img, text,
			image.Pt(combinedRect.Min.X, combinedRect.Min.Y-5),
			gocv.FontHersheySimplex, 0.8,
			color.RGBA{0, 0, 255, 0}, 2)
	}

	window := gocv.NewWindow("Card Number Detection")
	defer func() { _ = window.Close() }()

	for {
		window.IMShow(img)
		if window.WaitKey(1) == 27 {
			break
		}
	}
}

// GetUncollectedCardsByGmini 透過 Gmini 取得未收集的卡片
func (logic *GocvLogic) GetUncollectedCardsByGmini() {
	img := gocv.IMRead(capturePath, gocv.IMReadColor)
	if img.Empty() {
		panic("無法讀取影像")
	}

	defer func() { _ = img.Close() }()

	// 灰階
	gray := gocv.NewMat()
	defer func() { _ = gray.Close() }()
	gocv.CvtColor(img, &gray, gocv.ColorBGRToGray)

	// 二值化
	binary := gocv.NewMat()
	defer func() { _ = binary.Close() }()
	gocv.Threshold(gray, &binary, 200, 220, gocv.ThresholdBinaryInv)

	// 圖片暫存
	gocv.IMWrite(tempPath, binary)

	numbers, err := logic.recognizeDigitsByGmini(tempPath)
	if err != nil {
		fmt.Println("辨識失敗", err.Error())
		return
	}

	fmt.Println("辨識結果：", numbers)
}

// 整理組合輪廓
func (logic *GocvLogic) tidyContours(contours gocv.PointsVector) []image.Rectangle {
	rectMap := logic.getContourRectMap(contours)

	for key, rects := range rectMap {
		// 先按照 X 做排序
		sort.Slice(rects, func(i, j int) bool {
			return rects[i].Min.X < rects[j].Min.X
		})

		// 剛好三個一組的跳過分組
		if len(rects) == 3 {
			continue
		}

		for i := 0; i < len(rects)-2; i++ {
			// 判斷後面 1-2 個的距離是否是數字的距離
			d := 0
			var combination []image.Rectangle
			if rects[i+1].Min.X < rects[i].Min.X+50 {
				d++
				if rects[i+2].Min.X < rects[i+1].Min.X+50 {
					d++
				}
			}

			if d == 0 {
				continue
			}

			combination = append(combination, rects[i:i+d+1]...)

			rectMap[key*10+i] = combination
			i += d
		}
	}
	for k := range rectMap {
		if len(rectMap[k]) > 3 {
			delete(rectMap, k)
		}
	}

	var combineRects []image.Rectangle
	for k := range rectMap {
		pv := gocv.NewPointVector()
		for _, rect := range rectMap[k] {
			pv.Append(image.Point{X: rect.Min.X - 10, Y: rect.Min.Y - 10})
			pv.Append(image.Point{X: rect.Max.X + 10, Y: rect.Max.Y + 10})
		}
		combinedRect := gocv.BoundingRect(pv)
		if combinedRect.Max.X-combinedRect.Min.X > 200 {
			continue
		}
		combineRects = append(combineRects, combinedRect)
	}

	return combineRects
}

// 取得輪廓矩形集合
func (logic *GocvLogic) getContourRectMap(contours gocv.PointsVector) map[int][]image.Rectangle {
	var (
		rects       = logic.getContoursRects(contours)
		minMap      = make(map[int][]image.Rectangle)
		maxMap      = make(map[int][]image.Rectangle)
		remainRects []image.Rectangle
	)

	// 將 min Y 相同的組一起
	for _, rect := range rects {
		minMap[rect.Min.Y] = append(minMap[rect.Min.Y], rect)
	}

	// 相同小於三個的忽略
	for k, v := range minMap {
		if len(minMap[k]) < 3 {
			remainRects = append(remainRects, v...)
			delete(minMap, k)
		}
	}

	// 將 max Y 相同的組一起
	for _, rect := range remainRects {
		maxMap[rect.Max.Y] = append(maxMap[rect.Max.Y], rect)
	}

	// 相同小於三個的忽略
	for k := range maxMap {
		if len(maxMap[k]) < 3 {
			delete(maxMap, k)
		}
	}

	// 將重複的忽略
	for k := range maxMap {
		if _, ok := minMap[k]; ok {
			continue
		}
		minMap[k] = maxMap[k]
	}

	return minMap
}

// 取得輪廓的矩形
func (logic *GocvLogic) getContoursRects(contours gocv.PointsVector) []image.Rectangle {
	var rects []image.Rectangle
	for i := 0; i < contours.Size(); i++ {
		contour := contours.At(i)
		rect := gocv.BoundingRect(contour)
		if rect.Dx() < 90 && rect.Dy() < 90 {
			rects = append(rects, rect)
		}
	}
	return rects
}

// 檢查 ROI 是否超出圖像範圍
func (logic *GocvLogic) checkRegionSize(rect *image.Rectangle, mat *gocv.Mat) {
	if rect.Min.X < 0 {
		rect.Min.X = 0
	}
	if rect.Min.Y < 0 {
		rect.Min.Y = 0
	}

	// 調整 width 和 height
	if rect.Max.X > mat.Cols() {
		rect.Max.X = mat.Cols()
	}
	if rect.Max.Y > mat.Rows() {
		rect.Max.Y = mat.Rows()
	}
}

// 辨識數字
func (logic *GocvLogic) recognizeDigits(mat *gocv.Mat) (text string, err error) {
	// 儲存暫存圖片 (給 Tesseract 辨識用)
	gocv.IMWrite(tesseractTempPath, *mat)

	client := gosseract.NewClient()
	err = client.SetWhitelist(tesseractDigitalWhitelist)
	if err != nil {
		return "", err
	}

	// 使用 Tesseract 辨識數字
	_ = client.SetImage(tesseractTempPath)
	text, err = client.Text()
	fmt.Printf("辨識結果：%s\n", text)

	_ = client.Close()

	return
}

// 透過 Gmini 辨識數字
func (logic *GocvLogic) recognizeDigitsByGmini(path string) (numbers []string, _ error) {
	caller := NewGminiCaller()
	text, err := caller.RequestTextQuestionByImage("請告訴我圖片裡的數字，以逗點分隔，僅數字及逗點，無需回傳其他文字。", path)
	if err != nil {
		return nil, err
	}

	results := strings.Split(text, ",")

	return results, nil
}
