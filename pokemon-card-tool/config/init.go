package config

import (
	"github.com/yuyuancha-tool/pokemon-card-tool/contants"
	"log"
	"slices"

	"github.com/spf13/viper"
)

// Driver 資料庫環境 Config
var Driver struct {
	Mysql mysqlStruct
}

// MySql Config 結構
type mysqlStruct struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

// GoogleService google service 結構
var GoogleService googleServiceStruct

// google service 結構
type googleServiceStruct struct {
	GminiApiKey         string
	HasGminiApiParseLog bool
}

// PokemonCard 寶可夢卡牌結構
var PokemonCard pokemonCardStruct

// pokemonCardStruct 寶可夢卡牌結構
type pokemonCardStruct struct {
	UncollectedCardSeries int
	UncollectedCardsDir   string
}

func init() {
	v := viper.New()
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		log.Fatalln("無法讀取配置文件:", err.Error())
	}

	// 配置資料庫環境 Config
	Driver = struct{ Mysql mysqlStruct }{
		Mysql: mysqlStruct{
			Host:     v.GetString("MYSQL_HOST"),
			Port:     v.GetInt("MYSQL_PORT"),
			User:     v.GetString("MYSQL_USER"),
			Password: v.GetString("MYSQL_PASSWORD"),
			Database: v.GetString("MYSQL_DATABASE"),
		},
	}

	// 配置 google service 結構
	GoogleService = googleServiceStruct{
		GminiApiKey:         v.GetString("GOOGLE_GMINI_API_KEY"),
		HasGminiApiParseLog: v.GetBool("HAS_GMINI_API_PARSE_LOG"),
	}

	// 配置 pokemonCard 結構
	PokemonCard = pokemonCardStruct{
		UncollectedCardsDir:   v.GetString("UNCOLLECTED_CARDS_DIR"),
		UncollectedCardSeries: v.GetInt("UNCOLLECTED_CARD_SERIES"),
	}
	if slices.Index([]int{contants.SeriesIdA2, contants.SeriesIdA2a}, PokemonCard.UncollectedCardSeries) == -1 {
		PokemonCard.UncollectedCardSeries = contants.SeriesIdA2
	}
}
