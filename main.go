package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/smf8/xmansay/model"
	"github.com/smf8/xmansay/ui"
	"github.com/smf8/xmansay/util"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	var xManImg, xManFont string
	var xManSaySize int
	var duration int
	// check if res folder is available for default resources
	f, _ := ioutil.ReadDir("./")
	bo := false
	for _, dir := range f {
		if dir.Name() == "res" {
			bo = true
		}
	}
	if !bo {
		fmt.Println("Please obtain res folder and place it in root folder beside", os.Args[0])
		os.Exit(1)
	}
	flag.StringVar(&xManImg, "image", "", "/path/to/image (image must have a transparent background) default is poop")
	flag.StringVar(&xManFont, "font", "", "/path/to/ttf-font default is Roboto & IranSans5")
	flag.IntVar(&xManSaySize, "size", 0, "font size, default value is 50, recommended value is depended on screen resolution")
	flag.IntVar(&duration, "time", 5, "the duration which the man stays, default is 5, don't use more than 15")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(flag.CommandLine.Output(), "\\twrite XManSay text wrapped in \" \" after all flags\n")
	}
	flag.Parse()
	xManSay := flag.Args()
	if len(xManSay) != 1 {
		fmt.Println("Incorrect usage, for usage do ./main -h")
		os.Exit(1)
	}
	if duration > 15 {
		fmt.Println("Duration must be lower than 15. For usage check", os.Args[0], "-h")
		os.Exit(1)
	}

	//handling Arabic / Persian text
	var man *model.Xmansay
	if util.CheckIsEnglish(xManSay[0]) {
		// text is Not Persian
		man = model.NewXManSay(xManSay[0], xManImg, xManFont, float64(xManSaySize))
	} else {
		man = model.NewXManSay(util.Reverse(util.ToGlyph(xManSay[0])), xManImg, xManFont, float64(xManSaySize))
	}
	c, err := man.DrawMan()
	if err != nil {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	err = png.Encode(buf, c.Image())
	if err != nil {
		log.Fatal(err)
	}
	//c.SavePNG("output.png")
	b := buf.Bytes()
	go func() {
		time.Sleep(time.Duration(int(time.Second) * duration))
		os.Exit(0)
	}()
	ui.Display(b)
}
