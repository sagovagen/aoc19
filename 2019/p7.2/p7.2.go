package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
    "strconv"
    "log"
)

var inputs = []string{}

func toInt(input string) int {
  i, err := strconv.ParseInt(input, 10, 32)
  if (err != nil) { fmt.Println("ERROR: Unable to parse input", err) }
  return int(i)
}

func newStream() []int {
  return []int{}
}
func streamRead(stream *[]int) int {
  value := (*stream)[0]
  *stream = (*stream)[1:]
  //fmt.Printf("input stream %p: %v\n",stream, value)
  return value
}
func streamWrite(stream *[]int, value int) {
  log.Printf("output stream %p: %v\n",stream, value)
  *stream = append(*stream, value)
}
func streamPending(stream *[]int) bool {
  return len(*stream)>0
}

func run(mem []int, ip int, inputStream *[]int, outputStream *[]int) int {
  for ip<len(mem) {
    switch(mem[ip]) {
    case 1: //fmt.Printf("ip: %v, add [%v] + [%v] -> [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      mem[mem[ip+3]]=mem[mem[ip+1]]+mem[mem[ip+2]];
      ip+=4;
    case 2: //fmt.Printf("ip: %v, mul [%v] * [%v] -> [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      mem[mem[ip+3]]=mem[mem[ip+1]]*mem[mem[ip+2]];
      ip+=4;
    case 3: //fmt.Printf("ip: %v, inp -> [%v]\n", ip, mem[ip+1]);
      if !streamPending(inputStream) {
        //fmt.Printf("No more input for amp %p, passing control to next amp\n",&mem)
        return ip // wait until previous amp produces new output
      } else {
        mem[mem[ip+1]]=streamRead(inputStream)
        ip+=2;
      }
    case 4: //fmt.Printf("ip: %v, out [%v]\n", ip, mem[ip+1]);
      streamWrite(outputStream, mem[mem[ip+1]])
      ip+=2;
    case 104: //fmt.Printf("ip: %v, out %v\n", ip, mem[ip+1]);
      streamWrite(outputStream, mem[ip+1])
      ip+=2;
    case 5: //fmt.Printf("ip: %v, jit [%v] [%v]\n", ip, mem[ip+1], mem[ip+2]);
      if (mem[mem[ip+1]] != 0) { ip=mem[mem[ip+2]] } else { ip+=3 }
    case 105: //fmt.Printf("ip: %v, jit %v [%v]\n", ip, mem[ip+1], mem[ip+2]);
      if (mem[ip+1] != 0) { ip=mem[mem[ip+2]] } else { ip+=3 }
    case 1005: //fmt.Printf("ip: %v, jit [%v] %v\n", ip, mem[ip+1], mem[ip+2]);
      if (mem[mem[ip+1]] != 0) { ip=mem[ip+2] } else { ip+=3 }
    case 1105: //fmt.Printf("ip: %v, jit %v %v\n", ip, mem[ip+1], mem[ip+2]);
      if (mem[ip+1] != 0) { ip=mem[ip+2] } else { ip+=3 }
    case 6: //fmt.Printf("ip: %v, jif [%v] [%v]\n", ip, mem[ip+1], mem[ip+2]);
      if (mem[mem[ip+1]] == 0) { ip=mem[mem[ip+2]] } else { ip+=3 }
    case 106: //fmt.Printf("ip: %v, jif %v [%v]\n", ip, mem[ip+1], mem[ip+2]);
      if (mem[ip+1] == 0) { ip=mem[mem[ip+2]] } else { ip+=3 }
    case 1006: //fmt.Printf("ip: %v, jif [%v] %v\n", ip, mem[ip+1], mem[ip+2]);
      if (mem[mem[ip+1]] == 0) { ip=mem[ip+2] } else { ip+=3 }
    case 1106: //fmt.Printf("ip: %v, jif %v %v\n", ip, mem[ip+1], mem[ip+2]);
      if (mem[ip+1] == 0) { ip=mem[ip+2] } else { ip+=3 }
    case 7: //fmt.Printf("ip: %v, lt [%v] [%v] [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      if (mem[mem[ip+1]] < mem[mem[ip+2]]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
      ip+=4;
    case 107: //fmt.Printf("ip: %v, lt %v [%v] [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      if (mem[ip+1] < mem[mem[ip+2]]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
      ip+=4;
    case 1007: //fmt.Printf("ip: %v, lt [%v] %v [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      if (mem[mem[ip+1]] < mem[ip+2]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
      ip+=4;
    case 1107: //fmt.Printf("ip: %v, lt %v %v [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      if (mem[ip+1] < mem[ip+2]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
      ip+=4;
    case 8: //fmt.Printf("ip: %v, eq [%v] [%v] [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      if (mem[mem[ip+1]] == mem[mem[ip+2]]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
      ip+=4;
    case 108: //fmt.Printf("ip: %v, eq %v [%v] [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      if (mem[ip+1] == mem[mem[ip+2]]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
      ip+=4;
    case 1008: //fmt.Printf("ip: %v, eq [%v] %v [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      if (mem[mem[ip+1]] == mem[ip+2]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
      ip+=4;
    case 1108: //fmt.Printf("ip: %v, eq %v %v [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      if (mem[ip+1] == mem[ip+2]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
      ip+=4;
    case 101: //fmt.Printf("ip: %v, add %v + [%v] -> [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      mem[mem[ip+3]]=mem[ip+1]+mem[mem[ip+2]]
      ip+=4;
    case 1001: //fmt.Printf("ip: %v, add [%v] + %v -> [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      mem[mem[ip+3]]=mem[mem[ip+1]]+mem[ip+2]
      ip+=4;
    case 1101: //fmt.Printf("ip: %v, add %v + %v -> [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      mem[mem[ip+3]]=mem[ip+1]+mem[ip+2]
      ip+=4;
    case 102: //fmt.Printf("ip: %v, mul %v * [%v] -> [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      mem[mem[ip+3]]=mem[ip+1]*mem[mem[ip+2]]
      ip+=4;
    case 1002: //fmt.Printf("ip: %v, mul [%v] * %v -> [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      mem[mem[ip+3]]=mem[mem[ip+1]]*mem[ip+2]
      ip+=4;
    case 1102: //fmt.Printf("ip: %v, mul %v * %v -> [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      mem[mem[ip+3]]=mem[ip+1]*mem[ip+2]
      ip+=4;
    case 99:
      //fmt.Printf("ip: %v, Halt\n", ip);
      return -1
    default: fmt.Printf("ERROR: Unknown opcode %v at ip=%v\n", mem[ip], ip)
    }
  }
  return -1
}

func copyIntcode(intcode []int) []int {
  copy := make([]int, len(intcode))
  for i:=0;i<len(intcode);i++ {
    copy[i] = intcode[i]
  }
  return copy
}

func main() {
  if (len(os.Args) != 7) {
    fmt.Printf("usage: %v intcodefile setting\n", os.Args[0])
    os.Exit(2)
  }
  program := os.Args[1]
  config := os.Args[2:]
  n := 5

  if (len(config) != n) {
    fmt.Println("Error, need 5 amplifier settings, you only supplied ", len(config))
    os.Exit(2)
  }
  settings := make([]int, n)
  for i:=0;i<n;i++ {
    settings[i] = toInt(config[i])
  }

  pfile, err := os.Open(program)
  if err != nil {
    fmt.Println("Error reading program file")
    os.Exit(2)
  }
  defer pfile.Close()

  // Read program
  pscanner := bufio.NewScanner(pfile)
  pscanner.Split(bufio.ScanLines)
  if (!pscanner.Scan()) {
    fmt.Println("Error reading intcode program")
    os.Exit(2)
  }
  str := pscanner.Text()
  //fmt.Println("input: ",str)
  strarr := strings.Split(str,",")
  mem := make([]int, len(strarr)+1)
  codelen := len(strarr)
  for i := 0; i<codelen; i++ {
    c, err := strconv.Atoi(strarr[i])
    if err != nil {
      fmt.Println("Oops. Error reading intcode: ", strarr[i])
    } else {
      mem[i] = c
    }
  }

  // Setup all amplifiers
  amps := [][]int{}
  for i:=0;i<n;i++ {
    amps = append(amps, copyIntcode(mem))
  }

  // Create input/output streams
  instreams := [][]int{}
  for i:=0;i<n;i++ {
    instreams = append(instreams, newStream())
    instreams[i] = append(instreams[i], settings[i])
  }
  // Insert intial input to first amp
  instreams[0] = append(instreams[0], 0)

  // Create ip array
  ip := make([]int, n)
  for true {
    allDone := true
    for i:=0;i<n;i++ {
      if (ip[i]>=0) {
        ip[i] = run(amps[i],ip[i],&instreams[i], &instreams[(i+1)%n])
        //fmt.Printf("Amp %v stopped at ip: %v\n",i,ip[i])
        allDone = false
      } else {
        //fmt.Printf("AMP no %v has finished\n",i)
      }
    }
    if (allDone) {
      fmt.Printf("%v\n",instreams[0][0])
      break
    }
  }
}
