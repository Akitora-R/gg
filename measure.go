package gg

import (
	"golang.org/x/image/font"
)

type strCtx struct {
	fontHeight float64
	fontFace   font.Face
}

func (sc strCtx) MeasureString(s string) (w, h float64) {
	d := &font.Drawer{
		Face: sc.fontFace,
	}
	a := d.MeasureString(s)
	return float64(a >> 6), sc.fontHeight
}

func MeasureStringHeight(s string, width, fontHeight, lineSpacing float64, fontFace font.Face, breakWord bool) float64 {
	ctx := strCtx{
		fontHeight: fontHeight,
		fontFace:   fontFace,
	}
	result := wordWrap(ctx, s, width, breakWord)
	return float64(len(result)) * fontHeight * lineSpacing
}
