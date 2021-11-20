package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"math/rand"
	"os"
	"time"
)

func main() {
	const width = 1920
	const height = 1080
	const cf = 20 // 120

	var opt gif.Options
	var outGif gif.GIF
	for i := 0; i < 10; i++ {
		img := whiteNoise(width, height, cf)

		var buf bytes.Buffer
		err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: 100})
		if err != nil {
			fmt.Println(err)
			return
		}
		var buf2 bytes.Buffer
		// Write img2gif file to buffer.
		err = gif.Encode(&buf2, img, &opt)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Decode img2gif file from buffer to img.
		inImg, err := gif.Decode(&buf2)
		if err != nil {
			fmt.Println(err)
			return
		}

		outGif.Image = append(outGif.Image, inImg.(*image.Paletted))
		outGif.Delay = append(outGif.Delay, 10)
	}

	f, err := os.OpenFile("out.gif", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	err = gif.EncodeAll(f, &outGif)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func whiteNoise(width, height, cf int) *image.Gray {
	img := image.NewGray(image.Rect(0, 0, width, height))
	for i := 0; i < width/cf; i++ {
		for j := 0; j < height/cf; j++ {
			rand.Seed(time.Now().UnixNano())
			rect := image.Rect(i*cf, j*cf, (i+1)*cf, (j+1)*cf)
			c := color.Gray{uint8(rand.Intn(255))}
			draw.Draw(img, rect, &image.Uniform{c}, image.Point{}, draw.Src)
		}
	}
	return img
}
