package main

import (
    "fmt"
    "strings"
)

func numSame(d []string, i int) int {
    n := len(d)
    j := i
    for j<n && d[j]==d[i] { j++ }
    return j-i
}

func ok(code int) bool {
  //fmt.Println("code: ", code)
  digits := strings.Split(fmt.Sprintf("%d",code),"")
  var maxix=len(digits)
  var doubles bool = false


  for ix:=0;ix<maxix; {
    num := numSame(digits, ix)
    if (num == 2) { doubles = true }
    if (ix>0 && digits[ix]<digits[ix-1]) { return false }
    ix+=num
  }
  return doubles
}

func main() {
  var startcode int = 256310
  var stopcode int = 732736
  var numvalid int = 0
  for i:=startcode;i<=stopcode;i++ {
    if (ok(i)) {
      numvalid++
      fmt.Println(i)
    }
  }
  fmt.Println("Num valid codes: ", numvalid)
}
