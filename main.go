package main

import (
  "fmt"
  "bytes"
  "net/http"
  "image"
  "os"
  "log"
  _ "image/png"

  "github.com/disintegration/imaging"
  "github.com/kevin-cantwell/dotmatrix"
  "github.com/nfnt/resize"
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
  m := resize.Resize(128, 64, imaging.Invert(img), resize.NearestNeighbor)
  //log.Println(fmt.Sprintf("%+v", img.Bounds()))
  //Encode(m)
  fmt.Print(String(m))
}
