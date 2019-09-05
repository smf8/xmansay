package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"math/rand"
	"os"
	"time"
)

var (
	displayArea   *widgets.QWidget
	scene         *widgets.QGraphicsScene
	view          *widgets.QGraphicsView
	item          *widgets.QGraphicsPixmapItem
	mainApp       *widgets.QApplication
	imageFileName string
	b             []byte
)

func imageViewer() *widgets.QWidget {
	displayArea = widgets.NewQWidget(nil, 0)
	scene = widgets.NewQGraphicsScene(nil)
	view = widgets.NewQGraphicsView(nil)

	img := gui.QImage_FromData(b, len(b), "PNG")

	var pixmap = gui.QPixmap_FromImage(img, core.Qt__AutoColor)
	item = widgets.NewQGraphicsPixmapItem2(pixmap, nil)
	scene.AddItem(item)

	bounding := scene.ItemsBoundingRect() // to get the rect around all items in scene
	view.SetFixedWidth(int(bounding.Width()) + 20)
	view.SetFixedHeight(int(bounding.Height()) + 20)
	view.SetScene(scene)
	view.SetStyleSheet("background: transparent")

	var layout = widgets.NewQVBoxLayout()

	layout.AddWidget(view, 1, core.Qt__AlignCenter)
	displayArea.SetLayout(layout)
	displayArea.SetStyleSheet("background: transparent")
	return displayArea
}

//Display uses
func Display(imageData []byte) {
	b = imageData
	imageFileName = "output.png"
	mainApp = widgets.NewQApplication(len(os.Args), os.Args)
	window := widgets.NewQMainWindow(nil, 0)
	window.SetAttribute(core.Qt__WA_NoSystemBackground, true)
	window.SetAttribute(core.Qt__WA_TranslucentBackground, true)
	window.SetWindowFlag(core.Qt__Window, true)
	window.SetWindowFlag(core.Qt__FramelessWindowHint, true)
	widget := imageViewer()
	validX := mainApp.Screens()[0].Geometry().Width() - widget.Width()
	validY := mainApp.Screens()[0].Geometry().Height() - widget.Height()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	validX = r.Intn(validX)
	validY = r.Intn(validY)
	window.SetCentralWidget(widget)
	window.SetGeometry2(validX, validY, widget.Width(), widget.Height())
	window.Show()
	mainApp.Exec()
}
