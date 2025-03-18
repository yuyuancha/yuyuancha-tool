package contants

const (
	PokemonCardAssetsPath = "./assets/pokemon-card/"

	SeriesIdA2  = 1 // A2 系列：時間激鬥
	SeriesIdA2a = 2 // A2a 系列：超克之光
)

// SeriesWebCrawlerChildrenId 爬蟲的子節點 id
var SeriesWebCrawlerChildrenId = map[int]int{
	SeriesIdA2:  9,
	SeriesIdA2a: 11,
}

// SeriesWebCrawlerUrl 爬蟲的網址
var SeriesWebCrawlerUrl = map[int]string{
	SeriesIdA2:  "https://wiki.52poke.com/wiki/%E6%99%82%E7%A9%BA%E6%BF%80%E9%AC%A5%EF%BC%88TCGP%EF%BC%89#%E5%9B%BE%E5%86%8C",
	SeriesIdA2a: "https://wiki.52poke.com/wiki/%E8%B6%85%E5%85%8B%E4%B9%8B%E5%85%89%EF%BC%88TCGP%EF%BC%89",
}
