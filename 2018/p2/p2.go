package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

const MAXLENGTH = 1000

func main() {
  var count, ix, freq int = 0, 0, 0
  var input [MAXLENGTH]int
  var fmap map[int]bool
  fmap = make(map[int] bool)

  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanLines)
  for scanner.Scan() {
      str := scanner.Text()
      i, err := strconv.Atoi(str)
      if err != nil {
        fmt.Println("Oops. Error reading frequency: ", str)
      }
      if count+1 >= MAXLENGTH {
        fmt.Println("Input larger than ", MAXLENGTH)
        return
      }
      input[count] = i
      count++
  }

  fmt.Println("Read ", count, " lines of input")

  var found bool = false
  for found != true {
    freq += input[ix % count]
    //fmt.Println(freq)
    if fmap[freq] == true {
      fmt.Println("Repeated frequency: ", freq)
      return
    }
    fmap[freq] = true
    ix++
  }
}
