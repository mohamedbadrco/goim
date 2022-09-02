package main

import (
    "fmt"
     "image/png"
     "image/draw"
	"image/color"
     "image/color/palette"
    "log"
    "os"
    "github.com/nfnt/resize"
    "github.com/fogleman/gg"  
    "math/rand"
     "image/gif"
     // "encoding/binary"
    "image"
    // "io"
   // "github.com/andybons/gogif"
)

func main()  {
	
     catFile, err := os.Open("/home/mohmedbadrco/goimage/images/Dino4.png")
    if err != nil {
        log.Fatal(err)
    }
    defer catFile.Close()
 
    imData, err := png.Decode(catFile)
    if err != nil {
        fmt.Println(err)
    }
 
   
    if err != nil {
        log.Fatal(err)
    }


    
    
    
    imDatan := resize.Resize(0, 60, imData, resize.Lanczos3)

    
      
    // const int  ax = imDatan.Bounds().Max.X

    var am [][][2] int

    gscale := "$@ß%8&WM#*öõäåhkbdpqwmZÖØ0QLÇJÜÝXzçvünxrjft/|()1{}[]?-_+~<>i!lI;:,\"^`'. "

    glen := len(gscale)

    
    g := rand.Intn(20) + 10
    vis := true
    dc1 := gg.NewContext(imDatan.Bounds().Max.X * 14, imDatan.Bounds().Max.Y * 14)

    if err := dc1.LoadFontFace("/home/mohmedbadrco/goimage/font/Audiowide-Regular.ttf", 14); err != nil {
        panic(err)
    }

    dc1.SetRGB(0, 0, 0)
    dc1.Fill()
                
                for x := imDatan.Bounds().Min.X; x < imDatan.Bounds().Max.X; x++ {
                    var subam  [][2] int
                    for y := imDatan.Bounds().Min.Y; y < imDatan.Bounds().Max.Y; y++ {
            c := color.GrayModel.Convert(imDatan.At(x, y)).(color.Gray)
           
            if (g > 0){
                g--
            }

            if(g <= 0){
                vis = !(vis)
                fmt.Println(vis)
                g = rand.Intn(20) + 10
                fmt.Println(g)
            }
           // fmt.Println(int(c.Y))
            avg := int(((int(c.Y) * glen) / 255))
            if (avg > (glen - 1)){
                avg = glen - 1
            }
            // fmt.Println(avg)
         // k := gscale[avg];
         if(vis){
            if (g == 1){
                var gh = [2] int {avg,1}
                subam = append(subam,gh)
                pimage(dc1, color.RGBA{255, 255, 255, 255}, x * 14 ,14 + y *14,string(gscale[avg]))

            }else{
                pimage(dc1, imDatan.At(x, y), x * 14 ,14 + y *14,string(gscale[avg]))

                var gh = [2] int {avg,0}
                subam = append(subam,gh)
            }
            
         }else{
            var gh = [2] int {avg,2}
                subam = append(subam,gh)
         }
        
          
            
            // pimage(dc, imDatan.At(x, y), x * 14 ,14 + y *14,string(k))
        }
        am = append(am,subam)
    }
    dc1.Clip()
    
    // bbb := dc.Image()
    dc1.SavePNG("format.png")

    fmt.Println(len(am),len(am[int(len(am)/2)]))
    

    var images []*image.Paletted

    var delays []  int 

    for i := 0; i < 60; i++{
        dc := gg.NewContext(imDatan.Bounds().Max.X * 14, imDatan.Bounds().Max.Y * 14)

        if err := dc.LoadFontFace("/home/mohmedbadrco/goimage/font/Audiowide-Regular.ttf", 14); err != nil {
            panic(err)
        }

        dc.SetRGB(0, 0, 0)
        dc.Fill()
         

        for j := 0; j < len(am) - 1; j++{
        tem := am[j][0][1]
        am[j][0][1] = am[j][len(am)-1][1]
        
        fmt.Println("JJ")
            for k := len(am[0]) - 1; k > 1; k--{
                // fmt.Println(am[j][k][1],am[j][k-1][1])
               am[j][k][1] = am[j][k-1][1]
             //  fmt.Println(am[j][k])
                
            }
            am[j][1][1] = tem
        }
   


            for x := imDatan.Bounds().Min.X; x < imDatan.Bounds().Max.X; x++ { 
                for y := imDatan.Bounds().Min.Y; y < imDatan.Bounds().Max.Y; y++ {
                    
               if(am[x][y][1] == 0){
                
                pimage(dc, imDatan.At(x, y), x * 14 ,14 + y *14,string(gscale[am[x][y][0]]))
               // fmt.Println(aR, aG, aB, aA)



               }
               if(am[x][y][1] == 1){
                aR, aG, aB, aA := imDatan.At(x, y).RGBA() // no more error
                _ ,_,_= aR, aG, aB
                pimage(dc, color.RGBA{255, 255, 255, uint8(aA >>  8) }, x * 14 ,14 + y *14,string(gscale[am[x][y][0]]))
               }
                

            }
       
    }


    delays = append(delays,7)

    dc.Clip()
    
    bbb := dc.Image()
    dc.SavePNG(fmt.Sprintf("f/a%d.png", i))

    palettedImage := image.NewPaletted(image.Rect(0, 0, bbb.Bounds().Max.X, bbb.Bounds().Max.Y), palette.Plan9)
    draw.Draw(palettedImage, palettedImage.Rect, bbb,bbb.Bounds().Min, draw.Over)
     images = append(images,palettedImage)
}

f, err := os.OpenFile("rgb.gif", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})

}