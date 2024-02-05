package main

import (
	"fmt"
	"image"
	"log"
	"os"

	_ "image/jpeg"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func main() {
	file, err := os.Open("qrcode.jpg")
	if err != nil {
		log.Fatalln("Open file error:", err.Error())
	} else if file == nil {
		log.Fatalln("File is nil.")
	}

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalln("Decode image error:", err.Error())
	}

	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		log.Fatalln("Gozxing NewBinaryBitmapFromImage error:", err.Error())
	}

	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		log.Fatalln("Decode qr code reader error:", err.Error())
	}

	fmt.Println(result)
}
