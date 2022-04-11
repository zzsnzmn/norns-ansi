package main

import (
  "fmt"
  "bytes"
  "net/http"
  "image"
  "os"
  "log"
  _ "image/png"
  //"math"
  "image/color"

  //"github.com/disintegration/imaging"
  "github.com/kevin-cantwell/dotmatrix"
  //"github.com/nfnt/resize"
)

func Encode(img image.Image) error {
  return dotmatrix.Print(os.Stdout, img)
}

func String(img image.Image) string {
  buffer := &bytes.Buffer{}
  dotmatrix.Print(buffer, img)
  return buffer.String()
}

func getImageFromFilePath(filePath string) (image.Image, error) {
    f, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    image, _, err := image.Decode(f)
    return image, err
}

func getImageFromURL(url string) (image.Image, error) {
  resp, err := http.Get(url)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  image, _, err := image.Decode(resp.Body)
    return image, err
}

func main() {

  img, err := getImageFromURL("http://localhost:8889")


  if err != nil {
    log.Fatal(err)
  }
  //m := resize.Resize(128, 64, img, resize.NearestNeighbor)
  //img = imaging.Invert(img)
  //log.Println(fmt.Sprintf("%+v", img.Bounds()))
  //Encode(m)
  topLeft := image.Point{0, 0}
  bottomRight := image.Point{128, 64}
  m := image.NewRGBA(image.Rectangle{topLeft, bottomRight})
  a := 0
  b := 0
  for x := 0; x < 512; x = x + 4 {
    for y := 0; y < 256; y = y + 4 {
      b++
      //fmt.Println(a, b)
      //fmt.Println(img)
      //r, g, b, a := img.At(x, y).RGBA()
      red, _, blue, _ := img.At(x, y).RGBA()
      if red == 2056 && blue == 2056   {
      //fmt.Println(fmt.Sprintf("%v", alpha))
        m.Set(a, b, color.White)
      //} else {
        //fmt.Println(fmt.Sprintf("%v", img.At(x, y)))
        //m.Set(a, b, color.Gray16{0})
      } else {
      //fmt.Println(fmt.Sprintf("%v", alpha))
        m.Set(a, b, color.Black)
      }
    }
    a++
    b = 0
  }

  fmt.Print(String(m))
}
