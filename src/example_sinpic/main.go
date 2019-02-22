package main

import(
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main(){
	const iSize = 300

	// 创建一份灰度图
	pic := image.NewGray(image.Rect(0,0,iSize,iSize))

	for x := 0; x < iSize; x++{
		for y := 0; y < iSize; y++{
			pic.SetGray(x,y,color.Gray{255})
		}
	}

	for x := 0; x < iSize; x++{
		iSinValue := float64(x) * 2 * math.Pi / iSize
		y := iSize / 2 - math.Sin(iSinValue) * iSize / 2
		pic.SetGray(x,int(y),color.Gray{0})
	}

	file,err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}

	png.Encode(file,pic)

	file.Close()
}