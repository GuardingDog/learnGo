package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	// 创建一个新的RGBA图像对象
	width := 200
	height := 200
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))

	// 填充背景色
	bgColor := color.RGBA{255, 255, 255, 255}
	draw.Draw(rgba, rgba.Bounds(), &image.Uniform{bgColor}, image.Point{}, draw.Src)

	// 绘制笑脸
	centerX := width / 2
	centerY := height / 2
	radius := width / 4
	drawCircle(rgba, centerX, centerY, radius, color.White)

	eyeRadius := radius / 8
	drawCircle(rgba, centerX-radius/2, centerY-radius/2, eyeRadius, color.Black)
	drawCircle(rgba, centerX+radius/2, centerY-radius/2, eyeRadius, color.Black)

	mouthWidth := radius / 2
	mouthHeight := radius / 4
	mouthX := centerX - mouthWidth/2
	mouthY := centerY + radius/2 - mouthHeight/2
	drawEllipse(rgba, mouthX, mouthY, mouthWidth, mouthHeight, color.Black)

	// 保存图像到文件
	outputFile, err := os.Create("smiley_face.png")
	if err != nil {
		fmt.Println("Failed to create output file:", err)
		return
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, rgba)
	if err != nil {
		fmt.Println("Failed to encode image:", err)
		return
	}

	fmt.Println("Smiley face generated!")
}

// 绘制圆形
func drawCircle(img draw.Image, centerX, centerY, radius int, color color.Color) {
	for x := centerX - radius; x <= centerX+radius; x++ {
		for y := centerY - radius; y <= centerY+radius; y++ {
			if (x-centerX)*(x-centerX)+(y-centerY)*(y-centerY) <= radius*radius {
				img.Set(x, y, color)
			}
		}
	}
}

// 绘制椭圆
func drawEllipse(img draw.Image, x, y, width, height int, color color.Color) {
	for cx := x - width/2; cx <= x+width/2; cx++ {
		for cy := y - height/2; cy <= y+height/2; cy++ {
			if (cx-x)*(cx-x)*height*height+(cy-y)*(cy-y)*width*width <= height*height*width*width {
				img.Set(cx, cy, color)
			}
		}
	}
}
