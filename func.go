package main


import (
     "github.com/fogleman/gg"

     "math/rand"

     "image/color"
    // "fmt"

)

func pimage(dc *gg.Context, col color.Color, x, y int, k string) {

    dc.SetColor(col)
    
    dc.DrawString(k, float64(x),float64(y))
    dc.DrawString(k, float64(x),float64(y))
	// dc.DrawString(k, float64(x+7),float64(y+7))
 }
 func rpimage(dc *gg.Context, col color.Color, x, y int, k string) {
    ro := 3.14159
    if(rand.Intn(10) > 5){

    ro = ro*0.5
   }
	dc.Rotate(float64(ro))
	//aR, aG, aB, aA := col.RGBA()
	//_ ,_,_= aR, aG, aB
	//col = color.RGBA{74,246 ,38, uint8(aA >>  8) }
	
    dc.SetColor(col)
    
    dc.DrawString(k, float64(-x-14),float64(-y+14))
    dc.DrawString(k, float64(-x-14),float64(-y+14))
	// dc.DrawString(k, float64(x+7),float64(y+7))
    dc.Rotate(float64(((3.14159)*2) - ro))
   // print(float64(x))
 }


 func isAvailable(alpha [][2]int, num [2]int) bool {

    // iterate using the for loop
    for i := 0;
    i < len(alpha);
    i++ {
       // check      
       if alpha[i][0] == num[0] && alpha[i][1] == num[1] {
          // return true
          return true
       }
    }
    return false
 }



 func isAvailable2(alpha [][][2]int, num [][2]int) bool {

   // iterate using the for loop
   for k := 0;
   k < len(alpha);
   k++ {
   for i := 0;
   i < len(alpha[k]);
   i++ {
      // check   
      for j := 0;
   j < len(num);
   j++ {   
      if alpha[k][i][0] == num[j][0] && alpha[k][i][1] == num[j][1] {
         // return true
         return true
      }}}
   }
   return false
}




// HLine draws a horizontal line
// func HLine(dc *gg.Context,x1, y, x2 int) {
//    for ; x1 <= x2; x1++ {
//        dc.Set(x1, y, col)
//    }
// }

// // VLine draws a veritcal line
// func VLine(dc *gg.Context,x, y1, y2 int) {
//    for ; y1 <= y2; y1++ {
//        dc.Set(x, y1, col)
//    }
// }

// Rect draws a rectangle utilizing HLine() and VLine()
func Rect(dc *gg.Context,x, y, w, h float64,col color.Color,t int) {
   // HLine(x1, y1, x2)
   // HLine(x1, y2, x2)
   // VLine(x1, y1, y2)
   // VLine(x2, y1, y2)
   dc.SetColor(col)
   for i:= 0;i < t ;i++{
   dc.DrawRectangle(x,y,w,h)
   x -= 1
   y -= 1
   h -= 1
   w -= 1
  // fmt.Println("h")
}
}

func mapimag(dc *gg.Context){

    //     dc := gg.NewContext( 2 * 30, 2 * 30)
    //   // Rect(dc, float64(1),float64(1),float64(10),float64(10),color.RGBA{255, 255,255, 255 },5)
    //       //  dc.SetColor(color.RGBA{255, 255,255, 255 })
    //       //  dc.SetLineWidth(float64(5))
    //       //  dc.NewSolidPattern(color.RGBA{255, 255,255, 255 })

    //        // dc.Push()
    //   //  for i:= 0;i < 5 ;i++{
    //      dc.DrawLine(float64(10),float64(10),float64(30),float64(30))
    //      dc.SetRGB(255, 0, 0)
    //      dc.Fill()
        // dc.DrawString(string(byte(i)), float64(10),float64(10))
      //  dc.Fill()
		//dc.Pop()

        // fmt.Println("h")
//}
//dc := gg.NewContext(880, 880)
dc.SetLineWidth(float64(2))

dc.DrawRectangle(5, 5, 880 - 10 , 880 - 10)
dc.DrawRectangle(5, 5, 880 - 10 , 46)
dc.DrawRectangle(880 - 150, 5, 150 - 5 , 46)

// dc.SetRGB(1, 0, 0)
dc.SetColor(color.RGBA{255, 255,255, 255 })

dc.Stroke()

// dc.DrawRectangle(7, 7, 990 -2, 990 -2)
// dc.SetRGB(0, 0, 0)
// dc.Fill()

//dc.SavePNG(fmt.Sprintf("out/%s %d.png", string(byte(i)),i))
}
