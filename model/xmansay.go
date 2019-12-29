package model

import (
	"github.com/fogleman/gg"
	"strings"
)

const defaultSize = 40

//Xmansay struct holds an image and a font and text to represent a message
type Xmansay struct {
	text  string
	image string
	font  string
	size  float64
}

//NewXManSay creates a new xmansay instance
//defalut values are : image: poop, font: Roboto, size: 50
func NewXManSay(text, image, font string, size float64) *Xmansay {

	if size <= 0 {
		size = defaultSize
	}
	return &Xmansay{text, image, font, size}
}

//DrawMan clearly draws an xman Context, with x prefrences
func (x *Xmansay) DrawMan() (*gg.Context, error) {
	img, err := gg.LoadPNG(x.image)
	if err != nil {
		return nil, err
	}
	imgWidth := float64(img.Bounds().Dx())
	imgHeight := float64(img.Bounds().Dy())
	//init Context
	c := gg.NewContext(int(imgWidth*1.7), int(imgHeight*1.5))
	err = c.LoadFontFace(x.font, x.size)
	if err != nil {
		return nil, err
	}
	//Draw image to context
	c.DrawImage(img, 0, int(0.5*imgHeight))
	//Prepare the text
	wrappedText := c.WordWrap(x.text, 0.5*imgWidth)
	x.text = strings.Join(wrappedText, "\n")
	_, textHeight := c.MeasureMultilineString(x.text, 1.4)
	// if textHeight is greater than 0.5 * image height then some of the text is gonna overlap
	// to avoid it we move the dialouge down
	var dialougePosY float64
	dialougePosY = 30
	if textHeight > 0.5*imgHeight {
		dialougePosY += textHeight - (0.5 * imgHeight)
	}
	x.text = strings.ReplaceAll(x.text, "\n", " ")
	//Draw Dialouge box
	c.DrawRoundedRectangle(1.1*imgWidth, dialougePosY, 0.5*imgWidth+20, textHeight+20, 5)
	c.FillPreserve()
	c.SetLineWidth(2)
	c.Stroke()
	//Write Text into Dialouge
	c.SetRGB(0, 0, 0)
	c.DrawStringWrapped(x.text, 1.1*imgWidth+10, dialougePosY, 0, 0, 0.5*imgWidth, 1.4, gg.AlignCenter)
	return c, err
}
