package main

import (
    "fmt"
    "strings"
)

func ok(code int) bool {
  digits := strings.Split(fmt.Sprintf("%d",code),"")
  var doubles bool = false
  for i:=1;i<len(digits);i++ {
    if (digits[i] < digits[i-1]) { return false }
    if (digits[i] == digits[i-1]) { doubles=true }
  }
  return doubles
}


func main() {
  var startcode int = 256310
  var stopcode int = 732736
  var numvalid int = 0
  for i:=startcode;i<=stopcode;i++ {
    if (ok(i)) { numvalid++ }
  }
  fmt.Println("Num valid codes: ", numvalid)
}
