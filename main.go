package main
import (
    "fmt"
    "bufio"
	// "image/color"
    
    // "github.com/fogleman/gg"  
    
    //  "encoding/binary"
    // "strings"
    // "io"
//    "github.com/andybons/gogif"
  // "fmt"
    "io/ioutil"
     "log"
     "image/gif"
    "os"
)
func main(){
   
//     gscale := "$@B%8&WMÄZO0QLCJUYX#*oahkbdpqwmzcvűزظunxrjft?-_+~<>i!lI½¼;:,\"^`'. "


// // arg := os.Args[3]

//    glen := len(gscale)

//    for i := 0;i <glen;i++{
//     dc := gg.NewContext( 2 * 14, 2 * 14)

//     if err := dc.LoadFontFace("/home/mohmedbadrco/goimage/font/Audiowide-Regular.ttf", 14); err != nil {
//         panic(err)
//     }

//     dc.SetRGB(0, 0, 0)
//     dc.Fill()
//     pimage(dc, color.RGBA{255, 255,255, 255 }, 14 ,14,string(gscale[i]))
//     fmt.Println(i,string(gscale[i]))
//         dc.SavePNG(fmt.Sprintf("f/%s %d.png", string(gscale[i]),i))

//    }


    for i:= 0;i<1;i++{

    files, err := ioutil.ReadDir("/home/mohmedbadrco/goimage/images")
    if err != nil {
        log.Fatal(err)
    }
 
    for _, f := range files {
        p := fmt.Sprintf("/home/mohmedbadrco/goimage/images/%s",f.Name())
        fmt.Println(f.Name())
            //  g103d8(p)
            //  g103d8(p)
            inputFile , err := os.Open(p)
            defer inputFile.Close()
            if err != nil {
                panic(err)
            }
        
            r := bufio.NewReader(inputFile)
           g,err := gif.DecodeAll(r)
           fmt.Println(g.Image[0].Bounds().Max.X ,err)

       // gnc("/home/mohmedbadrco/goimage/images/Dino4.png")
    }
}


}