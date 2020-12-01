package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)


func main() {
  var freq int = 0
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        str := scanner.Text()
        i, err := strconv.Atoi(str)
        if err != nil {
          fmt.Println("Oops. Error reading frequency: ", str)
        } else {
          freq += i
        }
    }
    fmt.Println("Final frequency: ", freq)
}
