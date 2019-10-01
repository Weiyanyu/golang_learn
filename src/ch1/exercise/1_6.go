/*
练习 1.6： 修改Lissajous程序，修改其调色板来生成更丰富的颜色，然后修改SetColorIndex的第
三个参数，看看显示结果吧。
 */
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

func main55() {
	// The sequence of images is deterministic unless we seed
	// the pseudo‐random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())
	f, err := os.Create("E:\\go_project\\golang_learn\\ch1\\lissajous.gif")
	if err != nil {
		fmt.Fprintf(os.Stderr, "lissajous: %v\n", err)
		return
	}
	lissajous6(f)
}
func lissajous6(out io.Writer) {

	//调色板程序可以直接对应颜色的RGB
	var palette = []color.Color{
		color.White,
		color.Black,
		color.RGBA{0x00,0x80,0x00,0xff},
		color.RGBA{0x1e,0x90,0xff,0xff},
	}
	const (
		whiteIndex = 0 // first color in palette
		blackIndex = 1 // next color in palette
		greenIndex = 2
		DoderBlueIndex = 3
	)

	const (
		cycles = 5 // number of complete x oscillator revolutions
		res = 0.001 // angular resolution
		size = 100 // image canvas covers [‐size..+size]
		nframes = 64 // number of animation frames
		delay = 8 // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex := uint8(rand.Intn(4))
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}