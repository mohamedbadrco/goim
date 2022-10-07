package main

import (
    "fmt"
    //  "image/png"
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
   // "image/jpeg"
    
    // "io"
   // "github.com/andybons/gogif"
   "time"
   "io/ioutil"

    "github.com/golang/freetype"
    "github.com/golang/freetype/truetype"
)

// func init() {
//     // damn important or else At(), Bounds() functions will
//     // caused memory pointer error!!
//     image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
// }

func g103d6(path string)  {
    size := 60

	rand.Seed(time.Now().UnixNano())


    fontFile := "/home/mohmedbadrco/goimage/font/Audiowide-Regular.ttf"
    fontBytes, err := ioutil.ReadFile(fontFile)
    fontb, err := freetype.ParseFont(fontBytes)
    // font

    face := truetype.NewFace(fontb, &truetype.Options{Size: 14 ,DPI: 90})

    
     catFile, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    defer catFile.Close()
 
    // imData, err := png.Decode(catFile)
    // if err != nil {
	// 	// Handle error
	// }
    imData, imageType, err := image.Decode(catFile)
	if err != nil {
		// Handle error
	}
	
   // fmt.Println(imData)

   _ = imageType
    
    
    
    imDatan := resize.Resize(0, 60, imData, resize.Lanczos3)

    
      
    // const int  ax = imDatan.Bounds().Max.X

    var am [][][2] int


    

   // gscale := "$@B%8&WẄMÄÚŬÛŮÜŨ#*oöøäahkbdpqÞwmZŽO0QLCJUYŸXzcvunxrjft/|()1{}[]?-_+~<>i!lI½¼;:,\"^`'. "
   gscale := "$@B%8&WẄMÄÚŬÛŮÜŨ#*oöøäahkbdpqÞwmZŽŒO0QLCJUYŸXzcvúùŭûůüűunxrjft?-_+~<>i!lI½¼;:,\"^`'. "

    glen := len(gscale)

     rn := rand.Intn(len(gscale) - 10)

     gscale = gscale[rn:rn+rand.Intn(glen - rn )]

      glen = len(gscale)
    
    g := rand.Intn(20) + 10
    vis := true
   

   
        for x := imDatan.Bounds().Min.X; x < imDatan.Bounds().Max.X; x++ {
                    var subam  [][2] int
            for y := imDatan.Bounds().Min.Y; y < imDatan.Bounds().Max.Y; y++ {
            c := color.GrayModel.Convert(imDatan.At(x, y)).(color.Gray)
           
            if (g > 0){
                g--
            }

            if(g <= 0){
                vis = !(vis)
              //  fmt.Println(vis)
                g = rand.Intn(20) + 10
               // fmt.Println(g)
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

            }else{

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
    
    var rcolors [][3] int

	for i := 0;i<10;i++{
		var gh = [3] int {rand.Intn(255),rand.Intn(255),rand.Intn(255)}
    rcolors = append(rcolors,gh)
	}

    var images []*image.Paletted

    var delays []  int 



	var ob [][] int

	for x := imDatan.Bounds().Min.X; x < imDatan.Bounds().Max.X; x++ { 
            var  subo [] int
		for y := imDatan.Bounds().Min.Y; y < imDatan.Bounds().Max.Y; y++ {
			aR, aG, aB, aA := imDatan.At(x, y).RGBA() // no more error
		_ ,_,_= aR, aG, aB          
		o := rand.Intn(int(uint8(aA >>  8))+1)
		subo = append(subo,o)
	
	}
     ob = append(ob,subo)
}

    for i := 0; i < size; i++{
        dc := gg.NewContext(imDatan.Bounds().Max.X * 14, imDatan.Bounds().Max.Y * 14)

        dc.SetFontFace(face)

        dc.SetRGB(0, 0, 0)
        dc.Fill()
         

        for j := 0; j < len(am) - 1; j++{
        tem := am[j][0][1]
        am[j][0][1] = am[j][len(am[0])-1][1]
        
        
            for k := len(am[0]) - 1; k > 1; k--{
                // fmt.Println(am[j][k][1],am[j][k-1][1])
               am[j][k][1] = am[j][k-1][1]
              // fmt.Println(am[j][k])
                
            }
            am[j][1][1] = tem
        }

        offrand1 := 0
        offrand2 := 0
        // if(i == 28)  {
        //     v := 0
        //     if rand.Intn(2) > 0 {
        //         v = 1
        
        //         v = -1
        // }
        //     offrand1 = int(float64(imDatan.Bounds().Max.X)*0.35) * v
            
        //     if rand.Intn(2) > 0 {
        //         v = 1
        // } else {
        //         v = -1
        // }

        //     offrand2 = int(float64(imDatan.Bounds().Max.X)*0.35) * v }

            
          

        //    if(i == 31)  {
        //     v := 0
        //     if rand.Intn(1) > 0 {
        //         v = 1
        // } else {
        //         v = -1
        // }
        //     offrand1 = int(float64(imDatan.Bounds().Max.X)*0.35) * v
            
        //     if rand.Intn(1) > 0 {
        //         v = 1
        // } else {
        //         v = -1
        // }

        //     offrand2 = int(float64(imDatan.Bounds().Max.X)*0.35) * v }

        //     if(i == 31 || i > 32){
        //         offrand1 = 1
        //         offrand2 = 1
        //     }

   


            for x := imDatan.Bounds().Min.X; x < imDatan.Bounds().Max.X; x++ { 
              
                for y := imDatan.Bounds().Min.Y; y < imDatan.Bounds().Max.Y; y++ {
                    aR, aG, aB, aA := imDatan.At(x, y).RGBA() // no more error
                _ ,_,_= aR, aG, aB
                index := int((float64(am[x][y][0])/float64(glen))*float64(len(rcolors)))
				rc  := rcolors[index]
                // fmt.Println(index)
				
               if(am[x][y][1] == 0){
				// fmt.Println()
                
                pimage(dc, color.RGBA{uint8( rc[0]),uint8(rc[1]), uint8(rc[2]), uint8(rand.Intn(int(uint8(aA >>  8))+1)) }, x * 14 ,14 + y *14,string(gscale[am[x][y][0]]))
              pimage(dc, color.RGBA{uint8( rc[0]),uint8(rc[1]), uint8(rc[2]), uint8(rand.Intn(int(((aA/8 >>  8)))+1)) }, x * 14 + 7 ,21 + y *14,string(gscale[am[x][y][0]]))
				


               }
               if(am[x][y][1] == 1){
                
                pimage(dc, color.RGBA{255, 255, 255, uint8(rand.Intn(int(uint8(aA >>  8))+1)) }, x * 14 ,14 + y *14,string(gscale[am[x][y][0]]))
                 pimage(dc, color.RGBA{255, 255, 255, uint8(rand.Intn(int(((aA/8 >>  8)))+1)) }, x * 14 + 7 ,21 + y *14,string(gscale[am[x][y][0]]))

               } 

            }
			   
			


                

            
       
    }

  _ , _ = offrand1,offrand2
    delays = append(delays,7)

    dc.Clip()
    
    bbb := dc.Image()

    palettedImage := image.NewPaletted(image.Rect(0, 0, bbb.Bounds().Max.X, bbb.Bounds().Max.Y), palette.Plan9)
    draw.Draw(palettedImage, palettedImage.Rect, bbb,bbb.Bounds().Min, draw.Over)
     images = append(images,palettedImage)
	   // dc.SavePNG(fmt.Sprintf("f/a%d.png", i))

}
  now := time.Now()
    name := fmt.Sprintf("out/%srbg.gif",now)
f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0600)
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