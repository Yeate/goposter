package module

import (
	"fmt"

	"golang.org/x/image/font"

	"github.com/golang/freetype/truetype"

	"github.com/fogleman/gg"
)

type Text struct {
	X          int     `json:"x"`
	Y          int     `json:"y"`
	Text       string  `json:"text"`
	Width      float64 `json:"width"`
	FontSize   int     `json:"fontSize"`
	Color      string  `json:"color"`
	LineHeight int     `json:"lineHeight"`
	TextAlign  string  `json:"textAlign"`
}

const TextAlignCenter = "center"

func (t *Text) DrawX(w float64) float64 {
	if t.TextAlign == TextAlignCenter {
		return float64(t.X) - w/2
	}
	return float64(t.X)
}

func (text *Text) Draw(dc *gg.Context) error {
	var fontFace font.Face
	if text.LineHeight == 0 {
		text.LineHeight = text.FontSize
	}
	if Font != nil {
		fontFace = truetype.NewFace(Font, &truetype.Options{
			Size: float64(text.FontSize),
		})

	}
	dc.SetFontFace(fontFace)
	//dc.FontHeight()
	////_ = dc.LoadFontFace(viper.GetString("font.path"), float64(text.FontSize))
	dc.SetHexColor(text.Color)
	w, _ := dc.MeasureString(text.Text)
	words := dc.WordWrap(text.Text, text.Width)
	for index, word := range words {
		fmt.Println(float64(text.Y + text.LineHeight*(index+1)))
		dc.DrawString(word, text.DrawX(w), float64(text.Y+text.LineHeight*(index+1)-5))
	}
	return nil
}
