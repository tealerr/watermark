package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	//Add background image
	bgImg, err := os.Open("vertical.jpg")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	//Decode Background Image
	decodeBgImg, err := jpeg.Decode(bgImg)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer bgImg.Close()

	//Add watermark image
	watermark, err := os.Open("highKycIdCardVerticalWatermark.png")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	//Decode Watermark
	decodeWatermark, err := png.Decode(watermark)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer watermark.Close()

	//Resize watermark for vertical image
	newHeighVertical := 1350
	verticalResizedImg := resize.Thumbnail(uint(decodeWatermark.Bounds().Dy()), uint(newHeighVertical), decodeWatermark, resize.Lanczos3)
	//when image is horizontal, use code below.

	//Resize watermark for horizontal image

	// newWidthHorizontal := 1350
	// horizonResizedImg := resize.Thumbnail(uint(newWidthHorizontal), uint(decodeWatermark.Bounds().Dy()), decodeWatermark, resize.Lanczos3)

	//Set Position vertical
	offset := image.Pt(200, 0)
	b := decodeBgImg.Bounds()
	dupImg := image.NewRGBA(b)
	draw.Draw(dupImg, b, decodeBgImg, image.ZP, draw.Src)
	draw.Draw(dupImg, verticalResizedImg.Bounds().Add(offset), verticalResizedImg, image.ZP, draw.Over)

	//when image is horizontal, use code below for set position.
	//Set Position horizontal

	// offset := image.Pt(20, 400)
	// b := decodeBgImg.Bounds()
	// dupImg := image.NewRGBA(b)
	// draw.Draw(dupImg, b, decodeBgImg, image.ZP, draw.Src)
	// draw.Draw(dupImg, verticalResizedImg.Bounds().Add(offset), horizonResizedImg, image.ZP, draw.Over)

	//Export image with watermark
	resultDupImg, err := os.Create("dupVerticalTest.jpg")
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	jpeg.Encode(resultDupImg, dupImg, &jpeg.Options{jpeg.DefaultQuality})
	defer resultDupImg.Close()
}
