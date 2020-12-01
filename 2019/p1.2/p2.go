package main

import (
    "fmt"
    "bufio"
    "os"
    "math"
    "strconv"
)


func calc_fuel(mass int) int {
  if ( mass < 9) { return 0 }
   var fuel = int(math.Floor(float64(mass) / 3.0) - 2)
   return fuel + calc_fuel(fuel)
}

func main() {
  var tot int = 0
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        str := scanner.Text()
        mass, err := strconv.Atoi(str)
        if err != nil {
          fmt.Println("Oops. Error reading module mass: ", str)
        } else {
          var fuel int = calc_fuel(mass)
          fmt.Println("Mass: ", mass, " Fuel: ", fuel)
          tot += fuel
        }
    }
    fmt.Println("Total fuel: ", tot)
}
