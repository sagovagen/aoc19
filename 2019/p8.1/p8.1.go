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

  minzeros := numlayers+1
  minlayer := -1
  for i:=0;i<numlayers;i++ {
    numzeros := 0
    for j:=0;j<layersize;j++ {
      if (imagestr[i*layersize+j] == '0') {
        numzeros++
      }
    }
    if (numzeros < minzeros) {
      minzeros = numzeros
      minlayer = i
      fmt.Println("New min layer: ", i)
      fmt.Println("New min zeros:", numzeros)
    }
  }
  fmt.Println("Layer with least number of zeros: ", minlayer)
  fmt.Println("Number of zeros: ", minzeros)

  numones := 0
  numtwos := 0
  for i:=0;i<layersize;i++ {
    if (imagestr[minlayer*layersize+i] == '1') {
      numones++
    }
    if (imagestr[minlayer*layersize+i] == '2') {
      numtwos++
    }
  }
  fmt.Printf("Number of ones in layer %v: %v\n", minlayer, numones)
  fmt.Printf("Number of twos in layer %v: %v\n", minlayer, numtwos)
  fmt.Println("ones*twos: ", numones*numtwos)
}
