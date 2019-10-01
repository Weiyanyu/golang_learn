/*
练习 1.12： 修改Lissajour服务，从URL读取变量，比如你可以访问 http://localhost:8000/?
cycles=20 这个URL，这样访问可以将程序里的cycles默认的5修改为20。字符串转换为数字可以
调用strconv.Atoi函数。你可以在godoc里查看strconv.Atoi的详细说明。
*/
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/lissajous", func(writer http.ResponseWriter, request *http.Request) {
		lissajousService(writer, request)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajousService(out io.Writer, request *http.Request) {

	//调色板程序可以直接对应颜色的RGB
	var palette = []color.Color{color.White, color.RGBA{0x00, 0x80, 0x00, 0xff}}
	const (
		whiteIndex = 0 // first color in palette
		greenIndex = 1 // next color in palette
	)
	if err := request.ParseForm(); err != nil {
		log.Fatal(err)
	}



	cycles := float64(conveter(request.Form.Get("cycles"), 5))   // number of complete x oscillator revolutions
	res := 0.001  // angular resolution
	size := conveter(request.Form.Get("size"), 100)  // image canvas covers [‐size..+size]
	nframes := conveter(request.Form.Get("nframes"), 64) // number of animation frames
	delay := conveter(request.Form.Get("delay"), 8)    // delay between frames in 10ms units

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func conveter(str string, defautValue int) int {
	res, err := strconv.Atoi(str)
	if err != nil {
		res = defautValue
	}
	return res
}
