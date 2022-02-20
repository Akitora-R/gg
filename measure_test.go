package gg

import (
	"log"
	"testing"
)

const fontPath = "test_res/fonts/msyh.ttc"

func TestStringWrapped(t *testing.T) {
	dc := NewContext(1000, 1000)
	fontFace, err := LoadFontFace(fontPath, 38)
	if err != nil {
		panic(err)
	}
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetFontFace(fontFace)
	dc.SetRGB(0, 0, 0)
	dc.DrawStringWrapped("Hello, world! How are you?", 500, 500, 0.5, 0.5, 90, 1.5, AlignCenter, true)
	if err = dc.SavePNG("out.png"); err != nil {
		panic(err)
	}
}

func TestWordMeasure(t *testing.T) {
	fontFace, err := LoadFontFace(fontPath, 38)
	if err != nil {
		panic(err)
	}
	log.Println(MeasureStringHeight("Hello, world! How are you?", 90, 38, 1.5, fontFace, true))
}
