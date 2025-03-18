package main

const (
	capturePath               = "./img/card1.png"
	mattingImagePath          = "./img/matting2.png"
	templatePath              = "./img/template.png"
	tesseractDigitalWhitelist = "0123456789"
	tesseractTempPath         = "./img/tmp/temp_number.png"
	tempPath                  = "./img/tmp/temp_number.png"
)

func main() {
	logic := NewGocvLogic()
	logic.GetUncollectedCardsByGmini()
}
