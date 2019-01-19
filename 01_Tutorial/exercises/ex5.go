// Ex 1.5 - Change the Lissajous program's color palette to gren on black, for
// added authenticity
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

// `[]color.Color{...}` is a composite literal -- it instantiates a composite
// type from a sequence of element values. This one is a slice.
var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}

// const declarations can appear at the package level or at the function level
const (
	blackIndex = 0 // first color in palette
	greenIndex = 1 // next color in palette
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	// Why aren't these constants declared at the package level like whiteIndex
	// and blackIndex?
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0        // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes} // Another composite literal, this time a struct
	phase := 0.0                        // phase difference

	// Each iteration in this outer loop is a single frame of the animation
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		// Each pass through this inner loop sets some pixels to black to generate a new image
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: Ignoring encoding errors
}
