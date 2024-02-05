package service

import model "github.com/yuyuancha/yuyuancha-tool/model/gov"

// GetGovTravelCardShops 取得政府旅遊卡店家
func (logic *GovLogic) GetGovTravelCardShops(categoryId int) ([]model.GovTravelCardShop, error) {
	var shopModel model.GovTravelCardShop
	if categoryId == -1 {
		return shopModel.FindAll()
	}

	return shopModel.FindAllByCategoryId(categoryId)
}

// GetTravelCategories 取得旅遊類別
func (logic *GovLogic) GetTravelCategories() ([]model.GovTravelCardCategory, error) {
	var categoryModel model.GovTravelCardCategory
	return categoryModel.FindAll()
}
