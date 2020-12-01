package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanLines)
  scanner.Scan()
  imagestr := scanner.Text()
  numpixels := len(imagestr)

  layersize := 6 * 25
  numlayers := numpixels / layersize
  fmt.Println("Num pixels: ", numpixels)
  fmt.Println("Num layers: ", numlayers)
  image := make([]int, layersize)
  for i:=0;i<layersize;i++ { image[i]=2 }
  fmt.Println("Rendered image size: ", len(image))

  for i:=0;i<layersize;i++ {
    p := i
    for j:=0;j<numlayers;j++ {
      if (imagestr[j*layersize+i]!='2') {
        p = j*layersize+i
        break
      }
    }
    switch(imagestr[p]) {
    case '0': image[i] = 0;
    case '1': image[i] = 1;
    case '2': image[i] = 2;
    default: image[i] = 9;
    }
  }

  for y:=0;y<6;y++ {
    for x:=0;x<25;x++ {
      p := " "
      if (image[y*25+x] == 1) {
        p = "*"
      }
      fmt.Printf("%v",p)
    }
    fmt.Println()
  }
}
