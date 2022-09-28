package main
import (
    // "fmt"
     
	// "image/color"
    
    // "github.com/fogleman/gg"  
    
     // "encoding/binary"
    // "strings"
    // "io"
   // "github.com/andybons/gogif"
   "fmt"
    "io/ioutil"
     "log"
     //"os"
)
func main(){
   
//    gscale := "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/|()1{}[]?-_+~<>i!lI;:,\"^`'. ";


// arg := os.Args[3]

//    glen := len(gscale)

//    for i := 0;i <glen;i++{
//     dc := gg.NewContext( 2 * 14, 2 * 14)

//     if err := dc.LoadFontFace("/home/mohmedbadrco/goimage/font/Audiowide-Regular.ttf", 14); err != nil {
//         panic(err)
//     }

//     dc.SetRGB(0, 0, 0)
//     dc.Fill()
//     pimage(dc, color.RGBA{255, 255,255, 255 }, 14 ,14,string(gscale[i]))
   
//         dc.SavePNG(fmt.Sprintf("f/%s %d.png", string(gscale[i]),i))

//    }


    for i:= 0;i<100;i++{

    files, err := ioutil.ReadDir("/home/mohmedbadrco/goimage/images")
    if err != nil {
        log.Fatal(err)
    }
 
    for _, f := range files {
        p := fmt.Sprintf("/home/mohmedbadrco/goimage/images/%s",f.Name())
        fmt.Println(f.Name())
             g103d2(p)
       // gnc("/home/mohmedbadrco/goimage/images/Dino4.png")
    }
}


}