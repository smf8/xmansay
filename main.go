package main

import (
	"bytes"
	"github.com/smf8/xmansay/model"
	"image/png"
	"log"
)

func main() {
	man := model.NewXManSay("Lorem ipsum dolor", "res/poop.png", "res/chilanka.ttf", 60)
	c, err := man.DrawMan()
	if err != nil {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	err = png.Encode(buf, c.Image())
	if err != nil {
		log.Fatal(err)
	}
	c.SavePNG("output.png")

}
