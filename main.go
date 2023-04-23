package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/fzipp/canvas"
)

func main() {
	http := flag.String("http", ":8080", "HTTP service address (e.g., '127.0.0.1:8080' or just ':8080')")
	flag.Parse()

	fmt.Println("Listening on " + httpLink(*http))
	err := canvas.ListenAndServe(*http, run,
		canvas.Size(150, 150),
		canvas.Title("Clock"),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func run(ctx *canvas.Context) {
	for {
		select {
		case event := <-ctx.Events():
			if _, ok := event.(canvas.CloseEvent); ok {
				return
			}
		default:
			draw(ctx)
			ctx.Flush()
			time.Sleep(time.Second / 2)
		}
	}
}

const WIDTH int = 1000
const HEIGHT int = 1000

func draw(ctx *canvas.Context) {
	rgba := image.NewRGBA(image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: WIDTH, Y: HEIGHT}})
	for i := 0; i < WIDTH; i++ {
		for j := 0; j < HEIGHT; j++ {
			rgba.Set(i, j, color.RGBA{R: uint8(rand.Int()), G: uint8(rand.Int()), B: uint8(rand.Int()), A: 255})
		}
	}

	imageData := ctx.CreateImageData(rgba)
	ctx.PutImageData(imageData, 0, 0)

}

func httpLink(addr string) string {
	if addr[0] == ':' {
		addr = "localhost" + addr
	}
	return "http://" + addr
}
