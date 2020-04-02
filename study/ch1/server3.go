package main

import (
	"fmt"
	"log"
	"net/http"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	// http.HandleFunc("/", handler) // 所有/开头的URL 都使用handler函数处理
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		lissajous(w)
	}) // 匿名函数 func 也可以拆开写 作用域不同 
	log.Fatal(http.ListenAndServe("localhost:8000", nil))// 监听8000端口
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "%s: %s\n", k, v)
	}
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "RemotedAddr: %s\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil { // 缩小err变量的作用域
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "%s: %s\n", k, v)
	}
	// fmt.Fprintf(w, "\n")
	// lissajous(w)
}

func lissajous(out io.Writer) {
	const (
		cycles = 5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles * 2 * math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t * freq + phase)
			img.SetColorIndex(size + int(x*size + 0.5), size+int(y*size + 0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}