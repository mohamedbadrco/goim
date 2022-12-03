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
 func rpimage(dc *gg.Context, col color.Color, x, y int, k string) {
    dc.Rotate(float64(3.14159))
	
	//aR, aG, aB, aA := col.RGBA()
	//_ ,_,_= aR, aG, aB
	//col = color.RGBA{74,246 ,38, uint8(aA >>  8) }
	
    dc.SetColor(col)
    
    dc.DrawString(k, float64(-x-14),float64(-y+14))
    dc.DrawString(k, float64(-x-14),float64(-y+14))
	// dc.DrawString(k, float64(x+7),float64(y+7))
    dc.Rotate(float64(3.14159))
   // print(float64(x))
 }
