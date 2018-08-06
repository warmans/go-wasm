package main

import (
	"github.com/dennwc/dom"
	"image/color/palette"
	"math/rand"
	"time"
	"github.com/dennwc/dom/svg"
	"fmt"
)

func main() {

	doc := dom.GetDocument()
	container := doc.QuerySelector("#output")

	img := svg.New(dom.Px(800), dom.Px(600))

	//clear loading message and add image
	container.SetInnerHTML("")
	container.AppendChild(img.DOMElement())

	rect := img.NewRect(100, 100)
	wow := img.NewText("WOW")

	for {
		rect.Translate(float64(rand.Intn(800-100)), float64(rand.Intn(600-100)))
		rect.Fill(dom.Color(rgba2string(palette.WebSafe[rand.Intn(len(palette.WebSafe))].RGBA())))
		wow.Translate(float64(rand.Intn(800-100)), float64(rand.Intn(600-100)))

		time.Sleep(time.Millisecond * 100)
	}
}

func rgba2string(r, g, b, a uint32) string {
	return fmt.Sprintf("rgba(%d, %d, %d, %d)", r, g, b, a)
}
