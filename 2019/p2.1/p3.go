package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
    "strconv"
)


func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        strarr := strings.Split(scanner.Text(),",")
        codelen := len(strarr)
        var intcode [10000]int
        for i := 0; i<codelen; i++ {
          c, err := strconv.Atoi(strarr[i])
          if err != nil {
            fmt.Println("Oops. Error reading intcode: ", strarr[i])
          } else {
            intcode[i] = c
          }
        }

        fmt.Println("intcode length: ", codelen)

        // Set 1202 alarm state
        intcode[1] = 12
        intcode[2] = 2
        for op := 0; op<codelen && intcode[op] != 99; op+=4 {
            fmt.Println("opcode: ", intcode[op])
            i:=intcode[op+1]
            j:=intcode[op+2]
            k:=intcode[op+3]
            if (i<0 || i>=codelen) { fmt.Println("index out of bounds: ", i); os.Exit(1) }
            if (j<0 || j>=codelen) { fmt.Println("index out of bounds: ", j); os.Exit(1) }
            if (k<0 || k>=codelen) { fmt.Println("index out of bounds: ", k); os.Exit(1) }
            switch (intcode[op]) {
                case 1:
                  intcode[k] = intcode[i] + intcode[j]
                case 2:
                  intcode[k] = intcode[i] * intcode[j]
                default:
                  fmt.Println("Error: illegal op code: ", intcode[op])
            }
            fmt.Println("res: ", intcode[k])
        }
        fmt.Println(intcode[:codelen])
    }


}
