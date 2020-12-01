package main

import (
    "fmt"
    "bufio"
    "os"
    "math"
    "strconv"
)

func main() {
  var tot int = 0
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        str := scanner.Text()
        mass, err := strconv.Atoi(str)
        if err != nil {
          fmt.Println("Oops. Error reading module weight: ", str)
        } else {
          var fuel int = int(math.Floor(float64(mass) / 3) - 2)
          fmt.Println("Mass: ", mass, " Fuel: ", fuel)
          tot += fuel
        }
    }
    fmt.Println("Total fuel: ", tot)
}
