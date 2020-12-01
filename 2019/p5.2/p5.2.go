package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
    "strconv"
    "container/list"
)

func readInput(scanner *bufio.Scanner) int {
  str := scanner.Text()
  i, err := strconv.ParseInt(str, 10, 32)
  if (err != nil) { fmt.Println("ERROR: Unable to parse input") }
  return int(i)
}

func addNewRef(refs *list.List, ref int) {
  for e := refs.Front(); e != nil; e = e.Next() {
    if (e.Value == ref) { return }
  }
  refs.PushFront(ref)
}

func printRefVals(refs *list.List, mem []int) {
  fmt.Printf("    ")
  for e := refs.Front(); e != nil; e = e.Next() {
    value := (e.Value).(int)
    if (value >= 0 && value < len(mem)) {
      fmt.Printf("[%v]:%v ",e.Value, mem[value])
    } else {
      fmt.Printf("ERROR: (%v) out of bounds\n", value)
    }
  }
  fmt.Printf("\n")
}

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanLines)
  refs := list.New()

  if (scanner.Scan()) {
    strarr := strings.Split(scanner.Text(),",")
    mem := make([]int, len(strarr), 10000)
    codelen := len(strarr)
    for i := 0; i<codelen; i++ {
      c, err := strconv.Atoi(strarr[i])
      if err != nil {
        fmt.Println("Oops. Error reading intcode: ", strarr[i])
      } else {
        mem[i] = c
      }
    }
    fmt.Println("intcode length: ", codelen)
    //fmt.Println("code at ip=179: ",mem[179])
    if (scanner.Scan()) {
      input := strings.Split(scanner.Text(),",")
      fmt.Println("input num: ", len(input))

      for ip:=0; ip<codelen; ip++ {
        fmt.Println("Parsing codes:", mem[ip:ip+4])
        switch(mem[ip]) {
        case 1: fmt.Printf("ip: %v, add [%v] + [%v] -> [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
          mem[mem[ip+3]]=mem[mem[ip+1]]+mem[mem[ip+2]];
          addNewRef(refs, mem[ip+3])
          ip+=3;
        case 2: fmt.Printf("ip: %v, mul [%v] * [%v] -> [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
          mem[mem[ip+3]]=mem[mem[ip+1]]*mem[mem[ip+2]];
          addNewRef(refs, mem[ip+3])
          ip+=3;
        case 3: fmt.Printf("ip: %v, inp -> [%v]\n", ip, mem[ip+1]);
          mem[mem[ip+1]]=readInput(scanner)
          addNewRef(refs, mem[ip+1])
          ip++;
        case 4: fmt.Printf("ip: %v, out [%v]\n", ip, mem[ip+1]);
          fmt.Println("Output: ", mem[mem[ip+1]])
          ip++;
        case 5: fmt.Printf("ip: %v, jit [%v] [%v]\n", ip, mem[ip+1], mem[ip+2]);
          if (mem[mem[ip+1]] != 0) { ip=mem[mem[ip+2]] - 1; } else { ip+=2 } // ip will be increased with 1 by loop
        case 105: fmt.Printf("ip: %v, jit %v [%v]\n", ip, mem[ip+1], mem[ip+2]);
          if (mem[ip+1] != 0) { ip=mem[mem[ip+2]] - 1; } else { ip+=2 } // ip will be increased with 1 by loop
        case 1005: fmt.Printf("ip: %v, jit [%v] %v\n", ip, mem[ip+1], mem[ip+2]);
          if (mem[mem[ip+1]] != 0) { ip=mem[ip+2] - 1; } else { ip+=2 } // ip will be increased with 1 by loop
        case 1105: fmt.Printf("ip: %v, jit %v %v\n", ip, mem[ip+1], mem[ip+2]);
          if (mem[ip+1] != 0) { ip=mem[ip+2] - 1; } else { ip+=2 } // ip will be increased with 1 by loop
        case 6: fmt.Printf("ip: %v, jif [%v] [%v]\n", ip, mem[ip+1], mem[ip+2]);
          if (mem[mem[ip+1]] == 0) { ip=mem[mem[ip+2]] - 1; } else { ip+=2 } // ip will be increased with 1 by loop
        case 106: fmt.Printf("ip: %v, jif %v [%v]\n", ip, mem[ip+1], mem[ip+2]);
          if (mem[ip+1] == 0) { ip=mem[mem[ip+2]] - 1; } else { ip+=2 } // ip will be increased with 1 by loop
        case 1006: fmt.Printf("ip: %v, jif [%v] %v\n", ip, mem[ip+1], mem[ip+2]);
          if (mem[mem[ip+1]] == 0) { ip=mem[ip+2] - 1; } else { ip+=2 } // ip will be increased with 1 by loop
        case 1106: fmt.Printf("ip: %v, jif %v %v\n", ip, mem[ip+1], mem[ip+2]);
          if (mem[ip+1] == 0) { ip=mem[ip+2] - 1; } else { ip+=2 } // ip will be increased with 1 by loop
        case 7: fmt.Printf("ip: %v, lt [%v] [%v] [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
          if (mem[mem[ip+1]] < mem[mem[ip+2]]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
          addNewRef(refs, mem[ip+3])
          ip+=3;
        case 107: fmt.Printf("ip: %v, lt %v [%v] [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
          if (mem[ip+1] < mem[mem[ip+2]]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
          addNewRef(refs, mem[ip+3])
          ip+=3;
        case 1007: fmt.Printf("ip: %v, lt [%v] %v [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
          if (mem[mem[ip+1]] < mem[ip+2]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
          addNewRef(refs, mem[ip+3])
          ip+=3;
        case 1107: fmt.Printf("ip: %v, lt %v %v [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
          if (mem[ip+1] < mem[ip+2]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
          addNewRef(refs, mem[ip+3])
          ip+=3;
        case 8: fmt.Printf("ip: %v, eq [%v] [%v] [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
          if (mem[mem[ip+1]] == mem[mem[ip+2]]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
          addNewRef(refs, mem[ip+3])
          ip+=3;
        case 108: fmt.Printf("ip: %v, eq %v [%v] [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
          if (mem[ip+1] == mem[mem[ip+2]]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
          addNewRef(refs, mem[ip+3])
          ip+=3;
        case 1008: fmt.Printf("ip: %v, eq [%v] %v [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
          if (mem[mem[ip+1]] == mem[ip+2]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
          addNewRef(refs, mem[ip+3])
          ip+=3;
        case 1108: fmt.Printf("ip: %v, eq %v %v [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
          if (mem[ip+1] == mem[ip+2]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
          addNewRef(refs, mem[ip+3])
          ip+=3;
        case 104: fmt.Printf("ip: %v, out %v\n", ip, mem[ip+1]);
          fmt.Println("Output: ", mem[ip+1])
          ip++;
        case 101: fmt.Printf("ip: %v, add %v + [%v] -> [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
          mem[mem[ip+3]]=mem[ip+1]+mem[mem[ip+2]]
          addNewRef(refs, mem[ip+3])
          ip+=3;
        case 1001: fmt.Printf("ip: %v, add [%v] + %v -> [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
          mem[mem[ip+3]]=mem[mem[ip+1]]+mem[ip+2]
          addNewRef(refs, mem[ip+3])
          ip+=3;
        case 1101: fmt.Printf("ip: %v, add %v + %v -> [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
          mem[mem[ip+3]]=mem[ip+1]+mem[ip+2]
          addNewRef(refs, mem[ip+3])
          ip+=3;
        case 102: fmt.Printf("ip: %v, mul %v * [%v] -> [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
          mem[mem[ip+3]]=mem[ip+1]*mem[mem[ip+2]]
          addNewRef(refs, mem[ip+3])
          ip+=3;
        case 1002: fmt.Printf("ip: %v, mul [%v] * %v -> [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
          mem[mem[ip+3]]=mem[mem[ip+1]]*mem[ip+2]
          addNewRef(refs, mem[ip+3])
          ip+=3;
        case 1102: fmt.Printf("ip: %v, mul %v * %v -> [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
          mem[mem[ip+3]]=mem[ip+1]*mem[ip+2]
          addNewRef(refs, mem[ip+3])
          ip+=3;
        case 99:
          fmt.Printf("ip: %v, Halt\n", ip); os.Exit(1)
        default: fmt.Printf("ERROR: Unknown opcode %v at ip=%v\n", mem[ip], ip)
        }
        printRefVals(refs, mem)
      }
    }
  }
}
