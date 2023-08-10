package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/disintegration/imaging"
)

func main() {
	//Add background image
	bgImg, err := os.Open("backgroundImage/vertical.jpg")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	//Decode Background Image
	decodeBgImg, err := jpeg.Decode(bgImg)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer bgImg.Close()

	watermark, err := os.Open("kycWatermark/highKycIdCardVerticalWatermark.png")
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

	verticalResizedImg := imaging.Resize(decodeWatermark, 0, 1350, imaging.Lanczos)

	//when image is horizontal, use code below.
	//Resize watermark for horizontal image

	// horizonResizedImg := imaging.Resize(decodeWatermark, 1350, 0, imaging.Lanczos)

	//Set Position vertical
	offset := image.Pt(200, 0)
	b := decodeBgImg.Bounds()
	mergeImg := image.NewRGBA(b)

	//Create merge image with watermark
	draw.Draw(mergeImg, b, decodeBgImg, image.ZP, draw.Src)
	draw.Draw(mergeImg, verticalResizedImg.Bounds().Add(offset), verticalResizedImg, image.ZP, draw.Over)

	//when image is horizontal, use code below for set position.
	//Set Position horizontal

	// offset := image.Pt(20, 400)
	// b := decodeBgImg.Bounds()
	// mergeImg := image.NewRGBA(b)
	// draw.Draw(mergeImg, b, decodeBgImg, image.ZP, draw.Src)
	// draw.Draw(mergeImg, verticalWatermarkResized.Bounds().Add(offset), horizonWatermarkResized, image.ZP, draw.Over)

	//Export image with watermark
	resultMergeImg, err := os.Create("dupVerticalTest83.jpg")
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	jpeg.Encode(resultMergeImg, mergeImg, &jpeg.Options{jpeg.DefaultQuality})
	defer resultMergeImg.Close()
}
