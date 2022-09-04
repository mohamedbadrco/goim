package main


import (
     "github.com/fogleman/gg"


     "image/color"

)

func pimage(dc *gg.Context, col color.Color, x, y int, k string) {

    dc.SetColor(col)
    
    dc.DrawString(k, float64(x),float64(y))
    dc.DrawString(k, float64(x),float64(y))
	// dc.DrawString(k, float64(x+7),float64(y+7))
 }