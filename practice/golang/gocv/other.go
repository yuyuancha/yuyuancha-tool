package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
)

func capture() {
	// 讀取原圖:
	img := gocv.IMRead(capturePath, gocv.IMReadColor)
	// 讀取灰階圖：
	//img := gocv.IMRead(capturePath, gocv.IMReadGrayScale)

	gray := gocv.NewMat()
	defer func() {
		_ = gray.Close()
	}()

	gocv.CvtColor(img, &gray, gocv.ColorBGRToGray)

	// 讀取模板圖片
	template := gocv.IMRead(templatePath, gocv.IMReadGrayScale)
	if template.Empty() {
		panic("無法讀取模板圖片")
	}
	defer template.Close()

	// 準備空遮罩 (不使用遮罩)
	mask := gocv.NewMat()
	defer mask.Close()

	// 執行模板匹配
	result := gocv.NewMat()
	defer result.Close()
	gocv.MatchTemplate(gray, template, &result, gocv.TmCcoeffNormed, mask)

	// 設定相似度閾值
	threshold := 0.8
	_, maxVal, _, maxLoc := gocv.MinMaxLoc(result)
	if float64(maxVal) >= threshold {
		// 匹配成功，畫出矩形框
		topLeft := maxLoc
		bottomRight := image.Pt(topLeft.X+template.Cols(), topLeft.Y+template.Rows())
		gocv.Rectangle(&img, image.Rect(topLeft.X, topLeft.Y, bottomRight.X, bottomRight.Y), color.RGBA{255, 0, 0, 0}, 4)

		fmt.Println("偵測到：template")
	} else {
		fmt.Println("未偵測到：template")
	}

	window := gocv.NewWindow("capture")
	window.ResizeWindow(800, 600)
	window.IMShow(img)
	window.WaitKey(0)
	_ = window.Close()
}

// 圖像去背
func mattingImage() {
	img := gocv.IMRead(mattingImagePath, gocv.IMReadColor)
	if img.Empty() {
		panic("無法讀取影像")
	}
	defer img.Close()

	// 初始化蒙版
	mask := gocv.NewMatWithSize(img.Rows(), img.Cols(), gocv.MatTypeCV8UC1)
	defer mask.Close()

	// 初始化背景與前景模型
	bgModel := gocv.NewMat()
	defer bgModel.Close()

	fgModel := gocv.NewMat()
	defer fgModel.Close()

	// 定義前景區域 (以矩形方式框住)
	rect := image.Rect(10, 10, img.Cols()-10, img.Rows()-10)

	// 執行 GrabCut
	gocv.GrabCut(img, &mask, rect, &bgModel, &fgModel, 5, gocv.GCInitWithRect)

	// 將可能前景和確定前景設為白色
	fgd := gocv.NewMat()
	defer fgd.Close()
	gocv.Compare(mask, gocv.NewMatFromScalar(gocv.Scalar{Val1: float64(gocv.GCInitWithRect)}, gocv.MatTypeCV8UC1), &fgd, gocv.CompareEQ)
	gocv.Compare(mask, gocv.NewMatFromScalar(gocv.Scalar{Val1: float64(gocv.GCEval)}, gocv.MatTypeCV8UC1), &fgd, gocv.CompareEQ)

	// 建立輸出影像
	result := gocv.NewMat()
	defer result.Close()
	img.CopyToWithMask(&result, fgd)

	// 顯示結果
	window := gocv.NewWindow("GrabCut Result")
	defer window.Close()

	window.IMShow(result)
	window.WaitKey(0)
	_ = window.Close()
}
