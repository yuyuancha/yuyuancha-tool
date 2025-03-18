package logic

import (
	"errors"
	"fmt"
	"github.com/yuyuancha-tool/pokemon-card-tool/apiCaller"
	"github.com/yuyuancha-tool/pokemon-card-tool/config"
	"github.com/yuyuancha-tool/pokemon-card-tool/contants"
	"gocv.io/x/gocv"
	"image"
	"os"
	"path/filepath"
	"strings"
)

// PokemonCardLogic 寶可夢卡牌邏輯
type PokemonCardLogic struct{}

// NewPokemonCardLogic new pokemon card logic
func NewPokemonCardLogic() *PokemonCardLogic {
	return &PokemonCardLogic{}
}

// GetUncollectedCards 取得未收集卡牌
func (logic *PokemonCardLogic) GetUncollectedCards() []string {
	var tempPathList []string
	err := filepath.Walk(contants.PokemonCardAssetsPath+config.PokemonCard.UncollectedCardsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if !strings.HasSuffix(info.Name(), ".png") {
			fmt.Println("錯誤的檔案格式:", info.Name())
			return nil
		}

		fmt.Println("開始處理圖片：", info.Name())

		tempPath, err := logic.handleCardScreenShots(info.Name(), path)
		if err != nil {
			return err
		}

		tempPathList = append(tempPathList, tempPath)

		return nil
	})
	if err != nil {
		fmt.Println("錯誤：", err)
	}

	fmt.Println("處理完成，開始 AI 識別數字......")

	results, err := logic.recognizeDigitsByGmini(tempPathList...)
	if err != nil {
		fmt.Println("錯誤：", err)
	}

	return results
}

// 處理截圖，並回傳暫存路徑
func (logic *PokemonCardLogic) handleCardScreenShots(imageName, imagePath string) (string, error) {
	img := gocv.IMRead(imagePath, gocv.IMReadColor)
	if img.Empty() {
		return "", errors.New("無法讀取影像")
	}

	defer func() { _ = img.Close() }()

	// 裁切
	tempImage := img.Region(image.Rect(0, img.Rows()/7, img.Cols(), img.Rows()))
	tempPath := contants.AssetsTempPath + "/tmp-" + imageName

	// 圖片暫存
	gocv.IMWrite(tempPath, tempImage.Clone())

	return tempPath, nil
}

// 透過 Gmini 辨識數字
func (logic *PokemonCardLogic) recognizeDigitsByGmini(pathList ...string) (numbers []string, _ error) {
	caller := apiCaller.NewGminiCaller()
	text, err := caller.RequestTextQuestionByImage(
		"請告訴我這些圖片裡的數字。1.以逗點分隔2.僅需回傳數字及逗點，無需回傳其他內容3.數字如果有零也需回傳4.重複的需要過濾掉。",
		pathList...)
	if err != nil {
		return nil, err
	}

	results := strings.Split(text, ",")

	for index, result := range results {
		switch {
		case len(result) == 1:
			results[index] = "00" + result
		case len(result) == 2:
			results[index] = "0" + result
		}
	}

	return results, nil
}
