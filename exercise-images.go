package main

import (
  "golang.org/x/tour/pic"
  "image"
  "image/color"
)

type Image struct{
  width, height int
  color uint8
}

func (i Image) Bounds() image.Rectangle {
  return image.Rect(0, 0, i.width, i.height)
}

func (i Image) ColorModel () color.Model {
  return color.RGBAModel
}

func (i Image) At (x, y int) color.Color {
  return color.RGBA{i.color + uint8(x), i.color + uint8(y), 255, 255}
}

func main() {
  m := Image{100, 100, 100}
  pic.ShowImage(m)
}
