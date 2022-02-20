package gg

import (
	"golang.org/x/image/font"
	"strings"
)

func MeasureStringHeight(s string, fontWidth, fontPoints, lineSpacing float64, fontFace font.Face, breakWord bool) float64 {
	d := &font.Drawer{
		Face: fontFace,
	}
	fontHeight := fontPoints * 72 / 96
	var result []string
	for _, line := range strings.Split(s, "\n") {
		fields := splitString(line, breakWord)
		if len(fields)%2 == 1 {
			fields = append(fields, "")
		}
		x := ""
		for i := 0; i < len(fields); i += 2 {
			a := d.MeasureString(x + fields[i])
			w := float64(a >> 6)
			if w > fontWidth {
				if x == "" {
					result = append(result, fields[i])
					x = ""
					continue
				} else {
					result = append(result, x)
					x = ""
				}
			}
			x += fields[i] + fields[i+1]
		}
		if x != "" {
			result = append(result, x)
		}
	}
	for i, line := range result {
		result[i] = strings.TrimSpace(line)
	}
	return float64(len(result)) * fontHeight * lineSpacing
}
