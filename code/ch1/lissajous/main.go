package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var palette = []color.Color{color.RGBA{194, 249, 132, 230}, color.Black, color.Gray{4}, color.RGBA{44,250,135,123},color.RGBA{0,0,0,123},color.RGBA{0,0,126,0},color.RGBA{0,250,0,0}}

var cycles map[string]float64     // number of complete x oscillator revolutions

func main() {
	cycles = make(map[string]float64)

	cycles["initial"] = 5

	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		setfunc := func(w http.ResponseWriter, r *http.Request){

			if strings.HasPrefix(r.URL.Path, "/cycles="){
				string_cycle := r.URL.Path
				string_cycle = strings.TrimPrefix(string_cycle,"/cycles=")
				setting_to_20,err := strconv.Atoi(string_cycle)
				if err != nil {
					panic("not a good string")
				}
				cycles["initial"] = float64(setting_to_20)
			}
			lissajous(w)
			fmt.Println(cycles)
		}
		
		fmt.Println(cycles)
		
		http.HandleFunc("/", handler)
		http.HandleFunc("/cycles=20", setfunc)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	lissajous(os.Stdout)
}



func lissajous(out io.Writer) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles["initial"]*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(i))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
