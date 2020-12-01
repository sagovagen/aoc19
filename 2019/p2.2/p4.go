package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
    "strconv"
)

func run(code [1000]int, codelen int, nount int, verb int) int {
  code[1] = nount
  code[2] = verb
  for op := 0; op<codelen && code[op] != 99; op+=4 {
      i:=code[op+1]
      j:=code[op+2]
      k:=code[op+3]
      if (i<0 || i>=codelen) { fmt.Println("index out of bounds: ", i); os.Exit(1) }
      if (j<0 || j>=codelen) { fmt.Println("index out of bounds: ", j); os.Exit(1) }
      if (k<0 || k>=codelen) { fmt.Println("index out of bounds: ", k); os.Exit(1) }
      switch (code[op]) {
          case 1:
            code[k] = code[i] + code[j]
          case 2:
            code[k] = code[i] * code[j]
          default:
            fmt.Println("Error: illegal op code: ", code[op])
      }
  }
  return code[0]
}


func main() {
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanLines)
  for scanner.Scan() {
    var intcode [1000]int
    var newcode [1000]int

    strarr := strings.Split(scanner.Text(),",")
    codelen := len(strarr)
    for i := 0; i<codelen; i++ {
      c, err := strconv.Atoi(strarr[i])
      if err != nil {
        fmt.Println("Oops. Error reading intcode: ", strarr[i])
      } else {
        intcode[i] = c
      }
    }

    fmt.Println("intcode length: ", codelen)

    for noun:=0; noun<100; noun++ {
      for verb:=0; verb < 100; verb++ {
        newcode = intcode
        res := run(newcode, codelen, noun, verb)
        //fmt.Println(noun,verb,res)
        //fmt.Println(intcode[0],newcode[0])
        if ( res == 19690720) {
          fmt.Println("Found wanted output with noun: ", noun, " and verb: ", verb)
          fmt.Println("Result: ", noun*100+verb)
          os.Exit(1)
        }
      }
    }
  }
}
