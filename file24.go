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

func g103d8fm(path string)  {
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
  // gscale := "$@B%8&WMÄZO0QLCJUYX#*oahkbdpqwmzcvnxrjft?-_+~<>i!lI½¼;:,\"^`'. "

   // glen := len(gscale)

   //  rn := rand.Intn(len(gscale) - 10)

    //  gscale = gscale[rn:rn+rand.Intn(glen - rn )]
    tmpgs := "Äűز$@&ظu{}[]<>Þ½¼;:,\"^`'.ݝ"
	length := 10 + rand.Intn(41)
  
    ran_str := make([]byte, length)
	offset := 0
     noe := 3
    // Generating Random string
    for i := 0; i < length; i++ {
	    randn := rand.Intn(256)
		
		if((randn >= 127 && randn <= 159) || (randn >= 1 && randn <= 31) ){
		offset += 1
		}
		
		if(offset > noe){
		
		for (randn >= 127 && randn <= 159) || (randn >= 1 && randn <= 31){
		randn = rand.Intn(256)
		
		}
		
		ran_str[i] = byte(randn)
		
		} else{
        ran_str[i] = byte(randn) }
    }
  
    // Displaying the random string
    ir := rand.Intn(10)

    for i := 0 ; i< ir;i++{
        ran_str[rand.Intn(len(ran_str))] = (tmpgs[rand.Intn(len(tmpgs))])
    }

    str := string(ran_str)
    fmt.Println(str)
	gscale := str
	// if(rand.Intn(3) >= 2 ){
    //   gscale += " " }
      gscale += " " 
	// gscale = gscale[rn:]
    


      glen := len(gscale)
    
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
        
          
            
            // pimage(dc, imDatan.At(x, y), x * 14 ,14 + (y+3)*14,string(k))
        }
        am = append(am,subam)
    }
    
    var rcolors [][3] int
	
	numc := 10 + rand.Intn(10)

     randa := rand.Intn(5)
     _ = randa
	for i := 0;i<numc;i++{
        // if(randa < 2){
        //     var gh = [3] int {rand.Intn(255),rand.Intn(155),rand.Intn(155)}
        //     rcolors = append(rcolors,gh)
        // }
        // if(randa < 4 && randa > 2){
        //     var gh = [3] int {rand.Intn(155),rand.Intn(155),rand.Intn(255)}
        //     rcolors = append(rcolors,gh)
        // }
        // if(randa < 6 && randa > 4){
        //     var gh = [3] int {rand.Intn(155),rand.Intn(255),rand.Intn(155)}
        //     rcolors = append(rcolors,gh)
        // }
        // if(randa < 8 && randa > 6){
        //     var gh = [3] int {rand.Intn(155),rand.Intn(155),rand.Intn(155)}
        //     rcolors = append(rcolors,gh)
        // }
        // if(randa < 11 && randa > 8){
        //     var gh = [3] int {rand.Intn(255),rand.Intn(255),rand.Intn(255)}
        //     rcolors = append(rcolors,gh)
        // }else{
            // var gh = [3] int {rand.Intn(255),rand.Intn(255),rand.Intn(255)}
            // rcolors = append(rcolors,gh)
        //}
    
        // if(rand.Intn(10) > 4){
        // //     tmpc := rand.Intn(155)
        // //     var gh = [3] int {tmpc,tmpc,tmpc}
        // //   //  gh[randa] = gh[randa] // + rand.Intn(255 - tmpc)
        // //     rcolors = append(rcolors,gh)
        // var gh = [3] int {rand.Intn(155),rand.Intn(155),rand.Intn(155)}
        //  rcolors = append(rcolors,gh)  
        // }else{
        // var gh = [3] int {rand.Intn(155) + 100 ,rand.Intn(155) + 100 ,
        //     rand.Intn(155) + 100}
        //  rcolors = append(rcolors,gh)  
        // }
        if(randa == 0){
                  var gh = [3] int {rand.Intn(155),rand.Intn(155),rand.Intn(255)}
            rcolors = append(rcolors,gh)

        }
        if(randa == 1){
            var gh = [3] int {rand.Intn(255),rand.Intn(155),rand.Intn(155)}
        rcolors = append(rcolors,gh)

        }
        if(randa == 2){
            var gh = [3] int {rand.Intn(155),rand.Intn(255),rand.Intn(155)}
        rcolors = append(rcolors,gh)

        }
        if(randa == 3){
            var gh = [3] int {rand.Intn(255),rand.Intn(255),rand.Intn(255)}
        rcolors = append(rcolors,gh)

        }
        if(randa == 4){
            var gh = [3] int {rand.Intn(255),rand.Intn(155),rand.Intn(255)}
        rcolors = append(rcolors,gh)

        }
	}
    numaA := 0
    for x := imDatan.Bounds().Min.X; x < imDatan.Bounds().Max.X; x++ {
for y := imDatan.Bounds().Min.Y; y < imDatan.Bounds().Max.Y; y++ {
    aR, aG, aB, aA := imDatan.At(x, y).RGBA() // no more error
    _ ,_,_= aR, aG, aB
    if(uint8(aA  >>  8) > 0){
     numaA += 1
    }
}}
fmt.Println(numaA)
////////////////////////////////////////
////////////////////////////////////////
////////////////////////////////////////
////////////////////////////////////////

cy := (((numaA- 60)/60)) / 2
fmt.Println(cy)

// declare all array of linked strings
var st [][][2] int
// declare intital dircations 
xd := true
up := true
/* cross switch  horitanzal to virtcal */
cs := true
var sti= [1] int {60}



for j := 0; j <  cy ; j++{
    var  sst [][2] int
    jk := sti[rand.Intn(len(sti))] 
    var start [2] int
    if(xd == true){
    start = [2] int {rand.Intn(60),rand.Intn(60)} 
    }else{
    start = [2] int {rand.Intn(60),rand.Intn(60)}
    }
  





    aR, aG, aB, aA := imDatan.At(start[0], start[1]).RGBA() // no more error
    _, _ ,_,_= aR, aG, aB , aA
    for((uint8(aA  >>  8) <= 0)){
        start[0] =   rand.Intn(60)
        start[1] =   rand.Intn(60)
    }
    sst = append(sst,start)



    
    ofind := 1
    for k := 1; k < jk;k++{
    var block = [2] int {0,0}
    
    
    
    offst := 0
    if(cs){
    if(rand.Intn(30) < 2){
    xd = !xd
    cs = !cs
    }
    }else{
    if(rand.Intn(20) < 2){
    up = !(up)
    cs = !cs
    } }
    
    if(up){
    offst = 1 * ofind
    }  else{
    offst = -1 * ofind
    }    
    // fmt.Println(offst)
    
    if(xd){
    block[0] = sst[len(sst) - 1][0] + offst
    block[1] = sst[len(sst) - 1][1] 
    
    if(block[0] < 0){
    block[0] = 60 + block[0]
    }
    if(block[0] >= 60){
    block[0] = 0
    }
    
    }else{
    block[0] = sst[len(sst) - 1][0] 
    block[1] = sst[len(sst) - 1][1] + offst
    
    if(block[1] < 0){
    block[1] = 60 + block[1]
    }
    if(block[1] >= 60){
    block[1] = 0
    }} 
    
    ///////////////////////////////////////////////////////////////////////////////


    aR, aG, aB, aA := imDatan.At(block[0], block[1]).RGBA() // no more error
    _, _ ,_,_= aR, aG, aB , aA
    // if(uint8(aA  >>  8) > 0){
    // if(isAvailable(sst,block)){
    //     k--
    // }else{
    //     sst = append(sst,block)}

    // // else{
    // //     if(jk< (imDatan.Bounds().Max.X * imDatan.Bounds().Max.Y)){
    // //         jk++
    // //        // fmt.Println(k,jk)
    // //     }
    // // }
    
    // }else{
        
        for (isAvailable(sst,block) || (uint8(aA  >>  8) <= 0)){
           // block = [2] int {rand.Intn(60),rand.Intn(60)}
            block[0] =   rand.Intn(60)
            block[1] =   rand.Intn(60)
            aR, aG, aB, aA = imDatan.At(block[0], block[1]).RGBA() // no more error
            _, _ ,_,_= aR, aG, aB , aA
            //fmt.Println("HI")
            // fmt.Println(block)
            // fmt.Println(isAvailable(sst,block) || (uint8(aA  >>  8) <= 0))
            // fmt.Println(j,k)
            // if(!(isAvailable(sst,block) || (uint8(aA  >>  8) <= 0))){
            //     break
            // }
        }
        sst = append(sst,block)
       // fmt.Println("append")ac
    // }
    
    }
    fmt.Println("out K ")
     //fmt.Println(sst)
    if(isAvailable2(st,sst)){
        j--
        fmt.Println("j++")

    }else{
        st = append(st,sst)
       // fmt.Println(len(sst))
       fmt.Println("append2")
    
    }
    
    fmt.Println(j)}
    
      //fmt.Println(st)
     
    

for j := 0; j <  cy ; j++{
    var  sst [][2] int

    
    for x := imDatan.Bounds().Min.X; x < imDatan.Bounds().Max.X; x++ {
        for y := imDatan.Bounds().Min.Y; y < imDatan.Bounds().Max.Y; y++ { 
            var  point = [2] int {rand.Intn(60),rand.Intn(60)} 
            aR, aG, aB, aA := imDatan.At(x, y).RGBA() // no more error
            _, _ ,_,_= aR, aG, aB , aA
            if(uint8(aA  >>  8) > 0){

            if !isAvailable3(st,point) {
                sst = append(sst,point)
            }
        }
            if(len(sst) >= 60){
                break
            }
        }
        if(len(sst) >= 60){
            //  fmt.Println(sst)
            break
            
        }
    }
        st = append(st,sst)
}
fmt.Println("m Done")
//fmt.Println(st)

//ratioo := 0
////////////////////////////////
////////////////////////////////
////////////////////////////////
////////////////////////////////
////////////////////////////////
    ra := rcolors[rand.Intn(len(rcolors))]
      _ = ra
    var images []*image.Paletted

    var delays []  int 
    o := 0
    em := lom[rand.Intn(len(lom)-1)]
        emra := rcolors[rand.Intn(len(rcolors)-1)]
    for i := 0; i < size; i++{
        dc := gg.NewContext(imDatan.Bounds().Max.X * 14 + 40 , (imDatan.Bounds().Max.Y)* 14 + 40)

        dc.SetFontFace(face)

        dc.SetRGB(0, 0, 0)
        dc.Fill()
        if(i % 5 == 0){
        em = lom[rand.Intn(len(lom))]
        emra = rcolors[rand.Intn(len(rcolors))]
    }


        // for j := 0; j < len(am) - 1; j++{
        // tem := am[j][0][1]
        // am[j][0][1] = am[j][len(am[0])-1][1]
        
        
        //     for k := len(am[0]) - 1; k > 1; k--{
        //         // fmt.Println(am[j][k][1],am[j][k-1][1])
        //        am[j][k][1] = am[j][k-1][1]
        //       // fmt.Println(am[j][k])
                
        //     }
        //     am[j][1][1] = tem
        // }

        var xo []  int 
        for j := 0; j < size; j++{
            vo := []int{2,4,-3,1,3,-4,-7,7,14,-14}
           
            if(rand.Intn(1000) < 12){
               xo =  append(xo,14*vo[rand.Intn(6)])
            }else{
                xo =  append(xo,0)
            }

        }

        for j := 0; j < len(st); j++{
            // stroe first element in tem 
            tem := am[st[j][0][0]][st[j][0][1]][1]
           // tema:= aa[st[j][0][0]][st[j][0][1]]
            
    
            am[st[j][0][0]][st[j][0][1]][1] = am[st[j][len(st[j])-1][0]][st[j][len(st[j])-1][1]][1]
           // aa[st[j][0][0]][st[j][0][1]] = aa[st[j][len(st[j])-1][0]][st[j][len(st[j])-1][1]]
    
            
    
            for k := len(st[j]) - 1; k > 1; k--{
                am[st[j][k][0]][st[j][k][1]][1] = am[st[j][k-1][0]][st[j][k-1][1]][1]
               // aa[st[j][k][0]][st[j][k][1]] = aa[st[j][k-1][0]][st[j][k-1][1]]
              // fmt.Println("HI")
               }
                
               if(len(st[j]) > 1){
                am[st[j][1][0]][st[j][1][1]][1] = tem 
              //  aa[st[j][1][0]][st[j][1][1]] = tema 
             }
            }
    

        offrand1 := 0
        offrand2 := 0


		v := [2]int{-3,3}
		offrand1 = (rand.Intn(int(float64(imDatan.Bounds().Max.X)))*v[rand.Intn(2)] )
		offrand2 = (rand.Intn(int(float64(imDatan.Bounds().Max.X)))*v[rand.Intn(2)] )
		

    for x := imDatan.Bounds().Min.X ; x < imDatan.Bounds().Max.X; x++ { 
        if(rand.Intn(20) > 10){
            v := [2]int{-3,3}
            offrand1 = (rand.Intn(int(float64(imDatan.Bounds().Max.X)))*v[rand.Intn(2)] )
            offrand2 = (rand.Intn(int(float64(imDatan.Bounds().Max.X)))*v[rand.Intn(2)] )
        }
        if(rand.Intn(1000) < 12){
            vo := []int{2,4,-3,1,3,-4,-7,7,14,-14}
            o = 14*vo[rand.Intn(6)]
            }else{
                o = 0
            }   
        
        
        for y := imDatan.Bounds().Min.Y ; y < imDatan.Bounds().Max.Y; y++ {


            if(rand.Intn(20) > 10){
                v := [2]int{-3,3}
                offrand1 = (rand.Intn(int(float64(imDatan.Bounds().Max.X)))*v[rand.Intn(2)] )
                offrand2 = (rand.Intn(int(float64(imDatan.Bounds().Max.X)))*v[rand.Intn(2)] )
            }
        
            
            aR, aG, aB, aA := imDatan.At(x, y).RGBA() // no more error
        _ ,_,_= aR, aG, aB
        index := int((float64(am[x][y][0])/float64(glen))*float64(len(rcolors)))
        rc  := rcolors[index]
        // fmt.Println(index)
        // i<29 || i>31 
        if(true){
            if(rand.Intn(255) > 2){
                if(am[x][y][1] == 0){
                          
                    pimage(dc, color.RGBA{uint8( rc[0]),uint8(rc[1]), uint8(rc[2]), uint8(aA  >>  8) }, (x + 1) * 14 + xo[y] ,o + 16 + 14 + (y+3)*14,string(gscale[am[x][y][0]]))
                  pimage(dc, color.RGBA{uint8( rc[0]),uint8(rc[1]), uint8(rc[2]), uint8(((aA/8)) >>  8) }, (x + 1) * 14 + 7 + xo[y] ,o + 16 + 21 + (y+3)*14,string(gscale[am[x][y][0]]))
                    
    
    
                   }
                   if(am[x][y][1] == 1){
                    
                    pimage(dc, color.RGBA{255, 255, 255, uint8(aA >>  8) }, (x + 1) * 14 + xo[y] ,o + 14 + 16 + (y+3)*14,string(gscale[am[x][y][0]]))
                    pimage(dc, color.RGBA{255, 255, 255, uint8(((aA/8)) >>  8) }, (x + 1) * 14 + 7 + xo[y] , o + 16 + 21 + (y+3)*14,string(gscale[am[x][y][0]]))
    
                   }  }else{
                    if(am[x][y][1] == 0){
                    
                        rpimage(dc, color.RGBA{uint8( rc[0]),uint8(rc[1]), uint8(rc[2]), uint8(aA  >>  8) }, (x + 1) * 14 + xo[y] ,o + 16 + 14 + (y+3)*14,string(gscale[am[x][y][0]]))
                      rpimage(dc, color.RGBA{uint8( rc[0]),uint8(rc[1]), uint8(rc[2]), uint8(((aA/8)) >>  8) }, (x + 1) * 14 + 7 + xo[y] ,o + 16 + 21 + (y+3)*14,string(gscale[am[x][y][0]]))
                        
        
        
                       }
                       if(am[x][y][1] == 1){
                        
                        pimage(dc, color.RGBA{255, 255, 255, uint8(aA >>  8) }, (x + 1) * 14 + xo[y] ,o + 16 + 14 + (y+3)*14,string(gscale[am[x][y][0]]))
                        pimage(dc, color.RGBA{255, 255, 255, uint8(((aA/8)) >>  8) }, (x + 1) * 14 + 7 + xo[y] , o + 16 + 21 + (y+3)*14,string(gscale[am[x][y][0]]))
        
                       } 

            }
            } else{
            
            // if(am[x][y][1] == 0){
            
            //     pimage(dc, color.RGBA{uint8( rc[0]),uint8(rc[1]), uint8(rc[2]), uint8(aA  >>  8) },
            //         x * 14 + offrand1 ,16 + 3 *14 + (y+3)*14 + offrand2,string(gscale[am[x][y][0]]))
            // pimage(dc, color.RGBA{uint8( rc[0]),uint8(rc[1]), uint8(rc[2]), uint8(((aA/8)) >>  8) },
            //     x * 14 + 21 +offrand1,21 + 2 * 14 + (y+3)*14 + offrand2,string(gscale[am[x][y][0]]))
                


            // }
            // if(am[x][y][1] == 1){
            
            // pimage(dc, color.RGBA{255, 255, 255, uint8(aA >>  8) },
            //     x * 14 + offrand1 ,14 + (y+3)*14 + offrand2,string(gscale[am[x][y][0]]))
            // pimage(dc, color.RGBA{255, 255, 255, uint8(((aA/8)) >>  8) },
            //     x * 14 + 21 +offrand1,21 + 2 * 14  + (y+3)*14 + offrand2,string(gscale[am[x][y][0]]))

            // }

        }

    }
        
    
           

            
       
    }

  _ , _ = offrand1,offrand2

  

  for j := 0; j < glen; j++{
  rc  := rcolors[int(gscale[j])*(numc-1)/255]
  if(rand.Intn(255) > 2){
  pimage(dc, color.RGBA{uint8( rc[0]),uint8(rc[1]), uint8(rc[2]), 255 }, 
  j*14+15 ,26+5,string(gscale[j])) 
  }else
  {  rpimage(dc, color.RGBA{uint8( rc[0]),uint8(rc[1]), uint8(rc[2]), 255 },
   j*14+15 ,24+5,string(gscale[j]))
  }
  }
  pimage(dc, color.RGBA{uint8( emra[0]),uint8(emra[1]), uint8(emra[2]), 255 },
   880- 148 ,26+5,em) 

grand := gscale[rand.Intn(len(gscale))]

// rc  := rcolors[int(grand)*(numc-1)/255]

pimage(dc, color.RGBA{uint8( emra[0]),uint8(emra[1]), uint8(emra[2]), 255 },
 880- (48 + 21 + 14) ,26+5,fmt.Sprintf("%s:%d", string(grand),int(grand))) 



  // mapimag(dc,color.RGBA{uint8( ra[0]),uint8(ra[1]), uint8(ra[2]), 255 })
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
