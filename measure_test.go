package gg

import (
	"log"
	"testing"
)

func TestStringWrapped(t *testing.T) {
	dc := NewContext(1000, 1000)
	fontFace, err := LoadFontFace("test_res/fonts/msyh.ttc", 38)
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
	fontFace, err := LoadFontFace("test_res/fonts/msyh.ttc", 18)
	if err != nil {
		panic(err)
	}
	log.Println(MeasureStringHeight("Hello, world! How are you?", 90, 18, 1.5, fontFace, false))
}
