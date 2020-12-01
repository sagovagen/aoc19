package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
    "strconv"
    "container/list"
)

func readInput(scanner bufio.Scanner) int {
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

  for scanner.Scan() {
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
    fmt.Println("code at ip=179: ",mem[179])

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
        mem[mem[ip+1]]=1
        addNewRef(refs, mem[ip+1])
        ip++;
      case 4: fmt.Printf("ip: %v, out [%v]\n", ip, mem[ip+1]);
        fmt.Println("Output: ", mem[mem[ip+1]])
        ip++;
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
