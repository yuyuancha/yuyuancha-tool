package service

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"

	model "github.com/yuyuancha/yuyuancha-tool/model/dogHero"
)

var dogHeroPwd string
var dogHeroMonthlyTargetTableNames = []string{
	"box", "recruit", "highCatch", "point", "oil", "food", "arena", "lowCatch",
}

// DogHeroMonthlyTargetLogic 狗英雄邏輯結構
type DogHeroMonthlyTargetLogic struct{}

func init() {
	dogHeroPwd, _ = os.Getwd()
}

// GetTargetList 取得達標清單
func (logic *DogHeroMonthlyTargetLogic) GetTargetList() (
	map[string][]model.DogHeroMonthlyTarget, error) {
	var results = make(map[string][]model.DogHeroMonthlyTarget, 0)

	for _, tableName := range dogHeroMonthlyTargetTableNames {
		items, err := logic.GetMonthlyTargetListFromCsv(tableName)
		if err != nil {
			return nil, err
		}

		results[tableName] = items
	}

	return results, nil
}

func (logic *DogHeroMonthlyTargetLogic) GetMonthlyTargetListFromCsv(tableName string) (
	[]model.DogHeroMonthlyTarget, error) {
	var csvName = fmt.Sprintf("/service/dogHero/file/monthlyTarget/%s-表格 1.csv", tableName)
	var path = filepath.Join(dogHeroPwd, csvName)

	// os.O_RDONLY 表示只讀、0777 表示(owner/group/other)權限
	var file, err = os.OpenFile(path, os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}

	var reader = csv.NewReader(file)
	var items []model.DogHeroMonthlyTarget

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		var progressing, _ = strconv.Atoi(record[0])
		var luckyBoxNumber, _ = strconv.Atoi(record[1])

		if progressing == 0 {
			continue
		}

		var item = model.DogHeroMonthlyTarget{
			Progressing:    progressing,
			LuckyBoxNumber: luckyBoxNumber,
		}

		items = append(items, item)
	}

	return items, nil
}
