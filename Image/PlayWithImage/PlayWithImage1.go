package main

//https://github.com/golang-samples/image/blob/master/

import (
	//	"code.google.com/p/graphics-go/graphics"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"os/exec"
)

var (
	white color.Color = color.RGBA{255, 255, 255, 255}
	black color.Color = color.RGBA{0, 0, 0, 255}
	blue  color.Color = color.RGBA{0, 0, 255, 255}
)

// ref) http://golang.org/doc/articles/image_draw.html
func main() {

	m := image.NewRGBA(image.Rect(0, 0, 640, 480)) //*NRGBA (image.Image interface)

	// fill m in blue
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	// draw a line
	//	for i := m.Bounds().Min.X; i < m.Bounds().Max.X; i++ {
	//		m.Set(i, m.Bounds().Max.Y/2, white) // to change a single pixel
	//	}

	n := image.NewRGBA(image.Rect(100, 100, 500, 400))
	draw.Draw(n, n.Bounds(), m, image.Point{0, 0}, draw.Src)

	for y := n.Rect.Min.Y; y < n.Rect.Max.Y; y++ {
		for x := n.Rect.Min.X; x < n.Rect.Max.X; x++ {
			m.Set(x, y, color.RGBA{0x90, 0x90, 0x90, 0xff})
		}
	}
	//	//graphics.Rotate(n, img2, &graphics.RotateOptions{3.5})

	w, _ := os.Create("new.png")
	defer w.Close()
	png.Encode(w, m) //Encode writes the Image m to w in PNG format.

	//	Show(w.Name())

}

// show  a specified file by Preview.app for OS X(darwin)
func Show(name string) {
	command := "open"
	arg1 := "-a"
	arg2 := "/Applications/Preview.app"
	cmd := exec.Command(command, arg1, arg2, name)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
